package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewCustomEventClass() data.ClassStmt {
	return &CustomEventClass{
		source:      nil,
		construct:   &CustomEventConstructMethod{source: nil},
		cancel:      &CustomEventCancelMethod{source: nil},
		isCancelled: &CustomEventIsCancelledMethod{source: nil},
		toJSON:      &CustomEventToJSONMethod{source: nil},
	}
}

func NewCustomEventClassFrom(source *applicationsrc.CustomEvent) data.ClassStmt {
	return &CustomEventClass{
		source:      source,
		construct:   &CustomEventConstructMethod{source: source},
		cancel:      &CustomEventCancelMethod{source: source},
		isCancelled: &CustomEventIsCancelledMethod{source: source},
		toJSON:      &CustomEventToJSONMethod{source: source},
	}
}

type CustomEventClass struct {
	node.Node
	source      *applicationsrc.CustomEvent
	construct   data.Method
	cancel      data.Method
	isCancelled data.Method
	toJSON      data.Method
}

func (s *CustomEventClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewCustomEventClassFrom(&applicationsrc.CustomEvent{}), ctx.CreateBaseContext()), nil
}
func (s *CustomEventClass) GetName() string         { return "wails\\application\\CustomEvent" }
func (s *CustomEventClass) GetExtend() *string      { return nil }
func (s *CustomEventClass) GetImplements() []string { return nil }
func (s *CustomEventClass) AsString() string        { return "CustomEvent{}" }
func (s *CustomEventClass) GetSource() any          { return s.source }
func (s *CustomEventClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Name":
		return node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(s.source.Name)), true
	case "Data":
		return node.NewProperty(nil, "Data", "public", true, data.NewAnyValue(s.source.Data)), true
	case "Sender":
		return node.NewProperty(nil, "Sender", "public", true, data.NewAnyValue(s.source.Sender)), true
	}
	return nil, false
}

func (s *CustomEventClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Name":   node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(nil)),
		"Data":   node.NewProperty(nil, "Data", "public", true, data.NewAnyValue(nil)),
		"Sender": node.NewProperty(nil, "Sender", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *CustomEventClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Name":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Name = string(sv.AsString())
		}
	case "Data":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Data = v.Value.(interface{})
		}
	case "Sender":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Sender = string(sv.AsString())
		}
	}
}

func (s *CustomEventClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "cancel":
		return s.cancel, true
	case "isCancelled":
		return s.isCancelled, true
	case "toJSON":
		return s.toJSON, true
	}
	return nil, false
}

func (s *CustomEventClass) GetMethods() []data.Method {
	return []data.Method{
		s.cancel,
		s.isCancelled,
		s.toJSON,
	}
}

func (s *CustomEventClass) GetConstruct() data.Method { return s.construct }

type CustomEventConstructMethod struct {
	source *applicationsrc.CustomEvent
}

func (h *CustomEventConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *CustomEventConstructMethod) GetName() string               { return "construct" }
func (h *CustomEventConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *CustomEventConstructMethod) GetIsStatic() bool             { return false }
func (h *CustomEventConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *CustomEventConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *CustomEventConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
