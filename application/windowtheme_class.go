package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWindowThemeClass() data.ClassStmt {
	return &WindowThemeClass{
		source:    nil,
		construct: &WindowThemeConstructMethod{source: nil},
	}
}

func NewWindowThemeClassFrom(source *applicationsrc.WindowTheme) data.ClassStmt {
	return &WindowThemeClass{
		source:    source,
		construct: &WindowThemeConstructMethod{source: source},
	}
}

type WindowThemeClass struct {
	node.Node
	source    *applicationsrc.WindowTheme
	construct data.Method
}

func (s *WindowThemeClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWindowThemeClassFrom(&applicationsrc.WindowTheme{}), ctx.CreateBaseContext()), nil
}
func (s *WindowThemeClass) GetName() string         { return "wails\\application\\WindowTheme" }
func (s *WindowThemeClass) GetExtend() *string      { return nil }
func (s *WindowThemeClass) GetImplements() []string { return nil }
func (s *WindowThemeClass) AsString() string        { return "WindowTheme{}" }
func (s *WindowThemeClass) GetSource() any          { return s.source }
func (s *WindowThemeClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "BorderColour":
		return node.NewProperty(nil, "BorderColour", "public", true, data.NewAnyValue(s.source.BorderColour)), true
	case "TitleBarColour":
		return node.NewProperty(nil, "TitleBarColour", "public", true, data.NewAnyValue(s.source.TitleBarColour)), true
	case "TitleTextColour":
		return node.NewProperty(nil, "TitleTextColour", "public", true, data.NewAnyValue(s.source.TitleTextColour)), true
	}
	return nil, false
}

func (s *WindowThemeClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"BorderColour":    node.NewProperty(nil, "BorderColour", "public", true, data.NewAnyValue(nil)),
		"TitleBarColour":  node.NewProperty(nil, "TitleBarColour", "public", true, data.NewAnyValue(nil)),
		"TitleTextColour": node.NewProperty(nil, "TitleTextColour", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *WindowThemeClass) SetProperty(name string, value data.Value) {
	switch name {
	case "BorderColour":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.BorderColour = v.Value.(*uint32)
		}
	case "TitleBarColour":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.TitleBarColour = v.Value.(*uint32)
		}
	case "TitleTextColour":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.TitleTextColour = v.Value.(*uint32)
		}
	}
}

func (s *WindowThemeClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *WindowThemeClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *WindowThemeClass) GetConstruct() data.Method { return s.construct }

type WindowThemeConstructMethod struct {
	source *applicationsrc.WindowTheme
}

func (h *WindowThemeConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowThemeConstructMethod) GetName() string               { return "construct" }
func (h *WindowThemeConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowThemeConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowThemeConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowThemeConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowThemeConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
