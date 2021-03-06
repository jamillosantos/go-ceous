// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/transaction.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Transaction generates tpl/transaction.gohtml
func Transaction(env *models.Environment) string {
	var _b strings.Builder
	RenderTransaction(&_b, env)
	return _b.String()
}

// RenderTransaction render tpl/transaction.gohtml
func RenderTransaction(_buffer io.StringWriter, env *models.Environment) {
	_buffer.WriteString("\n\ntype Transaction struct {\n\t*ceous.BaseTxRunner\n}\n\nfunc NewTransaction(tx *ceous.BaseTxRunner) *Transaction {\n\treturn &Transaction{\n\t\tBaseTxRunner: tx,\n\t}\n}")
	for _, query := range env.Queries {
		_buffer.WriteString("\n// ")
		_buffer.WriteString(gorazor.HTMLEscape(query.FullName))
		_buffer.WriteString(" creates a new query from a transaction.\nfunc (c *Transaction) ")
		_buffer.WriteString(gorazor.HTMLEscape(query.FullName))
		_buffer.WriteString("(options ...ceous.CeousOption) ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(query.FullName))
		_buffer.WriteString(" {\n\treturn New")
		_buffer.WriteString(gorazor.HTMLEscape(query.FullName))
		_buffer.WriteString("(append(options, ceous.WithRunner(c))...)\n}")
	}
	for _, store := range env.Stores {
		_buffer.WriteString("\n// ")
		_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
		_buffer.WriteString(" creates a new store from a transaction.\nfunc (c *Transaction) ")
		_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
		_buffer.WriteString("(options ...ceous.CeousOption) ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
		_buffer.WriteString(" {\n\treturn New")
		_buffer.WriteString(gorazor.HTMLEscape(store.FullName))
		_buffer.WriteString("(append(options, ceous.WithRunner(c))...)\n}")
	}

}
