package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMacTitleBarClass() data.ClassStmt {
	return &MacTitleBarClass{
		source:    nil,
		construct: &MacTitleBarConstructMethod{source: nil},
	}
}

func NewMacTitleBarClassFrom(source *applicationsrc.MacTitleBar) data.ClassStmt {
	return &MacTitleBarClass{
		source:    source,
		construct: &MacTitleBarConstructMethod{source: source},
	}
}

type MacTitleBarClass struct {
	node.Node
	source    *applicationsrc.MacTitleBar
	construct data.Method
}

func (s *MacTitleBarClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMacTitleBarClassFrom(&applicationsrc.MacTitleBar{}), ctx.CreateBaseContext()), nil
}
func (s *MacTitleBarClass) GetName() string         { return "wails\\application\\MacTitleBar" }
func (s *MacTitleBarClass) GetExtend() *string      { return nil }
func (s *MacTitleBarClass) GetImplements() []string { return nil }
func (s *MacTitleBarClass) AsString() string        { return "MacTitleBar{}" }
func (s *MacTitleBarClass) GetSource() any          { return s.source }
func (s *MacTitleBarClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "AppearsTransparent":
		return node.NewProperty(nil, "AppearsTransparent", "public", true, data.NewAnyValue(s.source.AppearsTransparent)), true
	case "Hide":
		return node.NewProperty(nil, "Hide", "public", true, data.NewAnyValue(s.source.Hide)), true
	case "HideTitle":
		return node.NewProperty(nil, "HideTitle", "public", true, data.NewAnyValue(s.source.HideTitle)), true
	case "FullSizeContent":
		return node.NewProperty(nil, "FullSizeContent", "public", true, data.NewAnyValue(s.source.FullSizeContent)), true
	case "UseToolbar":
		return node.NewProperty(nil, "UseToolbar", "public", true, data.NewAnyValue(s.source.UseToolbar)), true
	case "HideToolbarSeparator":
		return node.NewProperty(nil, "HideToolbarSeparator", "public", true, data.NewAnyValue(s.source.HideToolbarSeparator)), true
	case "ShowToolbarWhenFullscreen":
		return node.NewProperty(nil, "ShowToolbarWhenFullscreen", "public", true, data.NewAnyValue(s.source.ShowToolbarWhenFullscreen)), true
	case "ToolbarStyle":
		return node.NewProperty(nil, "ToolbarStyle", "public", true, data.NewAnyValue(s.source.ToolbarStyle)), true
	}
	return nil, false
}

func (s *MacTitleBarClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"AppearsTransparent":        node.NewProperty(nil, "AppearsTransparent", "public", true, data.NewAnyValue(nil)),
		"Hide":                      node.NewProperty(nil, "Hide", "public", true, data.NewAnyValue(nil)),
		"HideTitle":                 node.NewProperty(nil, "HideTitle", "public", true, data.NewAnyValue(nil)),
		"FullSizeContent":           node.NewProperty(nil, "FullSizeContent", "public", true, data.NewAnyValue(nil)),
		"UseToolbar":                node.NewProperty(nil, "UseToolbar", "public", true, data.NewAnyValue(nil)),
		"HideToolbarSeparator":      node.NewProperty(nil, "HideToolbarSeparator", "public", true, data.NewAnyValue(nil)),
		"ShowToolbarWhenFullscreen": node.NewProperty(nil, "ShowToolbarWhenFullscreen", "public", true, data.NewAnyValue(nil)),
		"ToolbarStyle":              node.NewProperty(nil, "ToolbarStyle", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *MacTitleBarClass) SetProperty(name string, value data.Value) {
	switch name {
	case "AppearsTransparent":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.AppearsTransparent = bool(x)
			}
		}
	case "Hide":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.Hide = bool(x)
			}
		}
	case "HideTitle":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.HideTitle = bool(x)
			}
		}
	case "FullSizeContent":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.FullSizeContent = bool(x)
			}
		}
	case "UseToolbar":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.UseToolbar = bool(x)
			}
		}
	case "HideToolbarSeparator":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.HideToolbarSeparator = bool(x)
			}
		}
	case "ShowToolbarWhenFullscreen":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.ShowToolbarWhenFullscreen = bool(x)
			}
		}
	case "ToolbarStyle":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.ToolbarStyle = applicationsrc.MacToolbarStyle(x)
			}
		}
	}
}

func (s *MacTitleBarClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MacTitleBarClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MacTitleBarClass) GetConstruct() data.Method { return s.construct }

type MacTitleBarConstructMethod struct {
	source *applicationsrc.MacTitleBar
}

func (h *MacTitleBarConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MacTitleBarConstructMethod) GetName() string               { return "construct" }
func (h *MacTitleBarConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MacTitleBarConstructMethod) GetIsStatic() bool             { return false }
func (h *MacTitleBarConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MacTitleBarConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MacTitleBarConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
