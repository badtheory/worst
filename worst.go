package worst

import (
	"github.com/creasty/defaults"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jordan-wright/unindexed"
	. "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
	l "github.com/treastech/logger"
	"github.com/unrolled/render"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"time"
)

type Worst struct {
	Router 			*Router
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
	Url 		string `default:"/*"`
	Path 		string `default:""`
}

func New(opt ...Options) *Worst {

	var o Options
	var s Static

	if len(opt) == 0 {

		if err := defaults.Set(&s); err != nil {
			panic(err)
		}

		o = Options{
			secure.Options{
				STSSeconds:            31536000,
				STSIncludeSubdomains:  true,
				STSPreload:            true,
				FrameDeny:             true,
				ContentTypeNosniff:    true,
				BrowserXssFilter:      true,
				ContentSecurityPolicy: "script-src $NONCE",
			},
			s,
			&http.Server{
				Addr:         "localhost:1337",
				ReadTimeout:  60 * time.Second,
				WriteTimeout: 60 * time.Second,
				IdleTimeout:  60 * time.Second,
			},
			render.New(),
		}
	} else {

		if err := defaults.Set(&opt[0]); err != nil {
			panic(err)
		}

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

		if (Options{}.Render == opt[0].Render)  {
			opt[0].Render = render.New()
		}

		o = opt[0]
	}

	secureMiddleware := secure.New(o.Security)

	w := &Worst{
		Router: &Router{
			o.Render,
			chi.NewRouter(),
		},
		Security: secureMiddleware,
		Options: o,
	}


	logger, _ := zap.NewProduction()
	defer logger.Sync()

	w.Router.Use(
		secureMiddleware.Handler,
		middleware.RequestID,
		l.Logger(logger),
		middleware.Recoverer,
		middleware.Compress(3),
		middleware.Timeout(60 * time.Second),
	)

	w.Router.Handle(o.Static.Url, http.Handler(http.FileServer(unindexed.Dir(o.Static.Path))))
	return w

}

func (w *Worst) Run() {
	w.Options.Server.Handler = w.Router
	log.Println(Gray(1-1, Bold("Worst HTTP running on " + w.Options.Server.Addr)).BgGray(24-1))
	if err := w.Options.Server.ListenAndServe(); err == nil {
		log.Println(Red("Worst HTTP running on " + w.Options.Server.Addr).BgGray(24-1))
	}
}