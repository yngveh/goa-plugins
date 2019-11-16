package logrus

import (
	"os"
	"path"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

const (
	name           = "controllers"
	cmd            = "example"
	controllersDir = "controllers"
)

func init() {
	codegen.RegisterPlugin(name, cmd, Prepare, Generate)
}

func Generate(genPath string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {

	controllersPath := path.Join(genPath, "..", controllersDir)
	err := os.MkdirAll(controllersPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	for _, f := range files {

		changeImportForControllers(genPath, f)

		d, _ := path.Split(f.Path)
		if d == "" {
			f.Path = path.Join(controllersDir, f.Path)
		}
	}

	return files, nil
}

func changeImportForControllers(genPath string, f *codegen.File) {

	for _, sectionTemplate := range f.SectionTemplates {
		if data, ok := sectionTemplate.Data.(map[string]interface{}); ok {
			if imports, ok := data["Imports"]; ok {
				if importSpecs, ok := imports.([]*codegen.ImportSpec); ok {
					for _, importSpec := range importSpecs {
						if importSpec.Path+"/gen" == genPath {
							importSpec.Path += "/" + controllersDir
						}
					}
				}
			}
		}
	}
}

func Prepare(genpkg string, roots []eval.Root) error {
	return nil
}
