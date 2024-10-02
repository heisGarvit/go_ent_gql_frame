//go:generate ./src/utils/copier.sh
//go:generate go mod tidy
//go:generate go run -mod=mod ./src/utils/entc.go
//go:generate go mod tidy
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate git add src/models/ent
//go:generate go run -mod=mod ./src/utils/bgtasksgenerator/generator.go

package main

import (
	"flag"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/ravilushqa/otelgqlgen"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"log"
	"net/http"
	"project/src/infra/db"
	"project/src/infra/env"
	"project/src/infra/taskqueue"
	"project/src/middlewares/apmspan"
	"project/src/middlewares/authtoken"
	"project/src/middlewares/cors"
	"project/src/models/ent/resolvers"
)

func main() {
	defer db.Client.Close()

	port := flag.String("port", "8080", "Port to run the server on")
	bgTask := flag.Bool("bgTask", false, "Run background task worker")
	flag.Parse()

	srv := handler.New(resolvers.NewSchema(db.Client))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.Use(otelgqlgen.Middleware())

	mux := http.NewServeMux()
	httpHandler := cors.Middleware(srv)
	httpHandler = authtoken.Middleware(httpHandler)
	httpHandler = apmspan.Middleware(httpHandler)

	mux.Handle("/query", httpHandler)

	h := http.Handler(mux) // or use your router
	h = otelhttp.NewHandler(h, "")

	if env.Get("GRAPHQL_INTROSPECTION") == "true" {
		mux.HandleFunc("/explorer", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://studio.apollographql.com/sandbox/explorer?endpoint=http://"+r.Host+"/query", http.StatusSeeOther)
		})
		srv.Use(extension.Introspection{})
	}

	if bgTask != nil && *bgTask {
		go taskqueue.LaunchWorker()
	}

	log.Printf("connect to http://127.0.0.1:%s/explorer for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":"+*port, h))

}
