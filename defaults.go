package worst

import "github.com/badtheory/informer"

func (w * Worst) SetDefaults() {
	w.SetLogger()
	w.SetCompress(3)
	w.SetInformer()
}

func (w *Worst) SetLogger() {
	w.Router.Use(w.Middleware.Logger)
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
