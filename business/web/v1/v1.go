package v1

import (
	"os"

	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/auth"
	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/mid"
	"github.com/SyedMohamedHyder/ecoverify/foundation/logger"
	"github.com/SyedMohamedHyder/ecoverify/foundation/web"
	"github.com/jmoiron/sqlx"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
	Auth     *auth.Auth
	DB       *sqlx.DB
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) *web.App {
	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
		mid.Metrics(),
	)

	routeAdder.Add(app, cfg)

	return app
}
