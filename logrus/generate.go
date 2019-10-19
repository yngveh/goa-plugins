package logrus

import (
	"fmt"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

const (
	name           = "logrus"
	cmd            = "example"
	logrusTemplate = `
  var (
		logger *log.Logger
	)
	{
		logrus.SetFormatter( &logrus.JSONFormatter{} )
		logger = log.New(logrus.StandardLogger().Writer(), "", 0)
	}`
)

var (
	sectionsFunc = map[string]func(t *codegen.SectionTemplate){
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

			if f, ok := sectionsFunc[a.Name]; ok {
				f(a)
			}
		}
	}

	return files, nil
}

func serverMainLog(t *codegen.SectionTemplate) {
	fmt.Printf("Called %s \n", t.Name)
	fmt.Printf("Source:\n%s\n", t.Source)
	fmt.Printf("Data:%v\n", t.Data)
	t.Source = logrusTemplate
}
