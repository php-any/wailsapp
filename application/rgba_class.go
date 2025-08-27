package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewRGBAClass() data.ClassStmt {
	return &RGBAClass{
		source:    nil,
		construct: &RGBAConstructMethod{source: nil},
	}
}

func NewRGBAClassFrom(source *applicationsrc.RGBA) data.ClassStmt {
	return &RGBAClass{
		source:    source,
		construct: &RGBAConstructMethod{source: source},
	}
}

type RGBAClass struct {
	node.Node
	source    *applicationsrc.RGBA
	construct data.Method
}

func (s *RGBAClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewRGBAClassFrom(&applicationsrc.RGBA{}), ctx.CreateBaseContext()), nil
}
func (s *RGBAClass) GetName() string         { return "wails\\application\\RGBA" }
func (s *RGBAClass) GetExtend() *string      { return nil }
func (s *RGBAClass) GetImplements() []string { return nil }
func (s *RGBAClass) AsString() string        { return "RGBA{}" }
func (s *RGBAClass) GetSource() any          { return s.source }
func (s *RGBAClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Red":
		return node.NewProperty(nil, "Red", "public", true, data.NewAnyValue(s.source.Red)), true
	case "Green":
		return node.NewProperty(nil, "Green", "public", true, data.NewAnyValue(s.source.Green)), true
	case "Blue":
		return node.NewProperty(nil, "Blue", "public", true, data.NewAnyValue(s.source.Blue)), true
	case "Alpha":
		return node.NewProperty(nil, "Alpha", "public", true, data.NewAnyValue(s.source.Alpha)), true
	}
	return nil, false
}

func (s *RGBAClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Red":   node.NewProperty(nil, "Red", "public", true, data.NewAnyValue(nil)),
		"Green": node.NewProperty(nil, "Green", "public", true, data.NewAnyValue(nil)),
		"Blue":  node.NewProperty(nil, "Blue", "public", true, data.NewAnyValue(nil)),
		"Alpha": node.NewProperty(nil, "Alpha", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *RGBAClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Red":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Red = v.Value.(uint8)
		}
	case "Green":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Green = v.Value.(uint8)
		}
	case "Blue":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Blue = v.Value.(uint8)
		}
	case "Alpha":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Alpha = v.Value.(uint8)
		}
	}
}

func (s *RGBAClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *RGBAClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *RGBAClass) GetConstruct() data.Method { return s.construct }

type RGBAConstructMethod struct {
	source *applicationsrc.RGBA
}

func (h *RGBAConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *RGBAConstructMethod) GetName() string               { return "construct" }
func (h *RGBAConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *RGBAConstructMethod) GetIsStatic() bool             { return false }
func (h *RGBAConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *RGBAConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *RGBAConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
