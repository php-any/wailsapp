package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewRectClass() data.ClassStmt {
	return &RectClass{
		source:       nil,
		construct:    &RectConstructMethod{source: nil},
		contains:     &RectContainsMethod{source: nil},
		corner:       &RectCornerMethod{source: nil},
		insideCorner: &RectInsideCornerMethod{source: nil},
		intersect:    &RectIntersectMethod{source: nil},
		isEmpty:      &RectIsEmptyMethod{source: nil},
		origin:       &RectOriginMethod{source: nil},
		size:         &RectSizeMethod{source: nil},
	}
}

func NewRectClassFrom(source *applicationsrc.Rect) data.ClassStmt {
	return &RectClass{
		source:       source,
		construct:    &RectConstructMethod{source: source},
		contains:     &RectContainsMethod{source: source},
		corner:       &RectCornerMethod{source: source},
		insideCorner: &RectInsideCornerMethod{source: source},
		intersect:    &RectIntersectMethod{source: source},
		isEmpty:      &RectIsEmptyMethod{source: source},
		origin:       &RectOriginMethod{source: source},
		size:         &RectSizeMethod{source: source},
	}
}

type RectClass struct {
	node.Node
	source       *applicationsrc.Rect
	construct    data.Method
	contains     data.Method
	corner       data.Method
	insideCorner data.Method
	intersect    data.Method
	isEmpty      data.Method
	origin       data.Method
	size         data.Method
}

func (s *RectClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewRectClassFrom(&applicationsrc.Rect{}), ctx.CreateBaseContext()), nil
}
func (s *RectClass) GetName() string         { return "wails\\application\\Rect" }
func (s *RectClass) GetExtend() *string      { return nil }
func (s *RectClass) GetImplements() []string { return nil }
func (s *RectClass) AsString() string        { return "Rect{}" }
func (s *RectClass) GetSource() any          { return s.source }
func (s *RectClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "X":
		return node.NewProperty(nil, "X", "public", true, data.NewAnyValue(s.source.X)), true
	case "Y":
		return node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(s.source.Y)), true
	case "Width":
		return node.NewProperty(nil, "Width", "public", true, data.NewAnyValue(s.source.Width)), true
	case "Height":
		return node.NewProperty(nil, "Height", "public", true, data.NewAnyValue(s.source.Height)), true
	}
	return nil, false
}

func (s *RectClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"X":      node.NewProperty(nil, "X", "public", true, data.NewAnyValue(nil)),
		"Y":      node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(nil)),
		"Width":  node.NewProperty(nil, "Width", "public", true, data.NewAnyValue(nil)),
		"Height": node.NewProperty(nil, "Height", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *RectClass) SetProperty(name string, value data.Value) {
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

func (s *RectClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "contains":
		return s.contains, true
	case "corner":
		return s.corner, true
	case "insideCorner":
		return s.insideCorner, true
	case "intersect":
		return s.intersect, true
	case "isEmpty":
		return s.isEmpty, true
	case "origin":
		return s.origin, true
	case "size":
		return s.size, true
	}
	return nil, false
}

func (s *RectClass) GetMethods() []data.Method {
	return []data.Method{
		s.contains,
		s.corner,
		s.insideCorner,
		s.intersect,
		s.isEmpty,
		s.origin,
		s.size,
	}
}

func (s *RectClass) GetConstruct() data.Method { return s.construct }

type RectConstructMethod struct {
	source *applicationsrc.Rect
}

func (h *RectConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *RectConstructMethod) GetName() string               { return "construct" }
func (h *RectConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *RectConstructMethod) GetIsStatic() bool             { return false }
func (h *RectConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *RectConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *RectConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
