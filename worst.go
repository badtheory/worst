package worst

import (
	"fmt"
	"github.com/badtheory/informer"
	"github.com/creasty/defaults"
	"github.com/go-chi/chi"
	. "github.com/logrusorgru/aurora"
	"github.com/unrolled/render"
	"net/http"
)

type Worst struct {
	Router   *Router
	Server   *http.Server `default:"{\"Addr\": \"127.0.0.1:1337\"}"`
	Security Security
	Middleware Middleware
}

type Router struct {
	Render *render.Render
	*chi.Mux
}

type Options struct {
	Static   Static
	Server   *http.Server
	Render   *render.Render
	Logger   informer.Configuration
}

type Static struct {
	Url  string `default:"/*"`
	Path string `default:""`
}

func New() *Worst {
	w := &Worst{
		Router: &Router{
			render.New(),
			chi.NewRouter(),
		},
		Security: Security{},
	}

	if err := defaults.Set(w); err != nil {
		panic(err)
	}

	return w

}


func (w *Worst) Run() {
	w.Server.Handler = w.Router
	fmt.Println(Gray(1-1, Bold("Worst HTTP running on "+w.Server.Addr)).BgGray(24 - 1))
	if err := w.Server.ListenAndServe(); err == nil {
		fmt.Println(Red("Worst HTTP running on " + w.Server.Addr).BgGray(24 - 1))
	}
}
