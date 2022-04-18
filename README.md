# Leaf Migration Code Generator
This repository contains code generator for [leaf migration](https://github.com/paulusrobin/leaf-utilities/tree/main/leafMigration).

## Prerequisites

1. Go 1.18

## Installation
Simply download it from [here](https://github.com/enricodg/leaf-migration-codegen/releases/download/v1.0.0/leaf-migration-codegen) or do the following steps:
1. Clone this repository
```shell
$ git clone https://github.com/enricodg/leaf-migration-codegen
```
2. Go to the directory, and install dependencies and binaries
```shell
$ cd leaf-migration-codegen
$ go mod tidy
$ go install
```
3. You should be able to use the CLI by using `leaf-migration-codegen`
```shell
$ leaf-migration-codegen
NAME:
   Leaf migration codegen - Supporting leaf framework to initialize migration project

USAGE:
   leaf-migration-codegen command [command options] [arguments...]

VERSION:
   v1.0.0

DESCRIPTION:
   CLI Leaf migration code generator

COMMANDS:
   init      init --project <project URL>
   generate  generate --types <type> --name <name>
   new       new --types <type> --name <name>
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Usage

### 1. Initialize Project
Create your directory / git repository and pass the value to init command, example:
```shell
$ leaf-migration-codegen init --project github.com/enricodg/leaf-migration-example
```

### 2. Generate Files
Currently it only supports service parameter generation. Service parameter is used to change service configuration on the fly without restarting the service.
```shell
$ leaf-migration-codegen generate --types <mysql | postgre | mongo> --name <service_parameters>
```

### 3. Generate Migration
Leaf migration currently supports 3 databases (MongoDB, PostgreSQL, and MySQL), in order to specify which type of database migration you want to create please use `--types` flag. To name the migration file, use the `--name` flag. 
```shell
$ leaf-migration-codegen new --types <mysql | postgre | mongo> --name what_to_add_or_remove
```

### 4. Migrate
After you've done all the steps above, you can directly use the migration files to migrate, rollback, and check the version of your migration. If you want to migrate specific version only, use the `--specific` flag or otherwise it will migrate all available migrations. 

To migrate, use the following command:
```shell
$ go run main.go migrate [--types <mysql | postgre | mongo>] [--specific] [--verbose]
```

### 5. Rollback
If you don't specify which version to rollback (`--specific` flag) it will rollback the migration until the version that you mentioned.

To rollback, use the following command:
```shell
$ go run main.go rollback --types [mysql | postgre | mongo] --version <version number> [--specific] [--verbose]
```

### 6. Check Migration
To check, use the following command:
```shell
$ go run main.go check [--types <mysql | postgre | mongo> [--version <version number>]
```

---
Feel free to contribute to the project by creating issue or pull request. :)
