package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewScreenClass() data.ClassStmt {
	return &ScreenClass{
		source:    nil,
		construct: &ScreenConstructMethod{source: nil},
		origin:    &ScreenOriginMethod{source: nil},
	}
}

func NewScreenClassFrom(source *applicationsrc.Screen) data.ClassStmt {
	return &ScreenClass{
		source:    source,
		construct: &ScreenConstructMethod{source: source},
		origin:    &ScreenOriginMethod{source: source},
	}
}

type ScreenClass struct {
	node.Node
	source    *applicationsrc.Screen
	construct data.Method
	origin    data.Method
}

func (s *ScreenClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewScreenClassFrom(&applicationsrc.Screen{}), ctx.CreateBaseContext()), nil
}
func (s *ScreenClass) GetName() string         { return "wails\\application\\Screen" }
func (s *ScreenClass) GetExtend() *string      { return nil }
func (s *ScreenClass) GetImplements() []string { return nil }
func (s *ScreenClass) AsString() string        { return "Screen{}" }
func (s *ScreenClass) GetSource() any          { return s.source }
func (s *ScreenClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "ID":
		return node.NewProperty(nil, "ID", "public", true, data.NewAnyValue(s.source.ID)), true
	case "Name":
		return node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(s.source.Name)), true
	case "ScaleFactor":
		return node.NewProperty(nil, "ScaleFactor", "public", true, data.NewAnyValue(s.source.ScaleFactor)), true
	case "X":
		return node.NewProperty(nil, "X", "public", true, data.NewAnyValue(s.source.X)), true
	case "Y":
		return node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(s.source.Y)), true
	case "Size":
		return node.NewProperty(nil, "Size", "public", true, data.NewClassValue(NewSizeClassFrom(&s.source.Size), runtime.NewContextToDo())), true
	case "Bounds":
		return node.NewProperty(nil, "Bounds", "public", true, data.NewClassValue(NewRectClassFrom(&s.source.Bounds), runtime.NewContextToDo())), true
	case "PhysicalBounds":
		return node.NewProperty(nil, "PhysicalBounds", "public", true, data.NewClassValue(NewRectClassFrom(&s.source.PhysicalBounds), runtime.NewContextToDo())), true
	case "WorkArea":
		return node.NewProperty(nil, "WorkArea", "public", true, data.NewClassValue(NewRectClassFrom(&s.source.WorkArea), runtime.NewContextToDo())), true
	case "PhysicalWorkArea":
		return node.NewProperty(nil, "PhysicalWorkArea", "public", true, data.NewClassValue(NewRectClassFrom(&s.source.PhysicalWorkArea), runtime.NewContextToDo())), true
	case "IsPrimary":
		return node.NewProperty(nil, "IsPrimary", "public", true, data.NewAnyValue(s.source.IsPrimary)), true
	case "Rotation":
		return node.NewProperty(nil, "Rotation", "public", true, data.NewAnyValue(s.source.Rotation)), true
	}
	return nil, false
}

func (s *ScreenClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"ID":               node.NewProperty(nil, "ID", "public", true, data.NewAnyValue(nil)),
		"Name":             node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(nil)),
		"ScaleFactor":      node.NewProperty(nil, "ScaleFactor", "public", true, data.NewAnyValue(nil)),
		"X":                node.NewProperty(nil, "X", "public", true, data.NewAnyValue(nil)),
		"Y":                node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(nil)),
		"Size":             node.NewProperty(nil, "Size", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Bounds":           node.NewProperty(nil, "Bounds", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"PhysicalBounds":   node.NewProperty(nil, "PhysicalBounds", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"WorkArea":         node.NewProperty(nil, "WorkArea", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"PhysicalWorkArea": node.NewProperty(nil, "PhysicalWorkArea", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"IsPrimary":        node.NewProperty(nil, "IsPrimary", "public", true, data.NewAnyValue(nil)),
		"Rotation":         node.NewProperty(nil, "Rotation", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *ScreenClass) SetProperty(name string, value data.Value) {
	switch name {
	case "ID":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.ID = string(sv.AsString())
		}
	case "Name":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Name = string(sv.AsString())
		}
	case "ScaleFactor":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.ScaleFactor = v.Value.(float32)
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
	case "Size":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Size); ok {
					s.source.Size = *ptr
				} else {
					s.source.Size = src.(applicationsrc.Size)
				}
			}
		case *data.AnyValue:
			s.source.Size = v.Value.(applicationsrc.Size)
		}
	case "Bounds":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Rect); ok {
					s.source.Bounds = *ptr
				} else {
					s.source.Bounds = src.(applicationsrc.Rect)
				}
			}
		case *data.AnyValue:
			s.source.Bounds = v.Value.(applicationsrc.Rect)
		}
	case "PhysicalBounds":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Rect); ok {
					s.source.PhysicalBounds = *ptr
				} else {
					s.source.PhysicalBounds = src.(applicationsrc.Rect)
				}
			}
		case *data.AnyValue:
			s.source.PhysicalBounds = v.Value.(applicationsrc.Rect)
		}
	case "WorkArea":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Rect); ok {
					s.source.WorkArea = *ptr
				} else {
					s.source.WorkArea = src.(applicationsrc.Rect)
				}
			}
		case *data.AnyValue:
			s.source.WorkArea = v.Value.(applicationsrc.Rect)
		}
	case "PhysicalWorkArea":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Rect); ok {
					s.source.PhysicalWorkArea = *ptr
				} else {
					s.source.PhysicalWorkArea = src.(applicationsrc.Rect)
				}
			}
		case *data.AnyValue:
			s.source.PhysicalWorkArea = v.Value.(applicationsrc.Rect)
		}
	case "IsPrimary":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.IsPrimary = bool(x)
			}
		}
	case "Rotation":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Rotation = v.Value.(float32)
		}
	}
}

func (s *ScreenClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "origin":
		return s.origin, true
	}
	return nil, false
}

func (s *ScreenClass) GetMethods() []data.Method {
	return []data.Method{
		s.origin,
	}
}

func (s *ScreenClass) GetConstruct() data.Method { return s.construct }

type ScreenConstructMethod struct {
	source *applicationsrc.Screen
}

func (h *ScreenConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ScreenConstructMethod) GetName() string               { return "construct" }
func (h *ScreenConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ScreenConstructMethod) GetIsStatic() bool             { return false }
func (h *ScreenConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ScreenConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ScreenConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
