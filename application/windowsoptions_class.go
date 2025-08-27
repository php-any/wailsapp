package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWindowsOptionsClass() data.ClassStmt {
	return &WindowsOptionsClass{
		source:    nil,
		construct: &WindowsOptionsConstructMethod{source: nil},
	}
}

func NewWindowsOptionsClassFrom(source *applicationsrc.WindowsOptions) data.ClassStmt {
	return &WindowsOptionsClass{
		source:    source,
		construct: &WindowsOptionsConstructMethod{source: source},
	}
}

type WindowsOptionsClass struct {
	node.Node
	source    *applicationsrc.WindowsOptions
	construct data.Method
}

func (s *WindowsOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWindowsOptionsClassFrom(&applicationsrc.WindowsOptions{}), ctx.CreateBaseContext()), nil
}
func (s *WindowsOptionsClass) GetName() string         { return "wails\\application\\WindowsOptions" }
func (s *WindowsOptionsClass) GetExtend() *string      { return nil }
func (s *WindowsOptionsClass) GetImplements() []string { return nil }
func (s *WindowsOptionsClass) AsString() string        { return "WindowsOptions{}" }
func (s *WindowsOptionsClass) GetSource() any          { return s.source }
func (s *WindowsOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "WndClass":
		return node.NewProperty(nil, "WndClass", "public", true, data.NewAnyValue(s.source.WndClass)), true
	case "WndProcInterceptor":
		return node.NewProperty(nil, "WndProcInterceptor", "public", true, data.NewAnyValue(s.source.WndProcInterceptor)), true
	case "DisableQuitOnLastWindowClosed":
		return node.NewProperty(nil, "DisableQuitOnLastWindowClosed", "public", true, data.NewAnyValue(s.source.DisableQuitOnLastWindowClosed)), true
	case "WebviewUserDataPath":
		return node.NewProperty(nil, "WebviewUserDataPath", "public", true, data.NewAnyValue(s.source.WebviewUserDataPath)), true
	case "WebviewBrowserPath":
		return node.NewProperty(nil, "WebviewBrowserPath", "public", true, data.NewAnyValue(s.source.WebviewBrowserPath)), true
	}
	return nil, false
}

func (s *WindowsOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"WndClass":                      node.NewProperty(nil, "WndClass", "public", true, data.NewAnyValue(nil)),
		"WndProcInterceptor":            node.NewProperty(nil, "WndProcInterceptor", "public", true, data.NewAnyValue(nil)),
		"DisableQuitOnLastWindowClosed": node.NewProperty(nil, "DisableQuitOnLastWindowClosed", "public", true, data.NewAnyValue(nil)),
		"WebviewUserDataPath":           node.NewProperty(nil, "WebviewUserDataPath", "public", true, data.NewAnyValue(nil)),
		"WebviewBrowserPath":            node.NewProperty(nil, "WebviewBrowserPath", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *WindowsOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "WndClass":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.WndClass = string(sv.AsString())
		}
	case "WndProcInterceptor":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.WndProcInterceptor = v.Value.(func(uintptr, uint32, uintptr, uintptr) (uintptr, bool))
		}
	case "DisableQuitOnLastWindowClosed":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableQuitOnLastWindowClosed = bool(x)
			}
		}
	case "WebviewUserDataPath":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.WebviewUserDataPath = string(sv.AsString())
		}
	case "WebviewBrowserPath":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.WebviewBrowserPath = string(sv.AsString())
		}
	}
}

func (s *WindowsOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *WindowsOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *WindowsOptionsClass) GetConstruct() data.Method { return s.construct }

type WindowsOptionsConstructMethod struct {
	source *applicationsrc.WindowsOptions
}

func (h *WindowsOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowsOptionsConstructMethod) GetName() string               { return "construct" }
func (h *WindowsOptionsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowsOptionsConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowsOptionsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowsOptionsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowsOptionsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
