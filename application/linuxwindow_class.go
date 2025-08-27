package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewLinuxWindowClass() data.ClassStmt {
	return &LinuxWindowClass{
		source:    nil,
		construct: &LinuxWindowConstructMethod{source: nil},
	}
}

func NewLinuxWindowClassFrom(source *applicationsrc.LinuxWindow) data.ClassStmt {
	return &LinuxWindowClass{
		source:    source,
		construct: &LinuxWindowConstructMethod{source: source},
	}
}

type LinuxWindowClass struct {
	node.Node
	source    *applicationsrc.LinuxWindow
	construct data.Method
}

func (s *LinuxWindowClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewLinuxWindowClassFrom(&applicationsrc.LinuxWindow{}), ctx.CreateBaseContext()), nil
}
func (s *LinuxWindowClass) GetName() string         { return "wails\\application\\LinuxWindow" }
func (s *LinuxWindowClass) GetExtend() *string      { return nil }
func (s *LinuxWindowClass) GetImplements() []string { return nil }
func (s *LinuxWindowClass) AsString() string        { return "LinuxWindow{}" }
func (s *LinuxWindowClass) GetSource() any          { return s.source }
func (s *LinuxWindowClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Icon":
		return node.NewProperty(nil, "Icon", "public", true, data.NewAnyValue(s.source.Icon)), true
	case "WindowIsTranslucent":
		return node.NewProperty(nil, "WindowIsTranslucent", "public", true, data.NewAnyValue(s.source.WindowIsTranslucent)), true
	case "WebviewGpuPolicy":
		return node.NewProperty(nil, "WebviewGpuPolicy", "public", true, data.NewAnyValue(s.source.WebviewGpuPolicy)), true
	case "WindowDidMoveDebounceMS":
		return node.NewProperty(nil, "WindowDidMoveDebounceMS", "public", true, data.NewAnyValue(s.source.WindowDidMoveDebounceMS)), true
	case "Menu":
		return node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(NewMenuClassFrom(s.source.Menu), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *LinuxWindowClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Icon":                    node.NewProperty(nil, "Icon", "public", true, data.NewAnyValue(nil)),
		"WindowIsTranslucent":     node.NewProperty(nil, "WindowIsTranslucent", "public", true, data.NewAnyValue(nil)),
		"WebviewGpuPolicy":        node.NewProperty(nil, "WebviewGpuPolicy", "public", true, data.NewAnyValue(nil)),
		"WindowDidMoveDebounceMS": node.NewProperty(nil, "WindowDidMoveDebounceMS", "public", true, data.NewAnyValue(nil)),
		"Menu":                    node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *LinuxWindowClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Icon":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Icon = v.Value.([]uint8)
		}
	case "WindowIsTranslucent":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.WindowIsTranslucent = bool(x)
			}
		}
	case "WebviewGpuPolicy":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.WebviewGpuPolicy = applicationsrc.WebviewGpuPolicy(x)
			}
		}
	case "WindowDidMoveDebounceMS":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.WindowDidMoveDebounceMS = v.Value.(uint16)
		}
	case "Menu":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Menu); ok {
					s.source.Menu = ptr
				}
			}
		case *data.AnyValue:
			s.source.Menu = v.Value.(*applicationsrc.Menu)
		}
	}
}

func (s *LinuxWindowClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *LinuxWindowClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *LinuxWindowClass) GetConstruct() data.Method { return s.construct }

type LinuxWindowConstructMethod struct {
	source *applicationsrc.LinuxWindow
}

func (h *LinuxWindowConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *LinuxWindowConstructMethod) GetName() string               { return "construct" }
func (h *LinuxWindowConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *LinuxWindowConstructMethod) GetIsStatic() bool             { return false }
func (h *LinuxWindowConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *LinuxWindowConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *LinuxWindowConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
