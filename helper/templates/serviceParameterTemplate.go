package templates

const ServiceParameterFileName = "19700101000000_create_service_parameters"
const SqlServiceParameterMigrationTemplate = `package {{.MigrationType}}

import (
	"context"
	sqlConnection "github.com/enricodg/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/file"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type migration_19700101000000 struct {
	Log  leafLogger.Logger
	Conn sqlConnection.ORM
}

// NOTE: DO NOT CHANGE MIGRATION Version
func (m *migration_19700101000000) Version() uint64 {
	return uint64(19700101000000)
}

// NOTE: DO NOT CHANGE MIGRATION Name
func (m *migration_19700101000000) Name() string {
	return "create_service_parameters"
}

func (m *migration_19700101000000) Migrate() error {

	script, err := file.ReadToString("./scripts/{{.MigrationType}}/19700101000000_create_service_parameters_migrate.sql")
	if err != nil {
		return err
	}

	if err := m.Conn.Exec(context.Background(), script); err != nil {
		return err.Error()
	}

	return nil

}

func (m *migration_19700101000000) Rollback() error {
	script, err := file.ReadToString("./scripts/{{.MigrationType}}/19700101000000_create_service_parameters_rollback.sql")
	if err != nil {
		return err
	}

	if err := m.Conn.Exec(context.Background(), script); err != nil {
		return err.Error()
	}

	return nil
}
`

const MySqlServiceParameterMigrateTemplate = `CREATE TABLE IF NOT EXISTS service_parameters (
    id INT PRIMARY KEY AUTO_INCREMENT,
    variable VARCHAR(100) NOT NULL,
    value TEXT NOT NULL,
    description VARCHAR(255)NOT NULL,
    created_by varchar(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    is_deleted BOOLEAN DEFAULT false
);`

const PostgreSqlServiceParameterMigrateTemplate = `CREATE TABLE IF NOT EXISTS service_parameters (
    id SERIAL PRIMARY KEY,
    variable VARCHAR(100) NOT NULL,
    value TEXT NOT NULL,
    description VARCHAR(255)NOT NULL,
    created_by varchar(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100),
    updated_at TIMESTAMP NULL DEFAULT NULL,
    is_deleted BOOLEAN DEFAULT false
);`

const SqlServiceParameterRollbackTemplate = `DROP TABLE IF EXISTS service_parameters;`

const MongoServiceParameterMigrationTemplate = `package {{.MigrationType}}

import (
	"context"
	nosqlConnection "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type migration_19700101000000 struct {
	Log  leafLogger.Logger
	Conn nosqlConnection.Mongo
}

// NOTE: DO NOT CHANGE MIGRATION Version
func (m *migration_19700101000000) Version() uint64 {
	return uint64(19700101000000)
}

// NOTE: DO NOT CHANGE MIGRATION Name
func (m *migration_19700101000000) Name() string {
	return "create_service_parameters"
}

func (m *migration_19700101000000) Migrate() error {
	if err := m.Conn.DB().CreateCollection(context.Background(), "service_parameters"); err != nil {
		return err
	}

	if _, err := m.Conn.Indexes("service_parameters").CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{
					Key:   "variable",
					Value: bsonx.Int32(1),
				},
				{
					Key:   "is_deleted",
					Value: bsonx.Int32(1),
				},
			},
			Options: options.Index().
				SetName("variable_1_is_deleted_1").
				SetBackground(true),
		}); err != nil {
		return err
	}

	return nil
}

func (m *migration_19700101000000) Rollback() error {
	if err := m.Conn.DB().Collection("service_parameters").Drop(context.Background()); err != nil {
		return err
	}

	return nil
}
`
