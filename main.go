//go:generate ./utils/copier.sh
//go:generate go mod tidy
//go:generate go run -mod=mod ./utils/entc.go
//go:generate go mod tidy
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate git add ent
//go:generate go run -mod=mod ./utils/taskQueue/generator/generator.go

package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ravilushqa/otelgqlgen"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net"
	"net/http"
	"os"
	"project/ent/resolvers"
	"project/utils/apm"
	"project/utils/db"
	"project/utils/env"
	"project/utils/taskQueue"
	"strings"
)

type doneWriter struct {
	http.ResponseWriter
	ctx context.Context
	r   *http.Request
}

func (w *doneWriter) Write(b []byte) (int, error) {
	if resolverName, ok := w.ctx.Value("ResolverSpanNames").(*apm.ResolverSpanNames); ok {
		trace.SpanFromContext(w.ctx).SetName(strings.Join(resolverName.Names[:], " "))
	}
	return w.ResponseWriter.Write(b)
}

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		ctx := context.WithValue(r.Context(), "Authorization", r.Header.Get("Authorization"))
		rsn := apm.ResolverSpanNames{
			Names: []string{},
		}
		ctx = context.WithValue(ctx, "ResolverSpanNames", &rsn)
		r = r.WithContext(ctx)
		dw := &doneWriter{ResponseWriter: w, ctx: ctx, r: r}

		next.ServeHTTP(dw, r)
	})
}

func main() {
	defer db.Client.Close()

	// Configure the server and start listening on :8081.
	srv := handler.New(resolvers.NewSchema(db.Client))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.Use(otelgqlgen.Middleware())

	mux := http.NewServeMux()

	mux.Handle("/query", Middleware(srv))

	h := http.Handler(mux) // or use your router
	h = otelhttp.NewHandler(h, "")

	go taskQueue.LaunchWorker()

	port := env.Get("PORT")
	if port != "" {
		mux.HandleFunc("/explorer", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://studio.apollographql.com/sandbox/explorer?endpoint=http://"+r.Host+"/query", http.StatusSeeOther)
		})
		srv.Use(extension.Introspection{})

		log.Printf("connect to http://localhost:%s/explorer for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, h))
	}

	socketPath := env.Get("SOCKET_PATH")
	if socketPath == "" {
		socketPath = "./unix.sock"
	}

	_ = os.Remove(socketPath)
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal("Listen (UNIX socket) "+socketPath+" : ", err)
	}
	err = os.Chmod(socketPath, 0777)
	if err != nil {
		log.Println(err)
	}
	log.Printf("connect to unix socket for GraphQL playground")

	defer unixListener.Close()
	log.Fatal(http.Serve(unixListener, h))
}
