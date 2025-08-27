package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewButtonClass() data.ClassStmt {
	return &ButtonClass{
		source:       nil,
		construct:    &ButtonConstructMethod{source: nil},
		onClick:      &ButtonOnClickMethod{source: nil},
		setAsCancel:  &ButtonSetAsCancelMethod{source: nil},
		setAsDefault: &ButtonSetAsDefaultMethod{source: nil},
	}
}

func NewButtonClassFrom(source *applicationsrc.Button) data.ClassStmt {
	return &ButtonClass{
		source:       source,
		construct:    &ButtonConstructMethod{source: source},
		onClick:      &ButtonOnClickMethod{source: source},
		setAsCancel:  &ButtonSetAsCancelMethod{source: source},
		setAsDefault: &ButtonSetAsDefaultMethod{source: source},
	}
}

type ButtonClass struct {
	node.Node
	source       *applicationsrc.Button
	construct    data.Method
	onClick      data.Method
	setAsCancel  data.Method
	setAsDefault data.Method
}

func (s *ButtonClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewButtonClassFrom(&applicationsrc.Button{}), ctx.CreateBaseContext()), nil
}
func (s *ButtonClass) GetName() string         { return "wails\\application\\Button" }
func (s *ButtonClass) GetExtend() *string      { return nil }
func (s *ButtonClass) GetImplements() []string { return nil }
func (s *ButtonClass) AsString() string        { return "Button{}" }
func (s *ButtonClass) GetSource() any          { return s.source }
func (s *ButtonClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Label":
		return node.NewProperty(nil, "Label", "public", true, data.NewAnyValue(s.source.Label)), true
	case "IsCancel":
		return node.NewProperty(nil, "IsCancel", "public", true, data.NewAnyValue(s.source.IsCancel)), true
	case "IsDefault":
		return node.NewProperty(nil, "IsDefault", "public", true, data.NewAnyValue(s.source.IsDefault)), true
	case "Callback":
		return node.NewProperty(nil, "Callback", "public", true, data.NewAnyValue(s.source.Callback)), true
	}
	return nil, false
}

func (s *ButtonClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Label":     node.NewProperty(nil, "Label", "public", true, data.NewAnyValue(nil)),
		"IsCancel":  node.NewProperty(nil, "IsCancel", "public", true, data.NewAnyValue(nil)),
		"IsDefault": node.NewProperty(nil, "IsDefault", "public", true, data.NewAnyValue(nil)),
		"Callback":  node.NewProperty(nil, "Callback", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *ButtonClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Label":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Label = string(sv.AsString())
		}
	case "IsCancel":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.IsCancel = bool(x)
			}
		}
	case "IsDefault":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.IsDefault = bool(x)
			}
		}
	case "Callback":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Callback = v.Value.(func())
		}
	}
}

func (s *ButtonClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "onClick":
		return s.onClick, true
	case "setAsCancel":
		return s.setAsCancel, true
	case "setAsDefault":
		return s.setAsDefault, true
	}
	return nil, false
}

func (s *ButtonClass) GetMethods() []data.Method {
	return []data.Method{
		s.onClick,
		s.setAsCancel,
		s.setAsDefault,
	}
}

func (s *ButtonClass) GetConstruct() data.Method { return s.construct }

type ButtonConstructMethod struct {
	source *applicationsrc.Button
}

func (h *ButtonConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ButtonConstructMethod) GetName() string               { return "construct" }
func (h *ButtonConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ButtonConstructMethod) GetIsStatic() bool             { return false }
func (h *ButtonConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ButtonConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ButtonConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
