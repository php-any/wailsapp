package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewEnvironmentInfoClass() data.ClassStmt {
	return &EnvironmentInfoClass{
		source:    nil,
		construct: &EnvironmentInfoConstructMethod{source: nil},
	}
}

func NewEnvironmentInfoClassFrom(source *applicationsrc.EnvironmentInfo) data.ClassStmt {
	return &EnvironmentInfoClass{
		source:    source,
		construct: &EnvironmentInfoConstructMethod{source: source},
	}
}

type EnvironmentInfoClass struct {
	node.Node
	source    *applicationsrc.EnvironmentInfo
	construct data.Method
}

func (s *EnvironmentInfoClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewEnvironmentInfoClassFrom(&applicationsrc.EnvironmentInfo{}), ctx.CreateBaseContext()), nil
}
func (s *EnvironmentInfoClass) GetName() string         { return "wails\\application\\EnvironmentInfo" }
func (s *EnvironmentInfoClass) GetExtend() *string      { return nil }
func (s *EnvironmentInfoClass) GetImplements() []string { return nil }
func (s *EnvironmentInfoClass) AsString() string        { return "EnvironmentInfo{}" }
func (s *EnvironmentInfoClass) GetSource() any          { return s.source }
func (s *EnvironmentInfoClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "OS":
		return node.NewProperty(nil, "OS", "public", true, data.NewAnyValue(s.source.OS)), true
	case "Arch":
		return node.NewProperty(nil, "Arch", "public", true, data.NewAnyValue(s.source.Arch)), true
	case "Debug":
		return node.NewProperty(nil, "Debug", "public", true, data.NewAnyValue(s.source.Debug)), true
	case "OSInfo":
		return node.NewProperty(nil, "OSInfo", "public", true, data.NewAnyValue(s.source.OSInfo)), true
	case "PlatformInfo":
		return node.NewProperty(nil, "PlatformInfo", "public", true, data.NewAnyValue(s.source.PlatformInfo)), true
	}
	return nil, false
}

func (s *EnvironmentInfoClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"OS":           node.NewProperty(nil, "OS", "public", true, data.NewAnyValue(nil)),
		"Arch":         node.NewProperty(nil, "Arch", "public", true, data.NewAnyValue(nil)),
		"Debug":        node.NewProperty(nil, "Debug", "public", true, data.NewAnyValue(nil)),
		"OSInfo":       node.NewProperty(nil, "OSInfo", "public", true, data.NewAnyValue(nil)),
		"PlatformInfo": node.NewProperty(nil, "PlatformInfo", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *EnvironmentInfoClass) SetProperty(name string, value data.Value) {
	switch name {
	case "OS":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.OS = string(sv.AsString())
		}
	case "Arch":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.Arch = string(sv.AsString())
		}
	case "Debug":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.Debug = bool(x)
			}
		}
	case "OSInfo":

	case "PlatformInfo":
		if v, ok := value.(*data.AnyValue); ok {
			s.source.PlatformInfo = v.Value.(map[string]interface{})
		}
	}
}

func (s *EnvironmentInfoClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *EnvironmentInfoClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *EnvironmentInfoClass) GetConstruct() data.Method { return s.construct }

type EnvironmentInfoConstructMethod struct {
	source *applicationsrc.EnvironmentInfo
}

func (h *EnvironmentInfoConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *EnvironmentInfoConstructMethod) GetName() string               { return "construct" }
func (h *EnvironmentInfoConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *EnvironmentInfoConstructMethod) GetIsStatic() bool             { return false }
func (h *EnvironmentInfoConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *EnvironmentInfoConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *EnvironmentInfoConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
