package runserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"olympics/pkg/core"
	"olympics/pkg/core/actions"
	"olympics/pkg/server"
	"olympics/pkg/storage/db"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var ConfigPath string

func Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	fmt.Println(ConfigPath)
	configLoader := confita.NewLoader(file.NewBackend(ConfigPath))
	cfg, err := core.ParseCfg(ctx, configLoader)
	if err != nil {
		return err
	}

	storage, err := db.NewStorage(&cfg.Storage)
	if err != nil {
		return err
	}

	repo := &actions.Repo{Storage: storage, Config: cfg, Actions: actions.NewActions(storage)}
	app := server.NewApp(repo)

	gr, ctx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		<-app.Done()
		return app.Stop(ctx)
	})

	gr.Go(func() error {
		err := app.Start(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	return gr.Wait()
}
