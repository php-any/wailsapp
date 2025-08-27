package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewLinuxOptionsClass() data.ClassStmt {
	return &LinuxOptionsClass{
		source:    nil,
		construct: &LinuxOptionsConstructMethod{source: nil},
	}
}

func NewLinuxOptionsClassFrom(source *applicationsrc.LinuxOptions) data.ClassStmt {
	return &LinuxOptionsClass{
		source:    source,
		construct: &LinuxOptionsConstructMethod{source: source},
	}
}

type LinuxOptionsClass struct {
	node.Node
	source    *applicationsrc.LinuxOptions
	construct data.Method
}

func (s *LinuxOptionsClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewLinuxOptionsClassFrom(&applicationsrc.LinuxOptions{}), ctx.CreateBaseContext()), nil
}
func (s *LinuxOptionsClass) GetName() string         { return "wails\\application\\LinuxOptions" }
func (s *LinuxOptionsClass) GetExtend() *string      { return nil }
func (s *LinuxOptionsClass) GetImplements() []string { return nil }
func (s *LinuxOptionsClass) AsString() string        { return "LinuxOptions{}" }
func (s *LinuxOptionsClass) GetSource() any          { return s.source }
func (s *LinuxOptionsClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "DisableQuitOnLastWindowClosed":
		return node.NewProperty(nil, "DisableQuitOnLastWindowClosed", "public", true, data.NewAnyValue(s.source.DisableQuitOnLastWindowClosed)), true
	case "ProgramName":
		return node.NewProperty(nil, "ProgramName", "public", true, data.NewAnyValue(s.source.ProgramName)), true
	}
	return nil, false
}

func (s *LinuxOptionsClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"DisableQuitOnLastWindowClosed": node.NewProperty(nil, "DisableQuitOnLastWindowClosed", "public", true, data.NewAnyValue(nil)),
		"ProgramName":                   node.NewProperty(nil, "ProgramName", "public", true, data.NewAnyValue(nil)),
	}
}

func (s *LinuxOptionsClass) SetProperty(name string, value data.Value) {
	switch name {
	case "DisableQuitOnLastWindowClosed":
		if bv, ok := value.(*data.BoolValue); ok {
			if x, err := bv.AsBool(); err == nil {
				s.source.DisableQuitOnLastWindowClosed = bool(x)
			}
		}
	case "ProgramName":
		if sv, ok := value.(*data.StringValue); ok {
			s.source.ProgramName = string(sv.AsString())
		}
	}
}

func (s *LinuxOptionsClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	}
	return nil, false
}

func (s *LinuxOptionsClass) GetMethods() []data.Method {
	return []data.Method{}
}

func (s *LinuxOptionsClass) GetConstruct() data.Method { return s.construct }

type LinuxOptionsConstructMethod struct {
	source *applicationsrc.LinuxOptions
}

func (h *LinuxOptionsConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *LinuxOptionsConstructMethod) GetName() string               { return "construct" }
func (h *LinuxOptionsConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *LinuxOptionsConstructMethod) GetIsStatic() bool             { return false }
func (h *LinuxOptionsConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *LinuxOptionsConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *LinuxOptionsConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
