package logrus

import (
	"strings"
	"testing"

	"goa.design/goa/v3/codegen"
)

var (
	files = []*codegen.File{
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
)

func TestGenerate(t *testing.T) {

	g, err := Generate("my.pkg/myuser/myproject", nil, files)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	var mainFile *codegen.File
	for _, f := range g {
		if f.Path == "/mypath" {
			mainFile = f
			break
		}
	}

	if mainFile == nil || len(mainFile.SectionTemplates) < 1 {
		t.Fatalf("no section templates found")
	}

	st := mainFile.Section(sectionTemplateServerMainLogger)
	if len(st) < 1 {
		t.Fatalf("no section template %s found", sectionTemplateServerMainLogger)
	}

	if !strings.Contains(st[0].Source, "logrus") {
		t.Fatalf("source not containing logrus")
	}

}
