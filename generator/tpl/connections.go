// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: tpl/connections.gohtml

package tpl

import (
	. "github.com/jamillosantos/go-ceous/generator/helpers"
	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Connections generates tpl/connections.gohtml
func Connections(connections []*generatorModels.Connection, models []*generatorModels.Model) string {
	var _b strings.Builder
	RenderConnections(&_b, connections, models)
	return _b.String()
}

// RenderConnections render tpl/connections.gohtml
func RenderConnections(_buffer io.StringWriter, connections []*generatorModels.Connection, models []*generatorModels.Model) {
	_buffer.WriteString("\n\ntype Connection interface {\n\tceous.Connection")
	for _, model := range models {
		_buffer.WriteString("\n\t// ")
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString("Query creates a new query related with the connection set.\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString("Query(options ...ceous.CeousOption) ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
		_buffer.WriteString("\n\t// ")
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString("Store creates a new store related with the connection set.\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(model.Name))
		_buffer.WriteString("Store(options ...ceous.CeousOption) ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(model.StoreName()))
	}
	_buffer.WriteString("\n}")
	for _, conn := range connections {
		_buffer.WriteString("\ntype ")
		_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
		_buffer.WriteString("  struct {\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString("ceous.BaseConnection\n}\n\n\t")
		for _, model := range models {
			_buffer.WriteString("\n// ")
			_buffer.WriteString(gorazor.HTMLEscape(model.Name))
			_buffer.WriteString("Query creates a new query related with the connection ")
			_buffer.WriteString(gorazor.HTMLEscape(conn.Name))
			_buffer.WriteString(" set.\nfunc (c ")
			_buffer.WriteString(gorazor.HTMLEscape(Pointer))
			_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
			_buffer.WriteString(") ")
			_buffer.WriteString(gorazor.HTMLEscape(model.Name))
			_buffer.WriteString("Query(options ...ceous.CeousOption) ")
			_buffer.WriteString(gorazor.HTMLEscape(Pointer))
			_buffer.WriteString(gorazor.HTMLEscape(model.QueryName()))
			_buffer.WriteString(" {\n\treturn New")
			_buffer.WriteString(gorazor.HTMLEscape(model.Name))
			_buffer.WriteString("Query(append(options, ceous.WithConn(c))...)\n}\n\n// ")
			_buffer.WriteString(gorazor.HTMLEscape(model.Name))
			_buffer.WriteString("Store creates a new store related with the connection ")
			_buffer.WriteString(gorazor.HTMLEscape(conn.Name))
			_buffer.WriteString(" set.\nfunc (c ")
			_buffer.WriteString(gorazor.HTMLEscape(Pointer))
			_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
			_buffer.WriteString(") ")
			_buffer.WriteString(gorazor.HTMLEscape(model.Name))
			_buffer.WriteString("Store(options ...ceous.CeousOption) ")
			_buffer.WriteString(gorazor.HTMLEscape(Pointer))
			_buffer.WriteString(gorazor.HTMLEscape(model.StoreName()))
			_buffer.WriteString(" {\n\treturn New")
			_buffer.WriteString(gorazor.HTMLEscape(model.Name))
			_buffer.WriteString("Store(append(options, ceous.WithConn(c))...)\n}\n\t")
		}
		_buffer.WriteString("\n\n// Begin creates a new transaction with ")
		_buffer.WriteString(gorazor.HTMLEscape(conn.Name))
		_buffer.WriteString(" set.\nfunc (c ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
		_buffer.WriteString(") Begin() (*Transaction, error) {\n\ttx, err := c.BaseConnection.Begin()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn NewTransaction(tx), nil\n}\n\n// BeginTx creates a new transaction with extended config params with the\n// connection ")
		_buffer.WriteString(gorazor.HTMLEscape(conn.Name))
		_buffer.WriteString(" set.\nfunc (c ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
		_buffer.WriteString(") BeginTx(ctx context.Context, opts *sql.TxOptions) (*Transaction, error) {\n\ttx, err := c.BaseConnection.BeginTx(ctx, opts)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn NewTransaction(tx), nil\n}")
	}
	_buffer.WriteString("\n\nvar (")
	for _, conn := range connections {
		_buffer.WriteString("\n\t// ")
		_buffer.WriteString(gorazor.HTMLEscape(conn.Name))
		_buffer.WriteString(" is a database connection reference.\n\t")
		_buffer.WriteString(gorazor.HTMLEscape((conn.Name)))
		_buffer.WriteString(" ")
		_buffer.WriteString(gorazor.HTMLEscape(Pointer))
		_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
	}
	_buffer.WriteString("\n)")
	for _, conn := range connections {
		_buffer.WriteString("\n// Init")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(conn.Name)))
		_buffer.WriteString(" initializes the connection `")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(conn.Name)))
		_buffer.WriteString(":`.\nfunc Init")
		_buffer.WriteString(gorazor.HTMLEscape(PascalCase(conn.Name)))
		_buffer.WriteString("(db *sql.DB) {\n\t")
		_buffer.WriteString(gorazor.HTMLEscape(conn.Name))
		_buffer.WriteString(" = ")
		_buffer.WriteString(("&"))
		_buffer.WriteString(gorazor.HTMLEscape(conn.ConnectionName()))
		_buffer.WriteString("{\n\t\tBaseConnection: ceous.NewConnection(db),\n\t}\n}")
	}

}