package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMacOptionsClass() data.ClassStmt {
	return &MacOptionsClass{
		source:    nil,
		construct: &MacOptionsConstructMethod{source: nil},
	}
}

func NewMacOptionsClassFrom(source *applicationsrc.MacOptions) data.ClassStmt {
	return &MacOptionsClass{
		source:    source,
		construct: &MacOptionsConstructMethod{source: source},
	}
}

type MacOptionsClass struct {
	node.Node
	source    *applicationsrc.MacOptions
	construct data.Method
}

func (s *MacOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMacOptionsClassFrom(&applicationsrc.MacOptions{}), ctx.CreateBaseContext()), nil
}
func (s *MacOptionsClass) GetName() string         { return "wails\\application\\MacOptions" }
func (s *MacOptionsClass) GetExtend() *string      { return nil }
func (s *MacOptionsClass) GetImplements() []string { return nil }
func (s *MacOptionsClass) AsString() string        { return "MacOptions{}" }
func (s *MacOptionsClass) GetSource() any          { return s.source }
func (s *MacOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "ActivationPolicy":
		return node.NewProperty(nil, "ActivationPolicy", "public", true, data.NewAnyValue(s.source.ActivationPolicy)), true
	case "ApplicationShouldTerminateAfterLastWindowClosed":
		return node.NewProperty(nil, "ApplicationShouldTerminateAfterLastWindowClosed", "public", true, data.NewAnyValue(s.source.ApplicationShouldTerminateAfterLastWindowClosed)), true
	}
	return nil, false
}

func (s *MacOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"ActivationPolicy": node.NewProperty(nil, "ActivationPolicy", "public", true, data.NewAnyValue(nil)),
		"ApplicationShouldTerminateAfterLastWindowClosed": node.NewProperty(nil, "ApplicationShouldTerminateAfterLastWindowClosed", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *MacOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "ActivationPolicy":
		if iv, ok := value.(*data.IntValue); ok {
			if x, err := iv.AsInt(); err == nil {
				s.source.ActivationPolicy = applicationsrc.ActivationPolicy(x)
			}
		}
	case "ApplicationShouldTerminateAfterLastWindowClosed":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.ApplicationShouldTerminateAfterLastWindowClosed = bool(x)
			}
		}
	}
}

func (s *MacOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *MacOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *MacOptionsClass) GetConstruct() data.Method { return s.construct }

type MacOptionsConstructMethod struct {
	source *applicationsrc.MacOptions
}

func (h *MacOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MacOptionsConstructMethod) GetName() string               { return "construct" }
func (h *MacOptionsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MacOptionsConstructMethod) GetIsStatic() bool             { return false }
func (h *MacOptionsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MacOptionsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MacOptionsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
