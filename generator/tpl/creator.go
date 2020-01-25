// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/creator.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Creator generates tpl/creator.gohtml
func Creator(env *models.Environment) string {
	var _b strings.Builder
	RenderCreator(&_b, env)
	return _b.String()
}

// RenderCreator render tpl/creator.gohtml
func RenderCreator(_buffer io.StringWriter, env *models.Environment) {
	_buffer.WriteString("\n\ntype Creator struct {")
	for _, query := range env.Queries {
		_buffer.WriteString("\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(query.FullName))
		_buffer.WriteString("(options ...ceous.CeousOption) ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(query.FullName))
	}
	for _, store := range env.Stores {
		_buffer.WriteString("\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
		_buffer.WriteString("(options ...ceous.CeousOption) ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
	}
	_buffer.WriteString("\n}")

}
