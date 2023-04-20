package server

import (
	"fmt"
	"net/http"

	"olympics/lib/schema"
	"olympics/pkg/core/entities"
)

func (a *App) checkAuth() func(w http.ResponseWriter, r *http.Request) bool {
	handlersToRole := entities.ConstructHandlersMap(a.repo.Config.Auth.Handlers)
	return func(w http.ResponseWriter, r *http.Request) bool {
		requiredRole, ok := handlersToRole[entities.HandlerDescription{
			Path:   r.URL.Path,
			Method: r.Method,
		}]
		if !ok {
			return true
		}

		userName := r.Header.Get("X-User")
		if userName == "" {
			_ = schema.APIError(w, http.StatusUnauthorized, fmt.Errorf("no user is specified"))
			return false
		}

		userRole, err := a.repo.Actions.GetRole(r.Context(), userName)
		if err != nil {
			_ = schema.APIError(w, http.StatusInternalServerError, err)
			return false
		}

		if userRole < requiredRole {
			_ = schema.APIError(w, http.StatusUnauthorized, fmt.Errorf("user %s has no required role", userName))
			return false
		}
		return true
	}
}
