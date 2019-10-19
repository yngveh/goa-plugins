package logrus

import (
	"testing"

	"goa.design/goa/v3/codegen"
)

func TestGenerate(t *testing.T) {

	t.Logf("Testing")

	files := []*codegen.File{
		{
			Path: "/mypath",
			SectionTemplates: []*codegen.SectionTemplate{
				{
					Name: "server-main-logger",
					Source: `        {{ comment "Setup logger. Replace logger with your own log package of choice." }}
        var (
                logger *log.Logger
        )
        {
                logger = log.New(os.Stderr, "[{{ .APIPkg }}] ", log.Ltime)
        }
`,
				},
			},
		},
	}

	g, err := Generate("my.pkg/myuser/myproject", nil, files)

	for _, j := range g {
		debugf("%s", j.Path)
	}

	debugf("Result: %v, %v", g, err)
}
