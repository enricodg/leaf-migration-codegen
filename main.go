package main

import (
	"github.com/enricodg/leaf-migration-codegen/command"
	leafMigrationCommand "github.com/paulusrobin/leaf-utilities/leafMigration/command"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "Leaf migration codegen",
		Usage:       "Supporting leaf framework to initialize migration project",
		Description: "CLI Leaf migration code generator",
		UsageText:   "command [command options] [arguments...]",
		Version:     "v1.0.0",
		Commands: []*cli.Command{
			command.Init(),
			command.Generate(),
			leafMigrationCommand.New(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.GetLogger().StandardLogger().Errorf("Run Error: %+v", err.Error())
	}
}
