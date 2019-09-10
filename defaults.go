// This is just a for convenience sake, syntax sugar, no magic
package worst

import (
	"github.com/badtheory/informer"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

func (w * Worst) SetMiddlewareDefaults() {
	w.SetRequestId()
	w.SetLogger()
	w.SetInformer()
	w.SetRecover()
	w.SetCompress(3)
}

func (w * Worst) SetSecurityDefaults() {
	co, so := w.Security.fuse()
	w.Router.Use(w.Middleware.Secure(so))
	w.Router.Use(w.Middleware.Cors(co))
}

func (w *Worst) SetRequestId() {
	w.Router.Use(w.Middleware.RequestId)
}

func (w *Worst) SetLogger() {
	w.Router.Use(w.Middleware.Logger)
}

func (w *Worst) SetRecover() {
	w.Router.Use(w.Middleware.Recover)
}

func (w *Worst) SetCompress(level int, types ...string) {
	w.Router.Use(w.Middleware.Compress(level, types...))
}

func (w *Worst) SetInformer(opt ...informer.Configuration) {
	w.Router.Use(w.Middleware.Informer(opt...))
}

func (w *Worst) SetHeartbeat(endpoint string) {
	w.Router.Use(w.Middleware.Heartbeat(endpoint))
}

func (w *Worst) SetStatic(urlPrefix, location string, index bool) {
	w.Router.Use(w.Middleware.Static(urlPrefix, location, index))
}

func (w *Worst) Cors(options cors.Options) {
	w.Router.Use(w.Middleware.Cors(options))
}

func (w *Worst) Secure(options secure.Options) {
	w.Router.Use(w.Middleware.Secure(options))
}

