package helper

import (
	"github.com/enricodg/leaf-migration-codegen/helper/templates"
	"os"
	"text/template"
)

type (
	InitializeProjectRequestDTO struct {
		ProjectName string
	}
	CreateFileDTO struct {
		MigrationType string
		Template      string
	}
)

func Initialize(data InitializeProjectRequestDTO) error {
	main, err := template.New("main.go").Parse(templates.MainTemplate)
	if err != nil {
		return err
	}

	mainFile, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer mainFile.Close()

	err = main.Execute(mainFile, nil)
	if err != nil {
		return err
	}

	goMod, err := template.New("go.mod").Parse(templates.GoModTemplate)
	if err != nil {
		return err
	}

	goModFile, err := os.Create("go.mod")
	if err != nil {
		return err
	}
	defer goModFile.Close()

	err = goMod.Execute(goModFile, data)
	if err != nil {
		return err
	}

	env, err := template.New(".env.example").Parse(templates.EnvExampleTemplate)
	if err != nil {
		return err
	}

	envFile, err := os.Create(".env.example")
	if err != nil {
		return err
	}
	defer envFile.Close()

	err = env.Execute(envFile, nil)
	if err != nil {
		return err
	}

	gitIgnore, err := template.New(".gitignore").Parse(templates.GitIgnoreTemplate)
	if err != nil {
		return err
	}

	gitIgnoreFile, err := os.Create(".gitignore")
	if err != nil {
		return err
	}
	defer envFile.Close()

	err = gitIgnore.Execute(gitIgnoreFile, nil)
	if err != nil {
		return err
	}

	return nil
}

func CreateFile(outputPath string, data CreateFileDTO) error {
	tmpl, err := template.New(outputPath).Parse(data.Template)
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}
