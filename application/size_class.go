package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewSizeClass() data.ClassStmt {
	return &SizeClass{
		source:    nil,
		construct: &SizeConstructMethod{source: nil},
	}
}

func NewSizeClassFrom(source *applicationsrc.Size) data.ClassStmt {
	return &SizeClass{
		source:    source,
		construct: &SizeConstructMethod{source: source},
	}
}

type SizeClass struct {
	node.Node
	source    *applicationsrc.Size
	construct data.Method
}

func (s *SizeClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewSizeClassFrom(&applicationsrc.Size{}), ctx.CreateBaseContext()), nil
}
func (s *SizeClass) GetName() string         { return "wails\\application\\Size" }
func (s *SizeClass) GetExtend() *string      { return nil }
func (s *SizeClass) GetImplements() []string { return nil }
func (s *SizeClass) AsString() string        { return "Size{}" }
func (s *SizeClass) GetSource() any          { return s.source }
func (s *SizeClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Width":
		return node.NewProperty(nil, "Width", "public", true, data.NewAnyValue(s.source.Width)), true
	case "Height":
		return node.NewProperty(nil, "Height", "public", true, data.NewAnyValue(s.source.Height)), true
	}
	return nil, false
}

func (s *SizeClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Width":  node.NewProperty(nil, "Width", "public", true, data.NewAnyValue(nil)),
		"Height": node.NewProperty(nil, "Height", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *SizeClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Width":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Width = int(x)
			}
		}
	case "Height":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Height = int(x)
			}
		}
	}
}

func (s *SizeClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *SizeClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *SizeClass) GetConstruct() data.Method { return s.construct }

type SizeConstructMethod struct {
	source *applicationsrc.Size
}

func (h *SizeConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *SizeConstructMethod) GetName() string               { return "construct" }
func (h *SizeConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *SizeConstructMethod) GetIsStatic() bool             { return false }
func (h *SizeConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *SizeConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *SizeConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
