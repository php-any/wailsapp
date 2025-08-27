package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewContextMenuManagerClass() data.ClassStmt {
	return &ContextMenuManagerClass{
		source:    nil,
		construct: &ContextMenuManagerConstructMethod{source: nil},
		add:       &ContextMenuManagerAddMethod{source: nil},
		get:       &ContextMenuManagerGetMethod{source: nil},
		getAll:    &ContextMenuManagerGetAllMethod{source: nil},
		new:       &ContextMenuManagerNewMethod{source: nil},
		remove:    &ContextMenuManagerRemoveMethod{source: nil},
	}
}

func NewContextMenuManagerClassFrom(source *applicationsrc.ContextMenuManager) data.ClassStmt {
	return &ContextMenuManagerClass{
		source:    source,
		construct: &ContextMenuManagerConstructMethod{source: source},
		add:       &ContextMenuManagerAddMethod{source: source},
		get:       &ContextMenuManagerGetMethod{source: source},
		getAll:    &ContextMenuManagerGetAllMethod{source: source},
		new:       &ContextMenuManagerNewMethod{source: source},
		remove:    &ContextMenuManagerRemoveMethod{source: source},
	}
}

type ContextMenuManagerClass struct {
	node.Node
	source    *applicationsrc.ContextMenuManager
	construct data.Method
	add       data.Method
	get       data.Method
	getAll    data.Method
	new       data.Method
	remove    data.Method
}

func (s *ContextMenuManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewContextMenuManagerClassFrom(&applicationsrc.ContextMenuManager{}), ctx.CreateBaseContext()), nil
}
func (s *ContextMenuManagerClass) GetName() string                            { return "wails\\application\\ContextMenuManager" }
func (s *ContextMenuManagerClass) GetExtend() *string                         { return nil }
func (s *ContextMenuManagerClass) GetImplements() []string                    { return nil }
func (s *ContextMenuManagerClass) AsString() string                           { return "ContextMenuManager{}" }
func (s *ContextMenuManagerClass) GetSource() any                             { return s.source }
func (s *ContextMenuManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *ContextMenuManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *ContextMenuManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "add":
		return s.add, true
	case "get":
		return s.get, true
	case "getAll":
		return s.getAll, true
	case "new":
		return s.new, true
	case "remove":
		return s.remove, true
	}
	return nil, false
}

func (s *ContextMenuManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.add,
		s.get,
		s.getAll,
		s.new,
		s.remove,
	}
}

func (s *ContextMenuManagerClass) GetConstruct() data.Method { return s.construct }

type ContextMenuManagerConstructMethod struct {
	source *applicationsrc.ContextMenuManager
}

func (h *ContextMenuManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ContextMenuManagerConstructMethod) GetName() string               { return "construct" }
func (h *ContextMenuManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ContextMenuManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *ContextMenuManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ContextMenuManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ContextMenuManagerConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
