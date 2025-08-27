package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewLRTBClass() data.ClassStmt {
	return &LRTBClass{
		source:    nil,
		construct: &LRTBConstructMethod{source: nil},
	}
}

func NewLRTBClassFrom(source *applicationsrc.LRTB) data.ClassStmt {
	return &LRTBClass{
		source:    source,
		construct: &LRTBConstructMethod{source: source},
	}
}

type LRTBClass struct {
	node.Node
	source    *applicationsrc.LRTB
	construct data.Method
}

func (s *LRTBClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewLRTBClassFrom(&applicationsrc.LRTB{}), ctx.CreateBaseContext()), nil
}
func (s *LRTBClass) GetName() string         { return "wails\\application\\LRTB" }
func (s *LRTBClass) GetExtend() *string      { return nil }
func (s *LRTBClass) GetImplements() []string { return nil }
func (s *LRTBClass) AsString() string        { return "LRTB{}" }
func (s *LRTBClass) GetSource() any          { return s.source }
func (s *LRTBClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Left":
		return node.NewProperty(nil, "Left", "public", true, data.NewAnyValue(s.source.Left)), true
	case "Right":
		return node.NewProperty(nil, "Right", "public", true, data.NewAnyValue(s.source.Right)), true
	case "Top":
		return node.NewProperty(nil, "Top", "public", true, data.NewAnyValue(s.source.Top)), true
	case "Bottom":
		return node.NewProperty(nil, "Bottom", "public", true, data.NewAnyValue(s.source.Bottom)), true
	}
	return nil, false
}

func (s *LRTBClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Left":   node.NewProperty(nil, "Left", "public", true, data.NewAnyValue(nil)),
		"Right":  node.NewProperty(nil, "Right", "public", true, data.NewAnyValue(nil)),
		"Top":    node.NewProperty(nil, "Top", "public", true, data.NewAnyValue(nil)),
		"Bottom": node.NewProperty(nil, "Bottom", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *LRTBClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Left":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Left = int(x)
			}
		}
	case "Right":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Right = int(x)
			}
		}
	case "Top":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Top = int(x)
			}
		}
	case "Bottom":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Bottom = int(x)
			}
		}
	}
}

func (s *LRTBClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *LRTBClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *LRTBClass) GetConstruct() data.Method { return s.construct }

type LRTBConstructMethod struct {
	source *applicationsrc.LRTB
}

func (h *LRTBConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *LRTBConstructMethod) GetName() string               { return "construct" }
func (h *LRTBConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *LRTBConstructMethod) GetIsStatic() bool             { return false }
func (h *LRTBConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *LRTBConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *LRTBConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
