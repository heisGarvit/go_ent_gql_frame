package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"io/ioutil"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"project/bgTasks"
	"project/utils/taskQueue"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"text/template"
)

type TemplateArgs struct {
	Tasks []Task
}

type ArgType struct {
	ArgName  string
	GoType   string
	Category string
}

type Task struct {
	Name      string
	Args      []ArgType
	Signature string
	Ctx       string
}

var directoryContent map[string]string = make(map[string]string)

func main() {
	bgTasks.ImportWorkAround = ""
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var tasksList []Task

	for _, taskName := range taskQueue.Server.GetRegisteredTaskNames() {
		taskFunc, err := taskQueue.Server.GetRegisteredTask(taskName)
		if err != nil {
			slog.Error("Error getting task", "error", err)
		}
		if task, err := genTaskStruct(taskName, taskFunc); err != nil {
			slog.Error("Error generating task struct", "error", err)
		} else {
			tasksList = append(tasksList, task)
		}
	}
	var tmplFile = "utils/taskQueue/generator/template.gohtml"
	content, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		log.Println(err)
	} else {
		tmpl, err := template.New("t").Parse(string(content))
		if err != nil {
			panic(err)
		}
		f, err := os.Create("background/background.go")
		if err != nil {
			log.Println("create file: ", err)
			return
		}
		err = tmpl.Execute(f, TemplateArgs{Tasks: tasksList})
		if err != nil {
			panic(err)
		}
	}

}

func genTaskStruct(taskName string, taskFunc interface{}) (Task, error) {

	task := Task{
		Name: taskName,
		Ctx:  "context.Background()",
	}

	splits := strings.Split(runtime.FuncForPC(reflect.ValueOf(taskFunc).Pointer()).Name(), ".")

	code := readAllFilesInDirectory(strings.Replace(splits[0], "project", ".", 1))

	signature, err := extractFunctionSignature(splits[1], code)
	if err != nil {
		return Task{}, err
	}
	task.Signature = signature
	signature = strings.Replace(signature, "(", "", 1)
	signature = strings.Replace(signature, ")", "", 1)

	taskFuncType := reflect.TypeOf(taskFunc)
	if taskFuncType.NumIn() <= 0 {
		return task, nil
	}

	signatureArgs := strings.Split(signature, ",")

	iterateFrom := 0

	arg0Type := taskFuncType.In(0)
	if tasks.IsContextType(arg0Type) {
		task.Ctx = getSignatureArgDetails(signatureArgs[0]).ArgName
		iterateFrom = 1
	}

	for i := iterateFrom; i < taskFuncType.NumIn(); i++ {
		task.Args = append(task.Args, getSignatureArgDetails(signatureArgs[i]))
	}

	return task, nil
}

func getSignatureArgDetails(signatureArg string) ArgType {
	signatureArg = strings.TrimSpace(signatureArg)
	a := ArgType{}

	if strings.Contains(signatureArg, "...") {
		a.Category = "variadic"
		signatureArg = strings.Replace(signatureArg, "...", "", 1)
	} else {
		a.Category = "normal"
	}

	splits := strings.Split(signatureArg, " ")
	if len(splits) < 2 {
		return a
	}
	a.GoType = splits[1]
	a.ArgName = splits[0]

	return a
}

func readAllFilesInDirectory(directoryName string) string {
	if filesData, ok := directoryContent[directoryName]; ok {
		return filesData
	}

	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		log.Println(err)
		return ""
	}

	directoryContent[directoryName] = ""

	// Iterate through each file
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directoryName, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Println(err)
			} else {
				directoryContent[directoryName] += string(content)
			}
		}
	}

	return directoryContent[directoryName]

}

func extractFunctionSignature(funcName string, fileContent string) (string, error) {
	// Define the regular expression pattern to match the function signature
	pattern := fmt.Sprintf(`func\s+%s\s*\(([^)]*)\)`, funcName)
	re := regexp.MustCompile(pattern)

	// Find the submatches
	matches := re.FindStringSubmatch(fileContent)
	if len(matches) < 2 {
		return "", fmt.Errorf("function %s not found", funcName)
	}

	// Extract and trim the function signature
	signature := strings.TrimSpace(matches[1])

	return "(" + signature + ")", nil
}
