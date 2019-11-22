// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/models.gohtml

package tpl

import (
	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Models generates tpl/models.gohtml
func Models(ctxPkg *generatorModels.Ctx, models []*generatorModels.Model, embeddeds []*generatorModels.Model, connections []*generatorModels.Connection) string {
	var _b strings.Builder
	RenderModels(&_b, ctxPkg, models, embeddeds, connections)
	return _b.String()
}

// RenderModels render tpl/models.gohtml
func RenderModels(_buffer io.StringWriter, ctxPkg *generatorModels.Ctx, models []*generatorModels.Model, embeddeds []*generatorModels.Model, connections []*generatorModels.Connection) {
	_buffer.WriteString("package ")
	_buffer.WriteString(gorazor.HTMLEscape(ctxPkg.InputPkg.Name))
	_buffer.WriteString("\n\nimport (\n\t\"github.com/jamillosantos/go-ceous\"\n\t\"github.com/pkg/errors\"")
	for _, pkg := range ctxPkg.ModelsImports.Imports {
		if pkg.Alias == "-" || pkg.Pkg.ImportPath == "." {
			continue
		}

		_buffer.WriteString(("\n"))

		_buffer.WriteString("\t")
		_buffer.WriteString(gorazor.HTMLEscape(pkg.Alias))
		_buffer.WriteString(" \"")
		_buffer.WriteString(gorazor.HTMLEscape(pkg.Pkg.ImportPath))
		_buffer.WriteString("\"")
	}
	_buffer.WriteString("\n)")
	for _, m := range models {
		RenderModel(_buffer, ctxPkg, m)
	}
	for _, m := range embeddeds {

		_buffer.WriteString(("\n\n"))

		RenderEmbedded(_buffer, m)
	}
	RenderSchema(_buffer, ctxPkg, models)

}
