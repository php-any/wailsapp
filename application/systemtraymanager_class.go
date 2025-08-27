package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewSystemTrayManagerClass() data.ClassStmt {
	return &SystemTrayManagerClass{
		source:    nil,
		construct: &SystemTrayManagerConstructMethod{source: nil},
		new:       &SystemTrayManagerNewMethod{source: nil},
	}
}

func NewSystemTrayManagerClassFrom(source *applicationsrc.SystemTrayManager) data.ClassStmt {
	return &SystemTrayManagerClass{
		source:    source,
		construct: &SystemTrayManagerConstructMethod{source: source},
		new:       &SystemTrayManagerNewMethod{source: source},
	}
}

type SystemTrayManagerClass struct {
	node.Node
	source    *applicationsrc.SystemTrayManager
	construct data.Method
	new       data.Method
}

func (s *SystemTrayManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewSystemTrayManagerClassFrom(&applicationsrc.SystemTrayManager{}), ctx.CreateBaseContext()), nil
}
func (s *SystemTrayManagerClass) GetName() string                            { return "wails\\application\\SystemTrayManager" }
func (s *SystemTrayManagerClass) GetExtend() *string                         { return nil }
func (s *SystemTrayManagerClass) GetImplements() []string                    { return nil }
func (s *SystemTrayManagerClass) AsString() string                           { return "SystemTrayManager{}" }
func (s *SystemTrayManagerClass) GetSource() any                             { return s.source }
func (s *SystemTrayManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *SystemTrayManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *SystemTrayManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "new":
		return s.new, true
	}
	return nil, false
}

func (s *SystemTrayManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.new,
	}
}

func (s *SystemTrayManagerClass) GetConstruct() data.Method { return s.construct }

type SystemTrayManagerConstructMethod struct {
	source *applicationsrc.SystemTrayManager
}

func (h *SystemTrayManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *SystemTrayManagerConstructMethod) GetName() string               { return "construct" }
func (h *SystemTrayManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *SystemTrayManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *SystemTrayManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *SystemTrayManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *SystemTrayManagerConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
