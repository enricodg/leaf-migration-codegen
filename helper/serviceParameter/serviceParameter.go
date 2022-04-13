package serviceParameter

import (
	"fmt"
	"github.com/enricodg/leaf-migration-codegen/helper"
	"github.com/enricodg/leaf-migration-codegen/helper/templates"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"
)

func GenerateMigration(outputPath string, migrationType string) error {
	switch migrationType {
	case connection.MySQL, connection.Postgre:
		if err := helper.CreateFile(filepath.Join(outputPath, fmt.Sprintf("%s.%s", templates.ServiceParameterFileName, "go")),
			helper.CreateFileDTO{
				MigrationType: migrationType,
				Template:      templates.SqlServiceParameterMigrationTemplate,
			}); err != nil {
		}
		break
	case connection.Mongo:
		if err := helper.CreateFile(filepath.Join(outputPath, fmt.Sprintf("%s.%s", templates.ServiceParameterFileName, "go")),
			helper.CreateFileDTO{
				MigrationType: migrationType,
				Template:      templates.MongoServiceParameterMigrationTemplate,
			}); err != nil {
		}
		break
	default:
		return errors.New(fmt.Sprintf("[%s] service_parameters not implemented", strings.ToUpper(migrationType)))
	}

	return nil
}

func GenerateScript(outputPath, migrationType, migrateFileName, rollbackFileName string) error {
	switch migrationType {
	case connection.MySQL:
		return generateSqlScript(outputPath, migrateFileName, rollbackFileName,
			templates.MySqlServiceParameterMigrateTemplate, templates.SqlServiceParameterRollbackTemplate)
	case connection.Postgre:
		return generateSqlScript(outputPath, migrateFileName, rollbackFileName,
			templates.PostgreSqlServiceParameterMigrateTemplate, templates.SqlServiceParameterRollbackTemplate)
	}

	return nil
}

func generateSqlScript(outputPath, migrateFileName, rollbackFileName, migrateTemplate, rollbackTemplate string) error {
	if err := helper.CreateFile(filepath.Join(outputPath, migrateFileName),
		helper.CreateFileDTO{
			Template: migrateTemplate,
		}); err != nil {
	}
	if err := helper.CreateFile(filepath.Join(outputPath, rollbackFileName),
		helper.CreateFileDTO{
			Template: rollbackTemplate,
		}); err != nil {
	}

	return nil
}
