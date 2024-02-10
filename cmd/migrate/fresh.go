package migration

import (
	"context"

	"github.com/urfave/cli/v2"
)

func Fresh() *cli.Command {

	return &cli.Command{
		Name:  "fresh",
		Usage: "Rollback all migrations and re-run them",
		Action: func(c *cli.Context) error {
			migrations, err := initMigration()
			if err != nil {
				return err
			}

			if err := migrations.Fresh(context.Background()); err != nil {
				return err
			}

			return nil
		},
	}
}
