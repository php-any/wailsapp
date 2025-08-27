package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMacLiquidGlassClass() data.ClassStmt {
	return &MacLiquidGlassClass{
		source:    nil,
		construct: &MacLiquidGlassConstructMethod{source: nil},
	}
}

func NewMacLiquidGlassClassFrom(source *applicationsrc.MacLiquidGlass) data.ClassStmt {
	return &MacLiquidGlassClass{
		source:    source,
		construct: &MacLiquidGlassConstructMethod{source: source},
	}
}

type MacLiquidGlassClass struct {
	node.Node
	source    *applicationsrc.MacLiquidGlass
	construct data.Method
}

func (s *MacLiquidGlassClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMacLiquidGlassClassFrom(&applicationsrc.MacLiquidGlass{}), ctx.CreateBaseContext()), nil
}
func (s *MacLiquidGlassClass) GetName() string         { return "wails\\application\\MacLiquidGlass" }
func (s *MacLiquidGlassClass) GetExtend() *string      { return nil }
func (s *MacLiquidGlassClass) GetImplements() []string { return nil }
func (s *MacLiquidGlassClass) AsString() string        { return "MacLiquidGlass{}" }
func (s *MacLiquidGlassClass) GetSource() any          { return s.source }
func (s *MacLiquidGlassClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Style":
		return node.NewProperty(nil, "Style", "public", true, data.NewAnyValue(s.source.Style)), true
	case "Material":
		return node.NewProperty(nil, "Material", "public", true, data.NewAnyValue(s.source.Material)), true
	case "CornerRadius":
		return node.NewProperty(nil, "CornerRadius", "public", true, data.NewAnyValue(s.source.CornerRadius)), true
	case "TintColor":
		return node.NewProperty(nil, "TintColor", "public", true, data.NewClassValue(NewRGBAClassFrom(s.source.TintColor), runtime.NewContextToDo())), true
	case "GroupID":
		return node.NewProperty(nil, "GroupID", "public", true, data.NewAnyValue(s.source.GroupID)), true
	case "GroupSpacing":
		return node.NewProperty(nil, "GroupSpacing", "public", true, data.NewAnyValue(s.source.GroupSpacing)), true
	}
	return nil, false
}

func (s *MacLiquidGlassClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Style":        node.NewProperty(nil, "Style", "public", true, data.NewAnyValue(nil)),
		"Material":     node.NewProperty(nil, "Material", "public", true, data.NewAnyValue(nil)),
		"CornerRadius": node.NewProperty(nil, "CornerRadius", "public", true, data.NewAnyValue(nil)),
		"TintColor":    node.NewProperty(nil, "TintColor", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"GroupID":      node.NewProperty(nil, "GroupID", "public", true, data.NewAnyValue(nil)),
		"GroupSpacing": node.NewProperty(nil, "GroupSpacing", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *MacLiquidGlassClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Style":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Style = applicationsrc.MacLiquidGlassStyle(x)
			}
		}
	case "Material":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Material = applicationsrc.NSVisualEffectMaterial(x)
			}
		}
	case "CornerRadius":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.CornerRadius = v.Value.(float64)
		}
	case "TintColor":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.RGBA); ok {
					s.source.TintColor = ptr
				}
			}
		case *data.AnyValue:
			s.source.TintColor = v.Value.(*applicationsrc.RGBA)
		}
	case "GroupID":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.GroupID = string(sv.AsString())
		}
	case "GroupSpacing":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.GroupSpacing = v.Value.(float64)
		}
	}
}

func (s *MacLiquidGlassClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MacLiquidGlassClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MacLiquidGlassClass) GetConstruct() data.Method { return s.construct }

type MacLiquidGlassConstructMethod struct {
	source *applicationsrc.MacLiquidGlass
}

func (h *MacLiquidGlassConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MacLiquidGlassConstructMethod) GetName() string               { return "construct" }
func (h *MacLiquidGlassConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MacLiquidGlassConstructMethod) GetIsStatic() bool             { return false }
func (h *MacLiquidGlassConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MacLiquidGlassConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MacLiquidGlassConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
