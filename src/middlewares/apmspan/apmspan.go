package apmspan

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"project/src/infra/apm"
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
		rsn := apm.ResolverSpanNames{
			Names: []string{},
		}
		ctx := context.WithValue(r.Context(), "ResolverSpanNames", &rsn)
		r = r.WithContext(ctx)
		dw := &doneWriter{ResponseWriter: w, ctx: ctx, r: r}

		next.ServeHTTP(dw, r)
	})
}
