package worst

import (
	"github.com/badtheory/informer"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func Log() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				ctx := informer.WithFields(
					informer.Fields{
						"proto":   r.Proto,
						"path":    r.URL.Path,
						"latency": time.Since(t1),
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
