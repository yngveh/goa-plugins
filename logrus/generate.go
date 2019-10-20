package logrus

import (
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

const (
	name                            = "logrus"
	cmd                             = "example"
	importLogrus                    = "github.com/sirupsen/logrus"
	sectionTemplateServerMainLogger = "server-main-logger"

	logrusTemplate = `
        {{ comment "Setup logrus logger." }}
  var (
		logger *log.Logger
	)
	{
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)

		logger = log.New(logrus.StandardLogger().Writer(), "", 0)
	}
`
)

var (
	sectionsFunc = map[string]func(f *codegen.File, t *codegen.SectionTemplate){
		sectionTemplateServerMainLogger: serverMainLog,
	}
)

func init() {
	codegen.RegisterPlugin(name, cmd, nil, Generate)
}

func Generate(_ string, _ []eval.Root, files []*codegen.File) ([]*codegen.File, error) {

	for _, f := range files {
		for _, a := range f.SectionTemplates {
			if sf, ok := sectionsFunc[a.Name]; ok {
				sf(f, a)
			}
		}
	}

	return files, nil
}

func serverMainLog(f *codegen.File, t *codegen.SectionTemplate) {
	codegen.AddImport(f.SectionTemplates[0], &codegen.ImportSpec{Path: importLogrus})
	t.Source = logrusTemplate
}
