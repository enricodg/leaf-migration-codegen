package command

import (
	"github.com/enricodg/leaf-migration-codegen/handler"
	"github.com/enricodg/leaf-migration-codegen/helper/generateType"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"strings"
)

func Generate() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "generate --types <type> --name <name>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "types",
				Aliases:  []string{"t"},
				Value:    "mysql,mongo,postgre",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Value:    "service_parameters",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			log := logger.GetLogger()
			migrationType := strings.ToLower(c.String("types"))
			name := strings.ToLower(c.String("name"))
			log.StandardLogger().Infof("[%s] Generating %s files...", strings.ToUpper(migrationType), name)

			if !connection.IsValid(migrationType) {
				return errors.New("invalid migration type [mysql | mongo | postgre]")
			}

			if !generateType.IsValid(name) {
				return errors.New("invalid generate type [service_parameters]")
			}

			return handler.GetHandler().Generate(migrationType, name)
		},
	}
}
