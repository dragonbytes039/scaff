package fst

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed templates/*
var AllTemplates embed.FS

func SpreadTemplate(template string, projectName string) error {

	contentTemplate, err := fs.Sub(AllTemplates, "templates/"+template)

	i := 0
	projectNameCopy := projectName
	for DirExist(projectNameCopy) {
		projectNameCopy = fmt.Sprintf("%s%d", projectName, i)
		i++
	}
	projectName = projectNameCopy

	if err != nil {
		return fmt.Errorf("no se encontró el template")
	}

	return fs.WalkDir(contentTemplate, ".", func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		targetPath := filepath.Join("./"+projectName, path)

		if d.IsDir() {
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("error creando carpeta %s: %w", targetPath, err)
			}
		} else {
			datos, err := fs.ReadFile(contentTemplate, path)

			if err != nil {
				return fmt.Errorf("error leyendo archivo embed %s: %w", path, err)
			}

			if path == "go.mod.txt" {
				targetPath = filepath.Join("./"+projectName, "go.mod")
			}

			if err := os.WriteFile(targetPath, datos, 0644); err != nil {
				return fmt.Errorf("error escribiendo archivo en disco %s: %w", targetPath, err)
			}
		}

		return nil
	})

}
