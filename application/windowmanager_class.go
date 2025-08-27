package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWindowManagerClass() data.ClassStmt {
	return &WindowManagerClass{
		source:         nil,
		construct:      &WindowManagerConstructMethod{source: nil},
		add:            &WindowManagerAddMethod{source: nil},
		current:        &WindowManagerCurrentMethod{source: nil},
		get:            &WindowManagerGetMethod{source: nil},
		getAll:         &WindowManagerGetAllMethod{source: nil},
		getByID:        &WindowManagerGetByIDMethod{source: nil},
		getByName:      &WindowManagerGetByNameMethod{source: nil},
		new:            &WindowManagerNewMethod{source: nil},
		newWithOptions: &WindowManagerNewWithOptionsMethod{source: nil},
		onCreate:       &WindowManagerOnCreateMethod{source: nil},
		remove:         &WindowManagerRemoveMethod{source: nil},
		removeByName:   &WindowManagerRemoveByNameMethod{source: nil},
	}
}

func NewWindowManagerClassFrom(source *applicationsrc.WindowManager) data.ClassStmt {
	return &WindowManagerClass{
		source:         source,
		construct:      &WindowManagerConstructMethod{source: source},
		add:            &WindowManagerAddMethod{source: source},
		current:        &WindowManagerCurrentMethod{source: source},
		get:            &WindowManagerGetMethod{source: source},
		getAll:         &WindowManagerGetAllMethod{source: source},
		getByID:        &WindowManagerGetByIDMethod{source: source},
		getByName:      &WindowManagerGetByNameMethod{source: source},
		new:            &WindowManagerNewMethod{source: source},
		newWithOptions: &WindowManagerNewWithOptionsMethod{source: source},
		onCreate:       &WindowManagerOnCreateMethod{source: source},
		remove:         &WindowManagerRemoveMethod{source: source},
		removeByName:   &WindowManagerRemoveByNameMethod{source: source},
	}
}

type WindowManagerClass struct {
	node.Node
	source         *applicationsrc.WindowManager
	construct      data.Method
	add            data.Method
	current        data.Method
	get            data.Method
	getAll         data.Method
	getByID        data.Method
	getByName      data.Method
	new            data.Method
	newWithOptions data.Method
	onCreate       data.Method
	remove         data.Method
	removeByName   data.Method
}

func (s *WindowManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewWindowManagerClassFrom(&applicationsrc.WindowManager{}), ctx.CreateBaseContext()), nil
}
func (s *WindowManagerClass) GetName() string                            { return "wails\\application\\WindowManager" }
func (s *WindowManagerClass) GetExtend() *string                         { return nil }
func (s *WindowManagerClass) GetImplements() []string                    { return nil }
func (s *WindowManagerClass) AsString() string                           { return "WindowManager{}" }
func (s *WindowManagerClass) GetSource() any                             { return s.source }
func (s *WindowManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *WindowManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *WindowManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "add":
		return s.add, true
	case "current":
		return s.current, true
	case "get":
		return s.get, true
	case "getAll":
		return s.getAll, true
	case "getByID":
		return s.getByID, true
	case "getByName":
		return s.getByName, true
	case "new":
		return s.new, true
	case "newWithOptions":
		return s.newWithOptions, true
	case "onCreate":
		return s.onCreate, true
	case "remove":
		return s.remove, true
	case "removeByName":
		return s.removeByName, true
	}
	return nil, false
}

func (s *WindowManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.add,
		s.current,
		s.get,
		s.getAll,
		s.getByID,
		s.getByName,
		s.new,
		s.newWithOptions,
		s.onCreate,
		s.remove,
		s.removeByName,
	}
}

func (s *WindowManagerClass) GetConstruct() data.Method { return s.construct }

type WindowManagerConstructMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowManagerConstructMethod) GetName() string               { return "construct" }
func (h *WindowManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowManagerConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
