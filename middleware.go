package worst

import (
	"github.com/badtheory/informer"
	"github.com/badtheory/static"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type Middleware struct {
	http.Handler
}

type PlugInPlay interface {
	Logger(next http.Handler) http.Handler
	Compress(next http.Handler) http.Handler
	Heartbeat(next http.Handler) http.Handler
	Informer(next http.Handler) http.Handler
	Static(next http.Handler) http.Handler
}

func (m Middleware) Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

func (m Middleware) Compress(level int, types ...string) func(next http.Handler) http.Handler {
	return middleware.Compress(level, types...)
}

func (m Middleware) Heartbeat(endpoint string) func(http.Handler) http.Handler {
	return middleware.Heartbeat(endpoint)
}

func (m Middleware) Informer(opt ...informer.Configuration) func(next http.Handler) http.Handler {
	var o informer.Configuration
	if len(opt) == 0 {
		o = informer.Configuration{}
	} else {
		o = opt[0]
	}

	err := informer.NewLogger(o, informer.InstanceZapLogger)
	if err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()
			stop := time.Now()
			defer func() {
				ctx := informer.WithFields(
					informer.Fields{
						"proto":   r.Proto,
						"path":    r.URL.Path,
						"latency": stop.Sub(start).String(),
						"status":  ww.Status(),
						"size":    ww.BytesWritten(),
						"reqId":   middleware.GetReqID(r.Context()),
					},
				)
				ctx.Infof("req_level_log")
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}

}

func (m Middleware) Static(urlPrefix, location string, index bool) func(next http.Handler) http.Handler {
	return static.Serve(urlPrefix, static.LocalFile(location, index))
}