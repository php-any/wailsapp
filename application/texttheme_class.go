package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewTextThemeClass() data.ClassStmt {
	return &TextThemeClass{
		source:    nil,
		construct: &TextThemeConstructMethod{source: nil},
	}
}

func NewTextThemeClassFrom(source *applicationsrc.TextTheme) data.ClassStmt {
	return &TextThemeClass{
		source:    source,
		construct: &TextThemeConstructMethod{source: source},
	}
}

type TextThemeClass struct {
	node.Node
	source    *applicationsrc.TextTheme
	construct data.Method
}

func (s *TextThemeClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewTextThemeClassFrom(&applicationsrc.TextTheme{}), ctx.CreateBaseContext()), nil
}
func (s *TextThemeClass) GetName() string         { return "wails\\application\\TextTheme" }
func (s *TextThemeClass) GetExtend() *string      { return nil }
func (s *TextThemeClass) GetImplements() []string { return nil }
func (s *TextThemeClass) AsString() string        { return "TextTheme{}" }
func (s *TextThemeClass) GetSource() any          { return s.source }
func (s *TextThemeClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Text":
		return node.NewProperty(nil, "Text", "public", true, data.NewAnyValue(s.source.Text)), true
	case "Background":
		return node.NewProperty(nil, "Background", "public", true, data.NewAnyValue(s.source.Background)), true
	}
	return nil, false
}

func (s *TextThemeClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Text":       node.NewProperty(nil, "Text", "public", true, data.NewAnyValue(nil)),
		"Background": node.NewProperty(nil, "Background", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *TextThemeClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Text":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Text = v.Value.(*uint32)
		}
	case "Background":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Background = v.Value.(*uint32)
		}
	}
}

func (s *TextThemeClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *TextThemeClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *TextThemeClass) GetConstruct() data.Method { return s.construct }

type TextThemeConstructMethod struct {
	source *applicationsrc.TextTheme
}

func (h *TextThemeConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *TextThemeConstructMethod) GetName() string               { return "construct" }
func (h *TextThemeConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *TextThemeConstructMethod) GetIsStatic() bool             { return false }
func (h *TextThemeConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *TextThemeConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *TextThemeConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
