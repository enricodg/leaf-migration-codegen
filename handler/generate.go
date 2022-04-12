package handler

import (
	"fmt"
	"github.com/enricodg/leaf-migration-codegen/helper/generateType"
	"github.com/enricodg/leaf-migration-codegen/helper/serviceParameter"
	"github.com/enricodg/leaf-migration-codegen/helper/templates"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func (h handler) Generate(migrationType string, name string) error {
	migrationsPath := fmt.Sprintf("migrations/%s", migrationType)
	if err := os.MkdirAll(migrationsPath, os.ModePerm); err != nil {
		return err
	}

	scriptsPath := fmt.Sprintf("scripts/%s", migrationType)
	if err := os.MkdirAll(scriptsPath, os.ModePerm); err != nil {
		return err
	}
	var templateFileName string
	if generateType.ServiceParameters == name {
		// NOTE: create migration file
		templateFileName = templates.ServiceParameterFileName
		migrationFileName := fmt.Sprintf("%s.%s", templateFileName, "go")
		h.log.StandardLogger().Infof("[%s] generating migration file for %s: %s", strings.ToUpper(migrationType),
			name, migrationFileName)
		if err := serviceParameter.GenerateMigration(migrationsPath, migrationType); err != nil {
			h.log.StandardLogger().Errorf("[%s] error generating %s: %+v", strings.ToUpper(migrationType), name, err.Error())
			return err
		}

		extension := "sql"
		migrateFileName := fmt.Sprintf("%s_migrate.%s", templateFileName, extension)
		rollbackFileName := fmt.Sprintf("%s_rollback.%s", templateFileName, extension)
		if err := serviceParameter.GenerateScript(scriptsPath, migrationType, migrateFileName, rollbackFileName); err != nil {
			h.log.StandardLogger().Errorf("[%s] error generating %s: %+v", strings.ToUpper(migrationType), name, err.Error())
			return err
		}

		return h.createMigrationInitialization(migrationType, name, templateFileName+".go", migrationsPath)
	}
	return errors.New(fmt.Sprintf("[%s] %s is not implemented", strings.ToUpper(migrationType), name))
}

func (h handler) createMigrationInitialization(migrationType string, name string, templateFileName string, migrationsPath string) error {
	// NOTE: create migrations initialization
	h.log.StandardLogger().Infof("[%s] initialize migration file for %s: %s",
		strings.ToUpper(migrationType), name, templateFileName)
	files := migration.LoadMigrations(migrationType)
	if err := helper.CreateInitialization(filepath.Join(migrationsPath, "initialize.go"),
		helper.InitializeRequestDTO{
			MigrationType: migrationType,
			IsMongo:       connection.IsMongo(migrationType),
			Versions:      files,
		}); err != nil {
		h.log.StandardLogger().Errorf("[%s] error initialize migration file for %s: %s: %s",
			strings.ToUpper(migrationType), name, templateFileName, err.Error())
		return err
	}
	h.log.StandardLogger().Infof("[%s] finish generating %s files", strings.ToUpper(migrationType), name)
	return nil
}
