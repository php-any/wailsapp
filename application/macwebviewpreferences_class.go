package application

import (
	u "github.com/leaanthony/u"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMacWebviewPreferencesClass() data.ClassStmt {
	return &MacWebviewPreferencesClass{
		source:    nil,
		construct: &MacWebviewPreferencesConstructMethod{source: nil},
	}
}

func NewMacWebviewPreferencesClassFrom(source *applicationsrc.MacWebviewPreferences) data.ClassStmt {
	return &MacWebviewPreferencesClass{
		source:    source,
		construct: &MacWebviewPreferencesConstructMethod{source: source},
	}
}

type MacWebviewPreferencesClass struct {
	node.Node
	source    *applicationsrc.MacWebviewPreferences
	construct data.Method
}

func (s *MacWebviewPreferencesClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMacWebviewPreferencesClassFrom(&applicationsrc.MacWebviewPreferences{}), ctx.CreateBaseContext()), nil
}
func (s *MacWebviewPreferencesClass) GetName() string {
	return "wails\\application\\MacWebviewPreferences"
}
func (s *MacWebviewPreferencesClass) GetExtend() *string      { return nil }
func (s *MacWebviewPreferencesClass) GetImplements() []string { return nil }
func (s *MacWebviewPreferencesClass) AsString() string        { return "MacWebviewPreferences{}" }
func (s *MacWebviewPreferencesClass) GetSource() any          { return s.source }
func (s *MacWebviewPreferencesClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "TabFocusesLinks":
		return node.NewProperty(nil, "TabFocusesLinks", "public", true, data.NewAnyValue(s.source.TabFocusesLinks)), true
	case "TextInteractionEnabled":
		return node.NewProperty(nil, "TextInteractionEnabled", "public", true, data.NewAnyValue(s.source.TextInteractionEnabled)), true
	case "FullscreenEnabled":
		return node.NewProperty(nil, "FullscreenEnabled", "public", true, data.NewAnyValue(s.source.FullscreenEnabled)), true
	case "AllowsBackForwardNavigationGestures":
		return node.NewProperty(nil, "AllowsBackForwardNavigationGestures", "public", true, data.NewAnyValue(s.source.AllowsBackForwardNavigationGestures)), true
	}
	return nil, false
}

func (s *MacWebviewPreferencesClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"TabFocusesLinks":                     node.NewProperty(nil, "TabFocusesLinks", "public", true, data.NewAnyValue(nil)),
		"TextInteractionEnabled":              node.NewProperty(nil, "TextInteractionEnabled", "public", true, data.NewAnyValue(nil)),
		"FullscreenEnabled":                   node.NewProperty(nil, "FullscreenEnabled", "public", true, data.NewAnyValue(nil)),
		"AllowsBackForwardNavigationGestures": node.NewProperty(nil, "AllowsBackForwardNavigationGestures", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *MacWebviewPreferencesClass) SetProperty(name string, value data.Value) {
	switch name {
	case "TabFocusesLinks":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.TabFocusesLinks = v.Value.(u.Var[bool])
		}
	case "TextInteractionEnabled":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.TextInteractionEnabled = v.Value.(u.Var[bool])
		}
	case "FullscreenEnabled":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.FullscreenEnabled = v.Value.(u.Var[bool])
		}
	case "AllowsBackForwardNavigationGestures":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.AllowsBackForwardNavigationGestures = v.Value.(u.Var[bool])
		}
	}
}

func (s *MacWebviewPreferencesClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MacWebviewPreferencesClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MacWebviewPreferencesClass) GetConstruct() data.Method { return s.construct }

type MacWebviewPreferencesConstructMethod struct {
	source *applicationsrc.MacWebviewPreferences
}

func (h *MacWebviewPreferencesConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MacWebviewPreferencesConstructMethod) GetName() string { return "construct" }
func (h *MacWebviewPreferencesConstructMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *MacWebviewPreferencesConstructMethod) GetIsStatic() bool          { return false }
func (h *MacWebviewPreferencesConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *MacWebviewPreferencesConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *MacWebviewPreferencesConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
