package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewContextMenuDataClass() data.ClassStmt {
	return &ContextMenuDataClass{
		source:    nil,
		construct: &ContextMenuDataConstructMethod{source: nil},
	}
}

func NewContextMenuDataClassFrom(source *applicationsrc.ContextMenuData) data.ClassStmt {
	return &ContextMenuDataClass{
		source:    source,
		construct: &ContextMenuDataConstructMethod{source: source},
	}
}

type ContextMenuDataClass struct {
	node.Node
	source    *applicationsrc.ContextMenuData
	construct data.Method
}

func (s *ContextMenuDataClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewContextMenuDataClassFrom(&applicationsrc.ContextMenuData{}), ctx.CreateBaseContext()), nil
}
func (s *ContextMenuDataClass) GetName() string         { return "wails\\application\\ContextMenuData" }
func (s *ContextMenuDataClass) GetExtend() *string      { return nil }
func (s *ContextMenuDataClass) GetImplements() []string { return nil }
func (s *ContextMenuDataClass) AsString() string        { return "ContextMenuData{}" }
func (s *ContextMenuDataClass) GetSource() any          { return s.source }
func (s *ContextMenuDataClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Id":
		return node.NewProperty(nil, "Id", "public", true, data.NewAnyValue(s.source.Id)), true
	case "X":
		return node.NewProperty(nil, "X", "public", true, data.NewAnyValue(s.source.X)), true
	case "Y":
		return node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(s.source.Y)), true
	case "Data":
		return node.NewProperty(nil, "Data", "public", true, data.NewAnyValue(s.source.Data)), true
	}
	return nil, false
}

func (s *ContextMenuDataClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Id":   node.NewProperty(nil, "Id", "public", true, data.NewAnyValue(nil)),
		"X":    node.NewProperty(nil, "X", "public", true, data.NewAnyValue(nil)),
		"Y":    node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(nil)),
		"Data": node.NewProperty(nil, "Data", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *ContextMenuDataClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Id":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Id = string(sv.AsString())
		}
	case "X":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.X = int(x)
			}
		}
	case "Y":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Y = int(x)
			}
		}
	case "Data":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Data = string(sv.AsString())
		}
	}
}

func (s *ContextMenuDataClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *ContextMenuDataClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *ContextMenuDataClass) GetConstruct() data.Method { return s.construct }

type ContextMenuDataConstructMethod struct {
	source *applicationsrc.ContextMenuData
}

func (h *ContextMenuDataConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ContextMenuDataConstructMethod) GetName() string               { return "construct" }
func (h *ContextMenuDataConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ContextMenuDataConstructMethod) GetIsStatic() bool             { return false }
func (h *ContextMenuDataConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ContextMenuDataConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ContextMenuDataConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
