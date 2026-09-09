package cloudy

import (
	"context"
	"fmt"
	"log"

	"github.com/blue-monads/potatoverse/backend/xtypes"
	turso "turso.tech/database/tursogo"
)

type SubApp struct {
	Name    string
	App     xtypes.App
	UserId  int64
	tursoDB *turso.TursoSyncDb
}

func NewSubApp(ctx context.Context, config *Config, name string, userId int64) (*SubApp, error) {

	bootstrap := false

	tursoDB, err := turso.NewTursoSyncDb(ctx, turso.TursoSyncDbConfig{
		Path:             fmt.Sprintf("sub_%s.db", name),
		RemoteUrl:        config.TursoRemoteURL,
		AuthToken:        config.TursoAuthToken,
		BootstrapIfEmpty: &bootstrap,
	})

	if err != nil {
		return nil, err
	}

	return &SubApp{
		Name:    name,
		UserId:  userId,
		tursoDB: tursoDB,
	}, nil

}

func (s *SubApp) Sync(ctx context.Context) error {

	pulled, err := s.tursoDB.Pull(ctx)
	if err != nil {
		return err
	}

	if pulled {
		log.Println("Pulled latest changes from remote database")
	}

	return nil
}

func (s *SubApp) Load() error {

	return nil
}
