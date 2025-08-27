package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewPointClass() data.ClassStmt {
	return &PointClass{
		source:    nil,
		construct: &PointConstructMethod{source: nil},
	}
}

func NewPointClassFrom(source *applicationsrc.Point) data.ClassStmt {
	return &PointClass{
		source:    source,
		construct: &PointConstructMethod{source: source},
	}
}

type PointClass struct {
	node.Node
	source    *applicationsrc.Point
	construct data.Method
}

func (s *PointClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewPointClassFrom(&applicationsrc.Point{}), ctx.CreateBaseContext()), nil
}
func (s *PointClass) GetName() string         { return "wails\\application\\Point" }
func (s *PointClass) GetExtend() *string      { return nil }
func (s *PointClass) GetImplements() []string { return nil }
func (s *PointClass) AsString() string        { return "Point{}" }
func (s *PointClass) GetSource() any          { return s.source }
func (s *PointClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "X":
		return node.NewProperty(nil, "X", "public", true, data.NewAnyValue(s.source.X)), true
	case "Y":
		return node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(s.source.Y)), true
	}
	return nil, false
}

func (s *PointClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"X": node.NewProperty(nil, "X", "public", true, data.NewAnyValue(nil)),
		"Y": node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *PointClass) SetProperty(name string, value data.Value) {
	switch name {
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
	}
}

func (s *PointClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *PointClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *PointClass) GetConstruct() data.Method { return s.construct }

type PointConstructMethod struct {
	source *applicationsrc.Point
}

func (h *PointConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *PointConstructMethod) GetName() string               { return "construct" }
func (h *PointConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *PointConstructMethod) GetIsStatic() bool             { return false }
func (h *PointConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *PointConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *PointConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
