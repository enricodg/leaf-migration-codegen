package templates

const MainTemplate = `package main 

import (
	migration "github.com/paulusrobin/leaf-utilities/leafMigration"
)

func main() {
	migration.New().
		// WithMySql(mysql.InitializeMigrations).
		// WithPostgre(postgre.InitializeMigrations).
		// WithMongo(mongo.InitializeMigrations).
		Run()
}`

const GoModTemplate = `module {{.ProjectName}}

go 1.18

require github.com/paulusrobin/leaf-utilities/leafMigration v0.0.0-20220413034912-5e197528c618 // indirect
`

const EnvExampleTemplate = `LOG_LEVEL=INFO
LOG_FORMATTER=TEXT

# MySQL connection
MY_SQL_ADDRESS=
MY_SQL_USERNAME=
MY_SQL_PASSWORD=
MY_SQL_DB_NAME=
MY_SQL_MAX_IDLE_CONNECTION=1
MY_SQL_MAX_OPEN_CONNECTION=1
MY_SQL_MAX_LIFETIME_CONNECTION=30s
MY_SQL_LOG_MODE=false

# PostgreSQL connection
POSTGRE_SQL_ADDRESS=
POSTGRE_SQL_USERNAME=
POSTGRE_SQL_PASSWORD=
POSTGRE_SQL_DB_NAME=
POSTGRE_SQL_MAX_IDLE_CONNECTION=1
POSTGRE_SQL_MAX_OPEN_CONNECTION=1
POSTGRE_SQL_MAX_LIFETIME_CONNECTION=30s
POSTGRE_SQL_LOG_MODE=false

# MongoDB connection
MONGO_DB_URI=
MONGO_DB_DATABASE=
`

const GitIgnoreTemplate = `.idea
vendor
.env
`
