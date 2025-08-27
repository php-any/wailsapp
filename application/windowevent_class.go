package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWindowEventClass() data.ClassStmt {
	return &WindowEventClass{
		source:    nil,
		construct: &WindowEventConstructMethod{source: nil},
		cancel:    &WindowEventCancelMethod{source: nil},
		context:   &WindowEventContextMethod{source: nil},
	}
}

func NewWindowEventClassFrom(source *applicationsrc.WindowEvent) data.ClassStmt {
	return &WindowEventClass{
		source:    source,
		construct: &WindowEventConstructMethod{source: source},
		cancel:    &WindowEventCancelMethod{source: source},
		context:   &WindowEventContextMethod{source: source},
	}
}

type WindowEventClass struct {
	node.Node
	source    *applicationsrc.WindowEvent
	construct data.Method
	cancel    data.Method
	context   data.Method
}

func (s *WindowEventClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWindowEventClassFrom(&applicationsrc.WindowEvent{}), ctx.CreateBaseContext()), nil
}
func (s *WindowEventClass) GetName() string         { return "wails\\application\\WindowEvent" }
func (s *WindowEventClass) GetExtend() *string      { return nil }
func (s *WindowEventClass) GetImplements() []string { return nil }
func (s *WindowEventClass) AsString() string        { return "WindowEvent{}" }
func (s *WindowEventClass) GetSource() any          { return s.source }
func (s *WindowEventClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Cancelled":
		return node.NewProperty(nil, "Cancelled", "public", true, data.NewAnyValue(s.source.Cancelled)), true
	}
	return nil, false
}

func (s *WindowEventClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Cancelled": node.NewProperty(nil, "Cancelled", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *WindowEventClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Cancelled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.Cancelled = bool(x)
			}
		}
	}
}

func (s *WindowEventClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "cancel":
		return s.cancel, true
	case "context":
		return s.context, true
	}
	return nil, false
}

func (s *WindowEventClass) GetMethods() []data.Method {
	return []data.Method{
		s.cancel,
		s.context,
	}
}

func (s *WindowEventClass) GetConstruct() data.Method { return s.construct }

type WindowEventConstructMethod struct {
	source *applicationsrc.WindowEvent
}

func (h *WindowEventConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowEventConstructMethod) GetName() string               { return "construct" }
func (h *WindowEventConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowEventConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowEventConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowEventConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowEventConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
