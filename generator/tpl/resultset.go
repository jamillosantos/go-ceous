// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/resultset.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/naming"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Resultset generates tpl/resultset.gohtml
func Resultset(env *models.Environment, schema *models.Schema) string {
	var _b strings.Builder
	RenderResultset(&_b, env, schema)
	return _b.String()
}

// RenderResultset render tpl/resultset.gohtml
func RenderResultset(_buffer io.StringWriter, env *models.Environment, schema *models.Schema) {

	resultSetName := naming.CamelCase.Do(schema.Name) + "ResultSet"

	_buffer.WriteString("\n\ntype ")
	_buffer.WriteString(gorazor.HTMLEscape(resultSetName))
	_buffer.WriteString(" struct {\n\trs ceous.ResultSet\n\trecordScanner ceous.RecordScanner\n\tModel ")
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, schema.Name)))
	_buffer.WriteString("\n}\n\n// New")
	_buffer.WriteString(gorazor.HTMLEscape(naming.PascalCase.Do(resultSetName)))
	_buffer.WriteString(" create a new instance of the specialized\n// `ceous.ResultSet` for the model `")
	_buffer.WriteString(gorazor.HTMLEscape(schema.Name))
	_buffer.WriteString("`.\nfunc New")
	_buffer.WriteString(gorazor.HTMLEscape(naming.PascalCase.Do(resultSetName)))
	_buffer.WriteString("(rs ceous.ResultSet, err error) (")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(resultSetName))
	_buffer.WriteString(", error) {\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn ")
	_buffer.WriteString(("&"))
	_buffer.WriteString(gorazor.HTMLEscape(resultSetName))
	_buffer.WriteString("{\n\t\trs: rs,\n\t\trecordScanner: &ceous.DefaultRecordScanner,\n\t}, nil\n}\n\nfunc (rs ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(resultSetName))
	_buffer.WriteString(") Next() bool {\n\treturn rs.rs.Next()\n}\n\nfunc (rs ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(resultSetName))
	_buffer.WriteString(") Scan() error {\n\treturn rs.recordScanner.ScanRecord(rs.rs, &rs.Model)\n}\n\nfunc (rs ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(resultSetName))
	_buffer.WriteString(") Close() error {\n\terr := rs.rs.Close()\n\tif err != nil {\n\t\treturn err\n\t}\n\t// Disposes the instance...\n\trs.rs = nil\n\trs.recordScanner = nil\n\treturn nil\n}")

}
