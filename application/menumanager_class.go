package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMenuManagerClass() data.ClassStmt {
	return &MenuManagerClass{
		source:             nil,
		construct:          &MenuManagerConstructMethod{source: nil},
		getApplicationMenu: &MenuManagerGetApplicationMenuMethod{source: nil},
		new:                &MenuManagerNewMethod{source: nil},
		set:                &MenuManagerSetMethod{source: nil},
		setApplicationMenu: &MenuManagerSetApplicationMenuMethod{source: nil},
		showAbout:          &MenuManagerShowAboutMethod{source: nil},
	}
}

func NewMenuManagerClassFrom(source *applicationsrc.MenuManager) data.ClassStmt {
	return &MenuManagerClass{
		source:             source,
		construct:          &MenuManagerConstructMethod{source: source},
		getApplicationMenu: &MenuManagerGetApplicationMenuMethod{source: source},
		new:                &MenuManagerNewMethod{source: source},
		set:                &MenuManagerSetMethod{source: source},
		setApplicationMenu: &MenuManagerSetApplicationMenuMethod{source: source},
		showAbout:          &MenuManagerShowAboutMethod{source: source},
	}
}

type MenuManagerClass struct {
	node.Node
	source             *applicationsrc.MenuManager
	construct          data.Method
	getApplicationMenu data.Method
	new                data.Method
	set                data.Method
	setApplicationMenu data.Method
	showAbout          data.Method
}

func (s *MenuManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMenuManagerClassFrom(&applicationsrc.MenuManager{}), ctx.CreateBaseContext()), nil
}
func (s *MenuManagerClass) GetName() string                            { return "wails\\application\\MenuManager" }
func (s *MenuManagerClass) GetExtend() *string                         { return nil }
func (s *MenuManagerClass) GetImplements() []string                    { return nil }
func (s *MenuManagerClass) AsString() string                           { return "MenuManager{}" }
func (s *MenuManagerClass) GetSource() any                             { return s.source }
func (s *MenuManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *MenuManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *MenuManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "getApplicationMenu":
		return s.getApplicationMenu, true
	case "new":
		return s.new, true
	case "set":
		return s.set, true
	case "setApplicationMenu":
		return s.setApplicationMenu, true
	case "showAbout":
		return s.showAbout, true
	}
	return nil, false
}

func (s *MenuManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.getApplicationMenu,
		s.new,
		s.set,
		s.setApplicationMenu,
		s.showAbout,
	}
}

func (s *MenuManagerClass) GetConstruct() data.Method { return s.construct }

type MenuManagerConstructMethod struct {
	source *applicationsrc.MenuManager
}

func (h *MenuManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MenuManagerConstructMethod) GetName() string               { return "construct" }
func (h *MenuManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MenuManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *MenuManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MenuManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MenuManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
