package server

import (
	"net/http"
	"strings"

	"olympics/lib/middlewares/auth"
	"olympics/lib/server"
	"olympics/pkg/core/actions"
)

type App struct {
	*server.Server
	repo *actions.Repo
}

func NewApp(repo *actions.Repo) *App {
	app := &App{Server: server.NewServer(&repo.Config.Server), repo: repo}
	app.setMiddlewares()
	app.setRoutes()
	return app
}

func (a *App) setMiddlewares() {
	a.AddMiddleware(auth.AuthMiddleware(
		auth.WithCheck(a.checkAuth()),
		auth.WithDisabled(a.repo.Config.Auth.Disabled),
		auth.WithSkipFilter(func(r *http.Request) bool {
			return !strings.HasPrefix(r.URL.Path, "/api")
		}),
	))
}

func (a *App) setRoutes() {
	a.AddRoutes(a.GetRoutes())
}
