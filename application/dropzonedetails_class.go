package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewDropZoneDetailsClass() data.ClassStmt {
	return &DropZoneDetailsClass{
		source:    nil,
		construct: &DropZoneDetailsConstructMethod{source: nil},
	}
}

func NewDropZoneDetailsClassFrom(source *applicationsrc.DropZoneDetails) data.ClassStmt {
	return &DropZoneDetailsClass{
		source:    source,
		construct: &DropZoneDetailsConstructMethod{source: source},
	}
}

type DropZoneDetailsClass struct {
	node.Node
	source    *applicationsrc.DropZoneDetails
	construct data.Method
}

func (s *DropZoneDetailsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewDropZoneDetailsClassFrom(&applicationsrc.DropZoneDetails{}), ctx.CreateBaseContext()), nil
}
func (s *DropZoneDetailsClass) GetName() string         { return "wails\\application\\DropZoneDetails" }
func (s *DropZoneDetailsClass) GetExtend() *string      { return nil }
func (s *DropZoneDetailsClass) GetImplements() []string { return nil }
func (s *DropZoneDetailsClass) AsString() string        { return "DropZoneDetails{}" }
func (s *DropZoneDetailsClass) GetSource() any          { return s.source }
func (s *DropZoneDetailsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "X":
		return node.NewProperty(nil, "X", "public", true, data.NewAnyValue(s.source.X)), true
	case "Y":
		return node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(s.source.Y)), true
	case "ElementID":
		return node.NewProperty(nil, "ElementID", "public", true, data.NewAnyValue(s.source.ElementID)), true
	case "ClassList":
		return node.NewProperty(nil, "ClassList", "public", true, data.NewAnyValue(s.source.ClassList)), true
	case "Attributes":
		return node.NewProperty(nil, "Attributes", "public", true, data.NewAnyValue(s.source.Attributes)), true
	}
	return nil, false
}

func (s *DropZoneDetailsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"X":          node.NewProperty(nil, "X", "public", true, data.NewAnyValue(nil)),
		"Y":          node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(nil)),
		"ElementID":  node.NewProperty(nil, "ElementID", "public", true, data.NewAnyValue(nil)),
		"ClassList":  node.NewProperty(nil, "ClassList", "public", true, data.NewAnyValue(nil)),
		"Attributes": node.NewProperty(nil, "Attributes", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *DropZoneDetailsClass) SetProperty(name string, value data.Value) {
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
	case "ElementID":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.ElementID = string(sv.AsString())
		}
	case "ClassList":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.ClassList = v.Value.([]string)
		}
	case "Attributes":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Attributes = v.Value.(map[string]string)
		}
	}
}

func (s *DropZoneDetailsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *DropZoneDetailsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *DropZoneDetailsClass) GetConstruct() data.Method { return s.construct }

type DropZoneDetailsConstructMethod struct {
	source *applicationsrc.DropZoneDetails
}

func (h *DropZoneDetailsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *DropZoneDetailsConstructMethod) GetName() string               { return "construct" }
func (h *DropZoneDetailsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *DropZoneDetailsConstructMethod) GetIsStatic() bool             { return false }
func (h *DropZoneDetailsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *DropZoneDetailsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *DropZoneDetailsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
