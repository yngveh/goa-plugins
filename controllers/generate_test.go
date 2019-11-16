package controllers

import (
	"testing"

	"goa.design/goa/v3/codegen"
)

var (
	file = &codegen.File{
		SectionTemplates: []*codegen.SectionTemplate{
			{
				Data: map[string]interface{}{
					"Imports": []*codegen.ImportSpec{
						{
							Path: "github.com/test/me",
						},
					},
				},
			},
		},
	}
)

func TestChangePath(t *testing.T) {

	changeImportForControllers("github.com/test/me/gen", file)

	data := file.SectionTemplates[0].Data.(map[string]interface{})
	imports := data["Imports"].([]*codegen.ImportSpec)

	actual := imports[0].Path
	expected := "github.com/test/me/controllers"

	if actual != expected {
		t.Fatalf("Expected %s, actual %s", expected, actual)
	}
}
