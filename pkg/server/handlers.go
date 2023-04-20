package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"olympics/lib/schema"
	"olympics/pkg/core/actions"
	"olympics/pkg/core/entities"
)

func TestHandler(w http.ResponseWriter, r *http.Request, repo *actions.Repo) error {
	return schema.APIOk(w, http.StatusOK, "you are in the world of sheet")
}

// swagger:operation POST /v1/configuration/validate v1_configuration_validate
//
// Validate configuration version.
//
// ---
// consumes:
//   - application/json
//
// produces:
//   - application/json
//
// parameters:
//   - name: configuration_validate
//     in: body
//     description: start configuration in memory
//     required: true
//     schema:
//     type: object
//     properties:
//     configuration_id:
//     type: integer
//     format: int64
//     template_version:
//     type: integer
//     format: int64
//     inputs:
//     type: object
//     environment:
//     type: string
//
// responses:
//
//	200:
//	  description: OK
//	  $ref: '#/responses/APIResponse'
func AddRoleHandler(w http.ResponseWriter, r *http.Request, repo *actions.Repo) error {
	var addRoleInfo struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	log.Println("method POST '/api/role' handled")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&addRoleInfo); err != nil {
		return schema.APIError(w, http.StatusInternalServerError, fmt.Errorf("error during parsing body"))
	}

	if addRoleInfo.Name == "" {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("parameter 'name' must be specified"))
	}

	role, err := entities.RoleFromString(addRoleInfo.Role)
	if err != nil {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("invalid role is specified: %v", err))
	}

	if err := repo.Actions.AddRole(r.Context(), addRoleInfo.Name, role); err != nil {
		return schema.APIError(w, http.StatusInternalServerError, err)
	}

	return schema.APIOk(w, http.StatusOK, nil)
}

func GetAthleteInfo(w http.ResponseWriter, r *http.Request, repo *actions.Repo) error {
	log.Println("method GET '/api/v1/athlete' handled")
	name := r.URL.Query().Get("name")
	if name == "" {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("parameter 'name' must be specified"))
	}

	info, err := repo.Actions.GetAthleteInfo(r.Context(), name)
	if err != nil {
		return schema.APIError(w, http.StatusInternalServerError, err)
	}

	return schema.APIOk(w, http.StatusOK, info)
}

func PostAthleteEvent(w http.ResponseWriter, r *http.Request, repo *actions.Repo) error {
	var event entities.Athlete
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		return schema.APIError(w, http.StatusInternalServerError, fmt.Errorf("error during parsing body"))
	}

	if event.Name == "" {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("parameter 'name' must be specified"))
	}

	if event.Sport == "" {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("parameter 'sport' must be specified"))
	}

	if event.Country == "" {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("parameter 'country' must be specified"))
	}

	if event.Gold+event.Silver+event.Bronze <= 0 {
		return schema.APIError(w, http.StatusBadRequest, fmt.Errorf("at least one medal must be specified"))
	}

	if err := repo.Actions.AddAthleteEvent(r.Context(), event); err != nil {
		return schema.APIError(w, http.StatusInternalServerError, err)
	}

	return schema.APIOk(w, http.StatusOK, nil)
}
