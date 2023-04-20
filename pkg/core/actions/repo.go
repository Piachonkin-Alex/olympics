package actions

import (
	"olympics/pkg/core"
	"olympics/pkg/storage"
)

type Repo struct {
	Actions *Actions
	Config  *core.Config
	Storage storage.Storage
}
