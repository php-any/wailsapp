package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewEnvironmentManagerClass() data.ClassStmt {
	return &EnvironmentManagerClass{
		source:          nil,
		construct:       &EnvironmentManagerConstructMethod{source: nil},
		getAccentColor:  &EnvironmentManagerGetAccentColorMethod{source: nil},
		info:            &EnvironmentManagerInfoMethod{source: nil},
		isDarkMode:      &EnvironmentManagerIsDarkModeMethod{source: nil},
		openFileManager: &EnvironmentManagerOpenFileManagerMethod{source: nil},
	}
}

func NewEnvironmentManagerClassFrom(source *applicationsrc.EnvironmentManager) data.ClassStmt {
	return &EnvironmentManagerClass{
		source:          source,
		construct:       &EnvironmentManagerConstructMethod{source: source},
		getAccentColor:  &EnvironmentManagerGetAccentColorMethod{source: source},
		info:            &EnvironmentManagerInfoMethod{source: source},
		isDarkMode:      &EnvironmentManagerIsDarkModeMethod{source: source},
		openFileManager: &EnvironmentManagerOpenFileManagerMethod{source: source},
	}
}

type EnvironmentManagerClass struct {
	node.Node
	source          *applicationsrc.EnvironmentManager
	construct       data.Method
	getAccentColor  data.Method
	info            data.Method
	isDarkMode      data.Method
	openFileManager data.Method
}

func (s *EnvironmentManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewEnvironmentManagerClassFrom(&applicationsrc.EnvironmentManager{}), ctx.CreateBaseContext()), nil
}
func (s *EnvironmentManagerClass) GetName() string                            { return "wails\\application\\EnvironmentManager" }
func (s *EnvironmentManagerClass) GetExtend() *string                         { return nil }
func (s *EnvironmentManagerClass) GetImplements() []string                    { return nil }
func (s *EnvironmentManagerClass) AsString() string                           { return "EnvironmentManager{}" }
func (s *EnvironmentManagerClass) GetSource() any                             { return s.source }
func (s *EnvironmentManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *EnvironmentManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *EnvironmentManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "getAccentColor":
		return s.getAccentColor, true
	case "info":
		return s.info, true
	case "isDarkMode":
		return s.isDarkMode, true
	case "openFileManager":
		return s.openFileManager, true
	}
	return nil, false
}

func (s *EnvironmentManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.getAccentColor,
		s.info,
		s.isDarkMode,
		s.openFileManager,
	}
}

func (s *EnvironmentManagerClass) GetConstruct() data.Method { return s.construct }

type EnvironmentManagerConstructMethod struct {
	source *applicationsrc.EnvironmentManager
}

func (h *EnvironmentManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *EnvironmentManagerConstructMethod) GetName() string               { return "construct" }
func (h *EnvironmentManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *EnvironmentManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *EnvironmentManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *EnvironmentManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *EnvironmentManagerConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
