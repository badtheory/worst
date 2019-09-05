package worst

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jordan-wright/unindexed"
	. "github.com/logrusorgru/aurora"
	"github.com/unrolled/render"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Worst struct {
	Router 			Router
	Security		*secure.Secure
	Options			Options
}

type Router struct {
	Render		*render.Render
	*chi.Mux
}

type Options struct {
	Security		secure.Options
	Static 			Static
	Server          *http.Server
	Render			*render.Render
}

type Static struct {
	Url 		string
	Path 		string
}

func New(opt ...Options) *Worst {

	var s Options
	if len(opt) == 0 {
		s = Options{
			secure.Options{
				STSSeconds:            31536000,
				STSIncludeSubdomains:  true,
				STSPreload:            true,
				FrameDeny:             true,
				ContentTypeNosniff:    true,
				BrowserXssFilter:      true,
				ContentSecurityPolicy: "script-src $NONCE",
			},
			Static{
				"/public/*",
				"../public",
			},
			&http.Server{
				Addr:         "localhost:1337",
				ReadTimeout:  60 * time.Second,
				WriteTimeout: 60 * time.Second,
				IdleTimeout:  60 * time.Second,
			},
			render.New(),
		}
	} else {

		if reflect.DeepEqual(Options{}.Security, opt[0].Security)  {
			opt[0].Security = secure.Options{
				STSSeconds:            31536000,
				STSIncludeSubdomains:  true,
				STSPreload:            true,
				FrameDeny:             true,
				ContentTypeNosniff:    true,
				BrowserXssFilter:      true,
				ContentSecurityPolicy: "script-src $NONCE",
			}
		}

		if (Options{}.Static == opt[0].Static)  {
			opt[0].Static = Static{
				"/public/*",
				"../public",
			}
		}

		if (Options{}.Render == opt[0].Render)  {
			opt[0].Render = render.New()
		}

		s = opt[0]
	}

	secureMiddleware := secure.New(s.Security)

	w := &Worst{
		Router: Router{
			s.Render,
			chi.NewRouter(),
		},
		Security: secureMiddleware,
		Options: s,
	}

	w.Router.Use(secureMiddleware.Handler)
	w.Router.Use(middleware.RequestID)
	w.Router.Use(middleware.Logger)
	w.Router.Use(middleware.Recoverer)
	w.Router.Use(middleware.Compress(3))
	w.Router.Use(middleware.Timeout(60 * time.Second))
	w.Router.Handle(s.Static.Url, http.Handler(http.FileServer(unindexed.Dir(s.Static.Path))))

	w.Options.Server.Handler = w.Router

	return w

}

func (w *Worst) Run() {
	log.Println(Gray(1-1, Bold("Worst HTTP running on " + w.Options.Server.Addr)).BgGray(24-1))
	if err := w.Options.Server.ListenAndServe(); err == nil {
		log.Println(Red("Worst HTTP running on " + w.Options.Server.Addr).BgGray(24-1))
	}
}