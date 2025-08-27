package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWebviewWindowOptionsClass() data.ClassStmt {
	return &WebviewWindowOptionsClass{
		source:    nil,
		construct: &WebviewWindowOptionsConstructMethod{source: nil},
	}
}

func NewWebviewWindowOptionsClassFrom(source *applicationsrc.WebviewWindowOptions) data.ClassStmt {
	return &WebviewWindowOptionsClass{
		source:    source,
		construct: &WebviewWindowOptionsConstructMethod{source: source},
	}
}

type WebviewWindowOptionsClass struct {
	node.Node
	source    *applicationsrc.WebviewWindowOptions
	construct data.Method
}

func (s *WebviewWindowOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWebviewWindowOptionsClassFrom(&applicationsrc.WebviewWindowOptions{}), ctx.CreateBaseContext()), nil
}
func (s *WebviewWindowOptionsClass) GetName() string {
	return "wails\\application\\WebviewWindowOptions"
}
func (s *WebviewWindowOptionsClass) GetExtend() *string      { return nil }
func (s *WebviewWindowOptionsClass) GetImplements() []string { return nil }
func (s *WebviewWindowOptionsClass) AsString() string        { return "WebviewWindowOptions{}" }
func (s *WebviewWindowOptionsClass) GetSource() any          { return s.source }
func (s *WebviewWindowOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Name":
		return node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(s.source.Name)), true
	case "Title":
		return node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(s.source.Title)), true
	case "Width":
		return node.NewProperty(nil, "Width", "public", true, data.NewAnyValue(s.source.Width)), true
	case "Height":
		return node.NewProperty(nil, "Height", "public", true, data.NewAnyValue(s.source.Height)), true
	case "AlwaysOnTop":
		return node.NewProperty(nil, "AlwaysOnTop", "public", true, data.NewAnyValue(s.source.AlwaysOnTop)), true
	case "URL":
		return node.NewProperty(nil, "URL", "public", true, data.NewAnyValue(s.source.URL)), true
	case "DisableResize":
		return node.NewProperty(nil, "DisableResize", "public", true, data.NewAnyValue(s.source.DisableResize)), true
	case "Frameless":
		return node.NewProperty(nil, "Frameless", "public", true, data.NewAnyValue(s.source.Frameless)), true
	case "MinWidth":
		return node.NewProperty(nil, "MinWidth", "public", true, data.NewAnyValue(s.source.MinWidth)), true
	case "MinHeight":
		return node.NewProperty(nil, "MinHeight", "public", true, data.NewAnyValue(s.source.MinHeight)), true
	case "MaxWidth":
		return node.NewProperty(nil, "MaxWidth", "public", true, data.NewAnyValue(s.source.MaxWidth)), true
	case "MaxHeight":
		return node.NewProperty(nil, "MaxHeight", "public", true, data.NewAnyValue(s.source.MaxHeight)), true
	case "StartState":
		return node.NewProperty(nil, "StartState", "public", true, data.NewAnyValue(s.source.StartState)), true
	case "BackgroundType":
		return node.NewProperty(nil, "BackgroundType", "public", true, data.NewAnyValue(s.source.BackgroundType)), true
	case "BackgroundColour":
		return node.NewProperty(nil, "BackgroundColour", "public", true, data.NewClassValue(NewRGBAClassFrom(&s.source.BackgroundColour), runtime.NewContextToDo())), true
	case "HTML":
		return node.NewProperty(nil, "HTML", "public", true, data.NewAnyValue(s.source.HTML)), true
	case "JS":
		return node.NewProperty(nil, "JS", "public", true, data.NewAnyValue(s.source.JS)), true
	case "CSS":
		return node.NewProperty(nil, "CSS", "public", true, data.NewAnyValue(s.source.CSS)), true
	case "InitialPosition":
		return node.NewProperty(nil, "InitialPosition", "public", true, data.NewAnyValue(s.source.InitialPosition)), true
	case "X":
		return node.NewProperty(nil, "X", "public", true, data.NewAnyValue(s.source.X)), true
	case "Y":
		return node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(s.source.Y)), true
	case "Hidden":
		return node.NewProperty(nil, "Hidden", "public", true, data.NewAnyValue(s.source.Hidden)), true
	case "Zoom":
		return node.NewProperty(nil, "Zoom", "public", true, data.NewAnyValue(s.source.Zoom)), true
	case "ZoomControlEnabled":
		return node.NewProperty(nil, "ZoomControlEnabled", "public", true, data.NewAnyValue(s.source.ZoomControlEnabled)), true
	case "EnableDragAndDrop":
		return node.NewProperty(nil, "EnableDragAndDrop", "public", true, data.NewAnyValue(s.source.EnableDragAndDrop)), true
	case "OpenInspectorOnStartup":
		return node.NewProperty(nil, "OpenInspectorOnStartup", "public", true, data.NewAnyValue(s.source.OpenInspectorOnStartup)), true
	case "Mac":
		return node.NewProperty(nil, "Mac", "public", true, data.NewClassValue(NewMacWindowClassFrom(&s.source.Mac), runtime.NewContextToDo())), true
	case "Windows":
		return node.NewProperty(nil, "Windows", "public", true, data.NewClassValue(NewWindowsWindowClassFrom(&s.source.Windows), runtime.NewContextToDo())), true
	case "Linux":
		return node.NewProperty(nil, "Linux", "public", true, data.NewClassValue(NewLinuxWindowClassFrom(&s.source.Linux), runtime.NewContextToDo())), true
	case "MinimiseButtonState":
		return node.NewProperty(nil, "MinimiseButtonState", "public", true, data.NewAnyValue(s.source.MinimiseButtonState)), true
	case "MaximiseButtonState":
		return node.NewProperty(nil, "MaximiseButtonState", "public", true, data.NewAnyValue(s.source.MaximiseButtonState)), true
	case "CloseButtonState":
		return node.NewProperty(nil, "CloseButtonState", "public", true, data.NewAnyValue(s.source.CloseButtonState)), true
	case "DevToolsEnabled":
		return node.NewProperty(nil, "DevToolsEnabled", "public", true, data.NewAnyValue(s.source.DevToolsEnabled)), true
	case "DefaultContextMenuDisabled":
		return node.NewProperty(nil, "DefaultContextMenuDisabled", "public", true, data.NewAnyValue(s.source.DefaultContextMenuDisabled)), true
	case "KeyBindings":
		return node.NewProperty(nil, "KeyBindings", "public", true, data.NewAnyValue(s.source.KeyBindings)), true
	case "IgnoreMouseEvents":
		return node.NewProperty(nil, "IgnoreMouseEvents", "public", true, data.NewAnyValue(s.source.IgnoreMouseEvents)), true
	case "ContentProtectionEnabled":
		return node.NewProperty(nil, "ContentProtectionEnabled", "public", true, data.NewAnyValue(s.source.ContentProtectionEnabled)), true
	}
	return nil, false
}

func (s *WebviewWindowOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Name":                       node.NewProperty(nil, "Name", "public", true, data.NewAnyValue(nil)),
		"Title":                      node.NewProperty(nil, "Title", "public", true, data.NewAnyValue(nil)),
		"Width":                      node.NewProperty(nil, "Width", "public", true, data.NewAnyValue(nil)),
		"Height":                     node.NewProperty(nil, "Height", "public", true, data.NewAnyValue(nil)),
		"AlwaysOnTop":                node.NewProperty(nil, "AlwaysOnTop", "public", true, data.NewAnyValue(nil)),
		"URL":                        node.NewProperty(nil, "URL", "public", true, data.NewAnyValue(nil)),
		"DisableResize":              node.NewProperty(nil, "DisableResize", "public", true, data.NewAnyValue(nil)),
		"Frameless":                  node.NewProperty(nil, "Frameless", "public", true, data.NewAnyValue(nil)),
		"MinWidth":                   node.NewProperty(nil, "MinWidth", "public", true, data.NewAnyValue(nil)),
		"MinHeight":                  node.NewProperty(nil, "MinHeight", "public", true, data.NewAnyValue(nil)),
		"MaxWidth":                   node.NewProperty(nil, "MaxWidth", "public", true, data.NewAnyValue(nil)),
		"MaxHeight":                  node.NewProperty(nil, "MaxHeight", "public", true, data.NewAnyValue(nil)),
		"StartState":                 node.NewProperty(nil, "StartState", "public", true, data.NewAnyValue(nil)),
		"BackgroundType":             node.NewProperty(nil, "BackgroundType", "public", true, data.NewAnyValue(nil)),
		"BackgroundColour":           node.NewProperty(nil, "BackgroundColour", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"HTML":                       node.NewProperty(nil, "HTML", "public", true, data.NewAnyValue(nil)),
		"JS":                         node.NewProperty(nil, "JS", "public", true, data.NewAnyValue(nil)),
		"CSS":                        node.NewProperty(nil, "CSS", "public", true, data.NewAnyValue(nil)),
		"InitialPosition":            node.NewProperty(nil, "InitialPosition", "public", true, data.NewAnyValue(nil)),
		"X":                          node.NewProperty(nil, "X", "public", true, data.NewAnyValue(nil)),
		"Y":                          node.NewProperty(nil, "Y", "public", true, data.NewAnyValue(nil)),
		"Hidden":                     node.NewProperty(nil, "Hidden", "public", true, data.NewAnyValue(nil)),
		"Zoom":                       node.NewProperty(nil, "Zoom", "public", true, data.NewAnyValue(nil)),
		"ZoomControlEnabled":         node.NewProperty(nil, "ZoomControlEnabled", "public", true, data.NewAnyValue(nil)),
		"EnableDragAndDrop":          node.NewProperty(nil, "EnableDragAndDrop", "public", true, data.NewAnyValue(nil)),
		"OpenInspectorOnStartup":     node.NewProperty(nil, "OpenInspectorOnStartup", "public", true, data.NewAnyValue(nil)),
		"Mac":                        node.NewProperty(nil, "Mac", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Windows":                    node.NewProperty(nil, "Windows", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Linux":                      node.NewProperty(nil, "Linux", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"MinimiseButtonState":        node.NewProperty(nil, "MinimiseButtonState", "public", true, data.NewAnyValue(nil)),
		"MaximiseButtonState":        node.NewProperty(nil, "MaximiseButtonState", "public", true, data.NewAnyValue(nil)),
		"CloseButtonState":           node.NewProperty(nil, "CloseButtonState", "public", true, data.NewAnyValue(nil)),
		"DevToolsEnabled":            node.NewProperty(nil, "DevToolsEnabled", "public", true, data.NewAnyValue(nil)),
		"DefaultContextMenuDisabled": node.NewProperty(nil, "DefaultContextMenuDisabled", "public", true, data.NewAnyValue(nil)),
		"KeyBindings":                node.NewProperty(nil, "KeyBindings", "public", true, data.NewAnyValue(nil)),
		"IgnoreMouseEvents":          node.NewProperty(nil, "IgnoreMouseEvents", "public", true, data.NewAnyValue(nil)),
		"ContentProtectionEnabled":   node.NewProperty(nil, "ContentProtectionEnabled", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *WebviewWindowOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Name":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Name = string(sv.AsString())
		}
	case "Title":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Title = string(sv.AsString())
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
	case "AlwaysOnTop":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.AlwaysOnTop = bool(x)
			}
		}
	case "URL":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.URL = string(sv.AsString())
		}
	case "DisableResize":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableResize = bool(x)
			}
		}
	case "Frameless":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.Frameless = bool(x)
			}
		}
	case "MinWidth":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.MinWidth = int(x)
			}
		}
	case "MinHeight":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.MinHeight = int(x)
			}
		}
	case "MaxWidth":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.MaxWidth = int(x)
			}
		}
	case "MaxHeight":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.MaxHeight = int(x)
			}
		}
	case "StartState":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.StartState = applicationsrc.WindowState(x)
			}
		}
	case "BackgroundType":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.BackgroundType = applicationsrc.BackgroundType(x)
			}
		}
	case "BackgroundColour":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.RGBA); ok {
					s.source.BackgroundColour = *ptr
				} else {
					s.source.BackgroundColour = src.(applicationsrc.RGBA)
				}
			}
		case *data.AnyValue:
			s.source.BackgroundColour = v.Value.(applicationsrc.RGBA)
		}
	case "HTML":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.HTML = string(sv.AsString())
		}
	case "JS":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.JS = string(sv.AsString())
		}
	case "CSS":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.CSS = string(sv.AsString())
		}
	case "InitialPosition":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.InitialPosition = applicationsrc.WindowStartPosition(x)
			}
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
	case "Hidden":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.Hidden = bool(x)
			}
		}
	case "Zoom":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Zoom = v.Value.(float64)
		}
	case "ZoomControlEnabled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.ZoomControlEnabled = bool(x)
			}
		}
	case "EnableDragAndDrop":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.EnableDragAndDrop = bool(x)
			}
		}
	case "OpenInspectorOnStartup":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.OpenInspectorOnStartup = bool(x)
			}
		}
	case "Mac":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MacWindow); ok {
					s.source.Mac = *ptr
				} else {
					s.source.Mac = src.(applicationsrc.MacWindow)
				}
			}
		case *data.AnyValue:
			s.source.Mac = v.Value.(applicationsrc.MacWindow)
		}
	case "Windows":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowsWindow); ok {
					s.source.Windows = *ptr
				} else {
					s.source.Windows = src.(applicationsrc.WindowsWindow)
				}
			}
		case *data.AnyValue:
			s.source.Windows = v.Value.(applicationsrc.WindowsWindow)
		}
	case "Linux":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.LinuxWindow); ok {
					s.source.Linux = *ptr
				} else {
					s.source.Linux = src.(applicationsrc.LinuxWindow)
				}
			}
		case *data.AnyValue:
			s.source.Linux = v.Value.(applicationsrc.LinuxWindow)
		}
	case "MinimiseButtonState":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.MinimiseButtonState = applicationsrc.ButtonState(x)
			}
		}
	case "MaximiseButtonState":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.MaximiseButtonState = applicationsrc.ButtonState(x)
			}
		}
	case "CloseButtonState":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.CloseButtonState = applicationsrc.ButtonState(x)
			}
		}
	case "DevToolsEnabled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DevToolsEnabled = bool(x)
			}
		}
	case "DefaultContextMenuDisabled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DefaultContextMenuDisabled = bool(x)
			}
		}
	case "KeyBindings":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.KeyBindings = v.Value.(map[string]func(applicationsrc.Window))
		}
	case "IgnoreMouseEvents":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.IgnoreMouseEvents = bool(x)
			}
		}
	case "ContentProtectionEnabled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.ContentProtectionEnabled = bool(x)
			}
		}
	}
}

func (s *WebviewWindowOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *WebviewWindowOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *WebviewWindowOptionsClass) GetConstruct() data.Method { return s.construct }

type WebviewWindowOptionsConstructMethod struct {
	source *applicationsrc.WebviewWindowOptions
}

func (h *WebviewWindowOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WebviewWindowOptionsConstructMethod) GetName() string            { return "construct" }
func (h *WebviewWindowOptionsConstructMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowOptionsConstructMethod) GetIsStatic() bool          { return false }
func (h *WebviewWindowOptionsConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *WebviewWindowOptionsConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *WebviewWindowOptionsConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
