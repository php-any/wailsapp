package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	events "github.com/wailsapp/wails/v3/pkg/events"
)

func NewMacWindowClass() data.ClassStmt {
	return &MacWindowClass{
		source:    nil,
		construct: &MacWindowConstructMethod{source: nil},
	}
}

func NewMacWindowClassFrom(source *applicationsrc.MacWindow) data.ClassStmt {
	return &MacWindowClass{
		source:    source,
		construct: &MacWindowConstructMethod{source: source},
	}
}

type MacWindowClass struct {
	node.Node
	source    *applicationsrc.MacWindow
	construct data.Method
}

func (s *MacWindowClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMacWindowClassFrom(&applicationsrc.MacWindow{}), ctx.CreateBaseContext()), nil
}
func (s *MacWindowClass) GetName() string         { return "wails\\application\\MacWindow" }
func (s *MacWindowClass) GetExtend() *string      { return nil }
func (s *MacWindowClass) GetImplements() []string { return nil }
func (s *MacWindowClass) AsString() string        { return "MacWindow{}" }
func (s *MacWindowClass) GetSource() any          { return s.source }
func (s *MacWindowClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Backdrop":
		return node.NewProperty(nil, "Backdrop", "public", true, data.NewAnyValue(s.source.Backdrop)), true
	case "DisableShadow":
		return node.NewProperty(nil, "DisableShadow", "public", true, data.NewAnyValue(s.source.DisableShadow)), true
	case "TitleBar":
		return node.NewProperty(nil, "TitleBar", "public", true, data.NewClassValue(NewMacTitleBarClassFrom(&s.source.TitleBar), runtime.NewContextToDo())), true
	case "Appearance":
		return node.NewProperty(nil, "Appearance", "public", true, data.NewAnyValue(s.source.Appearance)), true
	case "InvisibleTitleBarHeight":
		return node.NewProperty(nil, "InvisibleTitleBarHeight", "public", true, data.NewAnyValue(s.source.InvisibleTitleBarHeight)), true
	case "EventMapping":
		return node.NewProperty(nil, "EventMapping", "public", true, data.NewAnyValue(s.source.EventMapping)), true
	case "EnableFraudulentWebsiteWarnings":
		return node.NewProperty(nil, "EnableFraudulentWebsiteWarnings", "public", true, data.NewAnyValue(s.source.EnableFraudulentWebsiteWarnings)), true
	case "WebviewPreferences":
		return node.NewProperty(nil, "WebviewPreferences", "public", true, data.NewClassValue(NewMacWebviewPreferencesClassFrom(&s.source.WebviewPreferences), runtime.NewContextToDo())), true
	case "WindowLevel":
		return node.NewProperty(nil, "WindowLevel", "public", true, data.NewAnyValue(s.source.WindowLevel)), true
	case "LiquidGlass":
		return node.NewProperty(nil, "LiquidGlass", "public", true, data.NewClassValue(NewMacLiquidGlassClassFrom(&s.source.LiquidGlass), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *MacWindowClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Backdrop":                        node.NewProperty(nil, "Backdrop", "public", true, data.NewAnyValue(nil)),
		"DisableShadow":                   node.NewProperty(nil, "DisableShadow", "public", true, data.NewAnyValue(nil)),
		"TitleBar":                        node.NewProperty(nil, "TitleBar", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"Appearance":                      node.NewProperty(nil, "Appearance", "public", true, data.NewAnyValue(nil)),
		"InvisibleTitleBarHeight":         node.NewProperty(nil, "InvisibleTitleBarHeight", "public", true, data.NewAnyValue(nil)),
		"EventMapping":                    node.NewProperty(nil, "EventMapping", "public", true, data.NewAnyValue(nil)),
		"EnableFraudulentWebsiteWarnings": node.NewProperty(nil, "EnableFraudulentWebsiteWarnings", "public", true, data.NewAnyValue(nil)),
		"WebviewPreferences":              node.NewProperty(nil, "WebviewPreferences", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"WindowLevel":                     node.NewProperty(nil, "WindowLevel", "public", true, data.NewAnyValue(nil)),
		"LiquidGlass":                     node.NewProperty(nil, "LiquidGlass", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *MacWindowClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Backdrop":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.Backdrop = applicationsrc.MacBackdrop(x)
			}
		}
	case "DisableShadow":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableShadow = bool(x)
			}
		}
	case "TitleBar":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MacTitleBar); ok {
					s.source.TitleBar = *ptr
				} else {
					s.source.TitleBar = src.(applicationsrc.MacTitleBar)
				}
			}
		case *data.AnyValue:
			s.source.TitleBar = v.Value.(applicationsrc.MacTitleBar)
		}
	case "Appearance":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Appearance = applicationsrc.MacAppearanceType(sv.AsString())
		}
	case "InvisibleTitleBarHeight":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.InvisibleTitleBarHeight = int(x)
			}
		}
	case "EventMapping":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.EventMapping = v.Value.(map[events.WindowEventType]events.WindowEventType)
		}
	case "EnableFraudulentWebsiteWarnings":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.EnableFraudulentWebsiteWarnings = bool(x)
			}
		}
	case "WebviewPreferences":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MacWebviewPreferences); ok {
					s.source.WebviewPreferences = *ptr
				} else {
					s.source.WebviewPreferences = src.(applicationsrc.MacWebviewPreferences)
				}
			}
		case *data.AnyValue:
			s.source.WebviewPreferences = v.Value.(applicationsrc.MacWebviewPreferences)
		}
	case "WindowLevel":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.WindowLevel = applicationsrc.MacWindowLevel(sv.AsString())
		}
	case "LiquidGlass":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MacLiquidGlass); ok {
					s.source.LiquidGlass = *ptr
				} else {
					s.source.LiquidGlass = src.(applicationsrc.MacLiquidGlass)
				}
			}
		case *data.AnyValue:
			s.source.LiquidGlass = v.Value.(applicationsrc.MacLiquidGlass)
		}
	}
}

func (s *MacWindowClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MacWindowClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MacWindowClass) GetConstruct() data.Method { return s.construct }

type MacWindowConstructMethod struct {
	source *applicationsrc.MacWindow
}

func (h *MacWindowConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MacWindowConstructMethod) GetName() string               { return "construct" }
func (h *MacWindowConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MacWindowConstructMethod) GetIsStatic() bool             { return false }
func (h *MacWindowConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MacWindowConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MacWindowConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
