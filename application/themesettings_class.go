package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewThemeSettingsClass() data.ClassStmt {
	return &ThemeSettingsClass{
		source:    nil,
		construct: &ThemeSettingsConstructMethod{source: nil},
	}
}

func NewThemeSettingsClassFrom(source *applicationsrc.ThemeSettings) data.ClassStmt {
	return &ThemeSettingsClass{
		source:    source,
		construct: &ThemeSettingsConstructMethod{source: source},
	}
}

type ThemeSettingsClass struct {
	node.Node
	source    *applicationsrc.ThemeSettings
	construct data.Method
}

func (s *ThemeSettingsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewThemeSettingsClassFrom(&applicationsrc.ThemeSettings{}), ctx.CreateBaseContext()), nil
}
func (s *ThemeSettingsClass) GetName() string         { return "wails\\application\\ThemeSettings" }
func (s *ThemeSettingsClass) GetExtend() *string      { return nil }
func (s *ThemeSettingsClass) GetImplements() []string { return nil }
func (s *ThemeSettingsClass) AsString() string        { return "ThemeSettings{}" }
func (s *ThemeSettingsClass) GetSource() any          { return s.source }
func (s *ThemeSettingsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "DarkModeActive":
		return node.NewProperty(nil, "DarkModeActive", "public", true, data.NewClassValue(NewWindowThemeClassFrom(s.source.DarkModeActive), runtime.NewContextToDo())), true
	case "DarkModeInactive":
		return node.NewProperty(nil, "DarkModeInactive", "public", true, data.NewClassValue(NewWindowThemeClassFrom(s.source.DarkModeInactive), runtime.NewContextToDo())), true
	case "LightModeActive":
		return node.NewProperty(nil, "LightModeActive", "public", true, data.NewClassValue(NewWindowThemeClassFrom(s.source.LightModeActive), runtime.NewContextToDo())), true
	case "LightModeInactive":
		return node.NewProperty(nil, "LightModeInactive", "public", true, data.NewClassValue(NewWindowThemeClassFrom(s.source.LightModeInactive), runtime.NewContextToDo())), true
	case "DarkModeMenuBar":
		return node.NewProperty(nil, "DarkModeMenuBar", "public", true, data.NewClassValue(NewMenuBarThemeClassFrom(s.source.DarkModeMenuBar), runtime.NewContextToDo())), true
	case "LightModeMenuBar":
		return node.NewProperty(nil, "LightModeMenuBar", "public", true, data.NewClassValue(NewMenuBarThemeClassFrom(s.source.LightModeMenuBar), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *ThemeSettingsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"DarkModeActive":    node.NewProperty(nil, "DarkModeActive", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"DarkModeInactive":  node.NewProperty(nil, "DarkModeInactive", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"LightModeActive":   node.NewProperty(nil, "LightModeActive", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"LightModeInactive": node.NewProperty(nil, "LightModeInactive", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"DarkModeMenuBar":   node.NewProperty(nil, "DarkModeMenuBar", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
		"LightModeMenuBar":  node.NewProperty(nil, "LightModeMenuBar", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *ThemeSettingsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "DarkModeActive":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowTheme); ok {
					s.source.DarkModeActive = ptr
				}
			}
		case *data.AnyValue:
			s.source.DarkModeActive = v.Value.(*applicationsrc.WindowTheme)
		}
	case "DarkModeInactive":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowTheme); ok {
					s.source.DarkModeInactive = ptr
				}
			}
		case *data.AnyValue:
			s.source.DarkModeInactive = v.Value.(*applicationsrc.WindowTheme)
		}
	case "LightModeActive":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowTheme); ok {
					s.source.LightModeActive = ptr
				}
			}
		case *data.AnyValue:
			s.source.LightModeActive = v.Value.(*applicationsrc.WindowTheme)
		}
	case "LightModeInactive":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.WindowTheme); ok {
					s.source.LightModeInactive = ptr
				}
			}
		case *data.AnyValue:
			s.source.LightModeInactive = v.Value.(*applicationsrc.WindowTheme)
		}
	case "DarkModeMenuBar":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MenuBarTheme); ok {
					s.source.DarkModeMenuBar = ptr
				}
			}
		case *data.AnyValue:
			s.source.DarkModeMenuBar = v.Value.(*applicationsrc.MenuBarTheme)
		}
	case "LightModeMenuBar":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.MenuBarTheme); ok {
					s.source.LightModeMenuBar = ptr
				}
			}
		case *data.AnyValue:
			s.source.LightModeMenuBar = v.Value.(*applicationsrc.MenuBarTheme)
		}
	}
}

func (s *ThemeSettingsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *ThemeSettingsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *ThemeSettingsClass) GetConstruct() data.Method { return s.construct }

type ThemeSettingsConstructMethod struct {
	source *applicationsrc.ThemeSettings
}

func (h *ThemeSettingsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ThemeSettingsConstructMethod) GetName() string               { return "construct" }
func (h *ThemeSettingsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ThemeSettingsConstructMethod) GetIsStatic() bool             { return false }
func (h *ThemeSettingsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ThemeSettingsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ThemeSettingsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
