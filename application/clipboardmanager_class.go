package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewClipboardManagerClass() data.ClassStmt {
	return &ClipboardManagerClass{
		source:    nil,
		construct: &ClipboardManagerConstructMethod{source: nil},
		setText:   &ClipboardManagerSetTextMethod{source: nil},
		text:      &ClipboardManagerTextMethod{source: nil},
	}
}

func NewClipboardManagerClassFrom(source *applicationsrc.ClipboardManager) data.ClassStmt {
	return &ClipboardManagerClass{
		source:    source,
		construct: &ClipboardManagerConstructMethod{source: source},
		setText:   &ClipboardManagerSetTextMethod{source: source},
		text:      &ClipboardManagerTextMethod{source: source},
	}
}

type ClipboardManagerClass struct {
	node.Node
	source    *applicationsrc.ClipboardManager
	construct data.Method
	setText   data.Method
	text      data.Method
}

func (s *ClipboardManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewClipboardManagerClassFrom(&applicationsrc.ClipboardManager{}), ctx.CreateBaseContext()), nil
}
func (s *ClipboardManagerClass) GetName() string                            { return "wails\\application\\ClipboardManager" }
func (s *ClipboardManagerClass) GetExtend() *string                         { return nil }
func (s *ClipboardManagerClass) GetImplements() []string                    { return nil }
func (s *ClipboardManagerClass) AsString() string                           { return "ClipboardManager{}" }
func (s *ClipboardManagerClass) GetSource() any                             { return s.source }
func (s *ClipboardManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *ClipboardManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *ClipboardManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "setText":
		return s.setText, true
	case "text":
		return s.text, true
	}
	return nil, false
}

func (s *ClipboardManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.setText,
		s.text,
	}
}

func (s *ClipboardManagerClass) GetConstruct() data.Method { return s.construct }

type ClipboardManagerConstructMethod struct {
	source *applicationsrc.ClipboardManager
}

func (h *ClipboardManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ClipboardManagerConstructMethod) GetName() string               { return "construct" }
func (h *ClipboardManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ClipboardManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *ClipboardManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ClipboardManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ClipboardManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
