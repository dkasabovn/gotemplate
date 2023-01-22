//go:build mage

package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	ProjectName = "template"
)

type (
	Run      mg.Namespace
	Test     mg.Namespace
	Build    mg.Namespace
	Util     mg.Namespace
	Generate mg.Namespace
)

func genURL(name string) string {
	return fmt.Sprintf("%s/%ssvc", ProjectName, name)
}

func (Generate) SQL() error {
	return sh.RunV("sqlc", "generate", "-f", "db/sqlc.yaml")
}

func (Run) Service(name string) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	servicePath := fmt.Sprintf("app/cmd/%s/main.go", name)
	return sh.RunV("go", "run", servicePath)
}

func buildRoot() error {
	return sh.RunV("docker", "build", "--no-cache", "--file", "Dockerfile.root", "--tag", genURL("root"), ".")
}

func (Build) Service(name string) error {
	mg.Deps(buildRoot)
	url := genURL(name)
	return sh.RunV("docker", "build", "--no-cache", "--tag", url, "--file", fmt.Sprintf("app/cmd/%s/Dockerfile", name), ".")
}

func (Build) All() error {
	mg.Deps(
		buildRoot,
	)
	return nil
}
