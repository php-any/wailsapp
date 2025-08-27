package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewKeyBindingManagerClass() data.ClassStmt {
	return &KeyBindingManagerClass{
		source:    nil,
		construct: &KeyBindingManagerConstructMethod{source: nil},
		add:       &KeyBindingManagerAddMethod{source: nil},
		getAll:    &KeyBindingManagerGetAllMethod{source: nil},
		process:   &KeyBindingManagerProcessMethod{source: nil},
		remove:    &KeyBindingManagerRemoveMethod{source: nil},
	}
}

func NewKeyBindingManagerClassFrom(source *applicationsrc.KeyBindingManager) data.ClassStmt {
	return &KeyBindingManagerClass{
		source:    source,
		construct: &KeyBindingManagerConstructMethod{source: source},
		add:       &KeyBindingManagerAddMethod{source: source},
		getAll:    &KeyBindingManagerGetAllMethod{source: source},
		process:   &KeyBindingManagerProcessMethod{source: source},
		remove:    &KeyBindingManagerRemoveMethod{source: source},
	}
}

type KeyBindingManagerClass struct {
	node.Node
	source    *applicationsrc.KeyBindingManager
	construct data.Method
	add       data.Method
	getAll    data.Method
	process   data.Method
	remove    data.Method
}

func (s *KeyBindingManagerClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewKeyBindingManagerClassFrom(&applicationsrc.KeyBindingManager{}), ctx.CreateBaseContext()), nil
}
func (s *KeyBindingManagerClass) GetName() string                            { return "wails\\application\\KeyBindingManager" }
func (s *KeyBindingManagerClass) GetExtend() *string                         { return nil }
func (s *KeyBindingManagerClass) GetImplements() []string                    { return nil }
func (s *KeyBindingManagerClass) AsString() string                           { return "KeyBindingManager{}" }
func (s *KeyBindingManagerClass) GetSource() any                             { return s.source }
func (s *KeyBindingManagerClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *KeyBindingManagerClass) GetProperties() map[string]data.Property    { return nil }
func (s *KeyBindingManagerClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "add":
		return s.add, true
	case "getAll":
		return s.getAll, true
	case "process":
		return s.process, true
	case "remove":
		return s.remove, true
	}
	return nil, false
}

func (s *KeyBindingManagerClass) GetMethods() []data.Method {
	return []data.Method{
		s.add,
		s.getAll,
		s.process,
		s.remove,
	}
}

func (s *KeyBindingManagerClass) GetConstruct() data.Method { return s.construct }

type KeyBindingManagerConstructMethod struct {
	source *applicationsrc.KeyBindingManager
}

func (h *KeyBindingManagerConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *KeyBindingManagerConstructMethod) GetName() string               { return "construct" }
func (h *KeyBindingManagerConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *KeyBindingManagerConstructMethod) GetIsStatic() bool             { return false }
func (h *KeyBindingManagerConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *KeyBindingManagerConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *KeyBindingManagerConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
