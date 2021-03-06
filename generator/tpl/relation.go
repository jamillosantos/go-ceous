// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/relation.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Relation generates tpl/relation.gohtml
func Relation(env *models.Environment, relation *models.Relation) string {
	var _b strings.Builder
	RenderRelation(&_b, env, relation)
	return _b.String()
}

// RenderRelation render tpl/relation.gohtml
func RenderRelation(_buffer io.StringWriter, env *models.Environment, relation *models.Relation) {
	_buffer.WriteString("\n\ntype ")
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString(" struct {\n\t_runner ceous.DBRunner\n\tkeys []interface{}\n\trecords map[")
	_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignFieldType))
	_buffer.WriteString("][]")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, relation.LocalModelType)))
	_buffer.WriteString("\n}\n\nfunc New")
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString("(runner ceous.DBRunner) ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString(" {\n\treturn ")
	_buffer.WriteString(("&"))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString("{\n\t\t_runner: runner,\n\t\tkeys:    make([]interface{}, 0),\n\t\trecords: make(map[")
	_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignFieldType))
	_buffer.WriteString("][]")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, relation.LocalModelType)))
	_buffer.WriteString("),\n\t}\n}\n\nfunc (relation ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString(") Aggregate(record ceous.Record) error {\n\tugRecord, ok := record.(")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, relation.LocalModelType)))
	_buffer.WriteString(")\n\tif !ok {\n\t\treturn ceous.ErrInvalidRecordType\n\t}\n\tif rs, ok := relation.records[ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.LocalFieldRef))
	_buffer.WriteString("]; ok {\n\t\trelation.records[ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.LocalFieldRef))
	_buffer.WriteString("] = append(rs, ugRecord)\n\t\t// No need to add the key here, since its will be already in the `keys`.\n\t} else {\n\t\trelation.records[ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.LocalFieldRef))
	_buffer.WriteString("] = append(rs, ugRecord)\n\t\trelation.keys = append(relation.keys, ugRecord.")
	_buffer.WriteString(gorazor.HTMLEscape(relation.LocalFieldRef))
	_buffer.WriteString(")\n\t}\n\treturn nil\n}\n\nfunc (relation ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString(") Realize() error {\n\trecords, err := NewUserQuery(ceous.WithRunner(relation._runner)).Where(ceous.Eq(")
	_buffer.WriteString(gorazor.HTMLEscape(env.InputPkgCtx.Ref(env.OutputPkg, "Schema")))
	_buffer.WriteString(".")
	_buffer.WriteString(gorazor.HTMLEscape(relation.ForeignModelType))
	_buffer.WriteString(".ID, relation.keys)).All()\n\tif err != nil {\n\t\treturn err // TODO(jota): Shall this be wrapped into a custom error?\n\t}\n\tfor _, record := range records {\n\t\tmasterRecords, ok := relation.records[record.ID]\n\t\tif !ok {\n\t\t\treturn ceous.ErrInconsistentRelationResult\n\t\t}\n\t\tfor _, r := range masterRecords {\n\t\t\tr.Assign")
	_buffer.WriteString(gorazor.HTMLEscape(PascalCase(relation.LocalField)))
	_buffer.WriteString("(record)\n\t\t}\n\t}\n\treturn nil\n}\n\nfunc (q ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(relation.Query.FullName))
	_buffer.WriteString(") With")
	_buffer.WriteString(gorazor.HTMLEscape(PascalCase(relation.LocalField)))
	_buffer.WriteString("() ")
	_buffer.WriteString(gorazor.HTMLEscape(Pointer))
	_buffer.WriteString(gorazor.HTMLEscape(relation.Query.FullName))
	_buffer.WriteString(" {\n\tq.BaseQuery.Relations = append(q.BaseQuery.Relations, New")
	_buffer.WriteString(gorazor.HTMLEscape(relation.FullName))
	_buffer.WriteString("(q.BaseQuery.Runner))\n\treturn q\n}")

}
