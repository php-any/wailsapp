package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	events "github.com/wailsapp/wails/v3/pkg/events"
)

func NewWindowsWindowClass() data.ClassStmt {
	return &WindowsWindowClass{
		source:    nil,
		construct: &WindowsWindowConstructMethod{source: nil},
	}
}

func NewWindowsWindowClassFrom(source *applicationsrc.WindowsWindow) data.ClassStmt {
	return &WindowsWindowClass{
		source:    source,
		construct: &WindowsWindowConstructMethod{source: source},
	}
}

type WindowsWindowClass struct {
	node.Node
	source    *applicationsrc.WindowsWindow
	construct data.Method
}

func (s *WindowsWindowClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWindowsWindowClassFrom(&applicationsrc.WindowsWindow{}), ctx.CreateBaseContext()), nil
}
func (s *WindowsWindowClass) GetName() string         { return "wails\\application\\WindowsWindow" }
func (s *WindowsWindowClass) GetExtend() *string      { return nil }
func (s *WindowsWindowClass) GetImplements() []string { return nil }
func (s *WindowsWindowClass) AsString() string        { return "WindowsWindow{}" }
func (s *WindowsWindowClass) GetSource() any          { return s.source }
func (s *WindowsWindowClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "BackdropType":
		return node.NewProperty(nil, "BackdropType", "public", true, data.NewAnyValue(s.source.BackdropType)), true
	case "DisableIcon":
		return node.NewProperty(nil, "DisableIcon", "public", true, data.NewAnyValue(s.source.DisableIcon)), true
	case "Theme":
		return node.NewProperty(nil, "Theme", "public", true, data.NewAnyValue(s.source.Theme)), true
	case "CustomTheme":
		return node.NewProperty(nil, "CustomTheme", "public", true, data.NewClassValue(NewThemeSettingsClassFrom(&s.source.CustomTheme), runtime.NewContextToDo())), true
	case "DisableFramelessWindowDecorations":
		return node.NewProperty(nil, "DisableFramelessWindowDecorations", "public", true, data.NewAnyValue(s.source.DisableFramelessWindowDecorations)), true
	case "WindowMask":
		return node.NewProperty(nil, "WindowMask", "public", true, data.NewAnyValue(s.source.WindowMask)), true
	case "WindowMaskDraggable":
		return node.NewProperty(nil, "WindowMaskDraggable", "public", true, data.NewAnyValue(s.source.WindowMaskDraggable)), true
	case "ResizeDebounceMS":
		return node.NewProperty(nil, "ResizeDebounceMS", "public", true, data.NewAnyValue(s.source.ResizeDebounceMS)), true
	case "WindowDidMoveDebounceMS":
		return node.NewProperty(nil, "WindowDidMoveDebounceMS", "public", true, data.NewAnyValue(s.source.WindowDidMoveDebounceMS)), true
	case "EventMapping":
		return node.NewProperty(nil, "EventMapping", "public", true, data.NewAnyValue(s.source.EventMapping)), true
	case "HiddenOnTaskbar":
		return node.NewProperty(nil, "HiddenOnTaskbar", "public", true, data.NewAnyValue(s.source.HiddenOnTaskbar)), true
	case "EnableSwipeGestures":
		return node.NewProperty(nil, "EnableSwipeGestures", "public", true, data.NewAnyValue(s.source.EnableSwipeGestures)), true
	case "Menu":
		return node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(NewMenuClassFrom(s.source.Menu), runtime.NewContextToDo())), true
	case "OnEnterEffect":
		return node.NewProperty(nil, "OnEnterEffect", "public", true, data.NewAnyValue(s.source.OnEnterEffect)), true
	case "OnOverEffect":
		return node.NewProperty(nil, "OnOverEffect", "public", true, data.NewAnyValue(s.source.OnOverEffect)), true
	case "Permissions":
		return node.NewProperty(nil, "Permissions", "public", true, data.NewAnyValue(s.source.Permissions)), true
	case "ExStyle":
		return node.NewProperty(nil, "ExStyle", "public", true, data.NewAnyValue(s.source.ExStyle)), true
	case "GeneralAutofillEnabled":
		return node.NewProperty(nil, "GeneralAutofillEnabled", "public", true, data.NewAnyValue(s.source.GeneralAutofillEnabled)), true
	case "PasswordAutosaveEnabled":
		return node.NewProperty(nil, "PasswordAutosaveEnabled", "public", true, data.NewAnyValue(s.source.PasswordAutosaveEnabled)), true
	case "EnabledFeatures":
		return node.NewProperty(nil, "EnabledFeatures", "public", true, data.NewAnyValue(s.source.EnabledFeatures)), true
	case "DisabledFeatures":
		return node.NewProperty(nil, "DisabledFeatures", "public", true, data.NewAnyValue(s.source.DisabledFeatures)), true
	case "AdditionalLaunchArgs":
		return node.NewProperty(nil, "AdditionalLaunchArgs", "public", true, data.NewAnyValue(s.source.AdditionalLaunchArgs)), true
	}
	return nil, false
}

func (s *WindowsWindowClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"BackdropType":                      node.NewProperty(nil, "BackdropType", "public", true, data.NewAnyValue(nil)),
		"DisableIcon":                       node.NewProperty(nil, "DisableIcon", "public", true, data.NewAnyValue(nil)),
		"Theme":                             node.NewProperty(nil, "Theme", "public", true, data.NewAnyValue(nil)),
		"CustomTheme":                       node.NewProperty(nil, "CustomTheme", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"DisableFramelessWindowDecorations": node.NewProperty(nil, "DisableFramelessWindowDecorations", "public", true, data.NewAnyValue(nil)),
		"WindowMask":                        node.NewProperty(nil, "WindowMask", "public", true, data.NewAnyValue(nil)),
		"WindowMaskDraggable":               node.NewProperty(nil, "WindowMaskDraggable", "public", true, data.NewAnyValue(nil)),
		"ResizeDebounceMS":                  node.NewProperty(nil, "ResizeDebounceMS", "public", true, data.NewAnyValue(nil)),
		"WindowDidMoveDebounceMS":           node.NewProperty(nil, "WindowDidMoveDebounceMS", "public", true, data.NewAnyValue(nil)),
		"EventMapping":                      node.NewProperty(nil, "EventMapping", "public", true, data.NewAnyValue(nil)),
		"HiddenOnTaskbar":                   node.NewProperty(nil, "HiddenOnTaskbar", "public", true, data.NewAnyValue(nil)),
		"EnableSwipeGestures":               node.NewProperty(nil, "EnableSwipeGestures", "public", true, data.NewAnyValue(nil)),
		"Menu":                              node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"OnEnterEffect":                     node.NewProperty(nil, "OnEnterEffect", "public", true, data.NewAnyValue(nil)),
		"OnOverEffect":                      node.NewProperty(nil, "OnOverEffect", "public", true, data.NewAnyValue(nil)),
		"Permissions":                       node.NewProperty(nil, "Permissions", "public", true, data.NewAnyValue(nil)),
		"ExStyle":                           node.NewProperty(nil, "ExStyle", "public", true, data.NewAnyValue(nil)),
		"GeneralAutofillEnabled":            node.NewProperty(nil, "GeneralAutofillEnabled", "public", true, data.NewAnyValue(nil)),
		"PasswordAutosaveEnabled":           node.NewProperty(nil, "PasswordAutosaveEnabled", "public", true, data.NewAnyValue(nil)),
		"EnabledFeatures":                   node.NewProperty(nil, "EnabledFeatures", "public", true, data.NewAnyValue(nil)),
		"DisabledFeatures":                  node.NewProperty(nil, "DisabledFeatures", "public", true, data.NewAnyValue(nil)),
		"AdditionalLaunchArgs":              node.NewProperty(nil, "AdditionalLaunchArgs", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *WindowsWindowClass) SetProperty(name string, value data.Value) {
	switch name {
	case "BackdropType":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.BackdropType = v.Value.(applicationsrc.BackdropType)
		}
	case "DisableIcon":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableIcon = bool(x)
			}
		}
	case "Theme":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Theme = applicationsrc.Theme(x)
			}
		}
	case "CustomTheme":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.ThemeSettings); ok {
					s.source.CustomTheme = *ptr
				} else {
					s.source.CustomTheme = src.(applicationsrc.ThemeSettings)
				}
			}
		case *data.AnyValue:
			s.source.CustomTheme = v.Value.(applicationsrc.ThemeSettings)
		}
	case "DisableFramelessWindowDecorations":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableFramelessWindowDecorations = bool(x)
			}
		}
	case "WindowMask":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.WindowMask = v.Value.([]uint8)
		}
	case "WindowMaskDraggable":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.WindowMaskDraggable = bool(x)
			}
		}
	case "ResizeDebounceMS":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.ResizeDebounceMS = v.Value.(uint16)
		}
	case "WindowDidMoveDebounceMS":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.WindowDidMoveDebounceMS = v.Value.(uint16)
		}
	case "EventMapping":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.EventMapping = v.Value.(map[events.WindowEventType]events.WindowEventType)
		}
	case "HiddenOnTaskbar":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.HiddenOnTaskbar = bool(x)
			}
		}
	case "EnableSwipeGestures":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.EnableSwipeGestures = bool(x)
			}
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
	case "OnEnterEffect":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.OnEnterEffect = v.Value.(applicationsrc.DragEffect)
		}
	case "OnOverEffect":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.OnOverEffect = v.Value.(applicationsrc.DragEffect)
		}
	case "Permissions":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.Permissions = v.Value.(map[applicationsrc.CoreWebView2PermissionKind]applicationsrc.CoreWebView2PermissionState)
		}
	case "ExStyle":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.ExStyle = int(x)
			}
		}
	case "GeneralAutofillEnabled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.GeneralAutofillEnabled = bool(x)
			}
		}
	case "PasswordAutosaveEnabled":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.PasswordAutosaveEnabled = bool(x)
			}
		}
	case "EnabledFeatures":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.EnabledFeatures = v.Value.([]string)
		}
	case "DisabledFeatures":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.DisabledFeatures = v.Value.([]string)
		}
	case "AdditionalLaunchArgs":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.AdditionalLaunchArgs = v.Value.([]string)
		}
	}
}

func (s *WindowsWindowClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *WindowsWindowClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *WindowsWindowClass) GetConstruct() data.Method { return s.construct }

type WindowsWindowConstructMethod struct {
	source *applicationsrc.WindowsWindow
}

func (h *WindowsWindowConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowsWindowConstructMethod) GetName() string               { return "construct" }
func (h *WindowsWindowConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowsWindowConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowsWindowConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowsWindowConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowsWindowConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
