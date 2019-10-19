package logrus

import (
	"fmt"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

const (
	name         = "logrus"
	cmd          = "example"
	importLogrus = "github.com/sirupsen/logrus"

	logrusTemplate = `        {{ comment "Setup logrus logger." }}
  var (
		logger *log.Logger
	)
	{
		logrus.SetFormatter( &logrus.JSONFormatter{} )
		logger = log.New(logrus.StandardLogger().Writer(), "", 0)
	}
`
)

var (
	sectionsFunc = map[string]func(f *codegen.File, t *codegen.SectionTemplate){
		"server-main-logger": serverMainLog,
	}
)

func init() {
	fmt.Println("[Init]")
	codegen.RegisterPlugin(name, cmd, nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {

	fmt.Printf("[Genpkg] %s", genpkg)

	for _, f := range files {
		fmt.Printf("[File] %s\n", f.Path)

		for _, a := range f.SectionTemplates {
			fmt.Printf("  [SectionTemplate] %s\n", a.Name)

			if sf, ok := sectionsFunc[a.Name]; ok {
				sf(f, a)
			}
		}
	}

	return files, nil
}

func serverMainLog(f *codegen.File, t *codegen.SectionTemplate) {
	fmt.Printf("Called %s \n", t.Name)
	fmt.Printf("Source:\n%s\n", t.Source)
	fmt.Printf("Data:%v\n", t.Data)
	codegen.AddImport(f.SectionTemplates[0], &codegen.ImportSpec{Path: importLogrus})
	t.Source = logrusTemplate
}
