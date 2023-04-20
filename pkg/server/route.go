package server

import (
	"log"
	"net/http"

	"olympics/lib/server"
	"olympics/pkg/core/actions"
)

func (a *App) toHandler(repoHandler func(w http.ResponseWriter, r *http.Request, repo *actions.Repo) error) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := repoHandler(writer, request, a.repo)
		if err != nil {
			log.Printf("error during writing: %v", err)
		}
	}
}

func (a *App) GetRoutes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/api/test",
			Handler: a.toHandler(TestHandler),
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/role",
			Handler: a.toHandler(AddRoleHandler),
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/v1/athlete",
			Handler: a.toHandler(GetAthleteInfo),
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/v1/event",
			Handler: a.toHandler(PostAthleteEvent),
		},
	}

}
