package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMenuClass() data.ClassStmt {
	return &MenuClass{
		source:         nil,
		construct:      &MenuConstructMethod{source: nil},
		add:            &MenuAddMethod{source: nil},
		addCheckbox:    &MenuAddCheckboxMethod{source: nil},
		addRadio:       &MenuAddRadioMethod{source: nil},
		addRole:        &MenuAddRoleMethod{source: nil},
		addSeparator:   &MenuAddSeparatorMethod{source: nil},
		addSubmenu:     &MenuAddSubmenuMethod{source: nil},
		append:         &MenuAppendMethod{source: nil},
		clear:          &MenuClearMethod{source: nil},
		clone:          &MenuCloneMethod{source: nil},
		destroy:        &MenuDestroyMethod{source: nil},
		findByLabel:    &MenuFindByLabelMethod{source: nil},
		findByRole:     &MenuFindByRoleMethod{source: nil},
		itemAt:         &MenuItemAtMethod{source: nil},
		prepend:        &MenuPrependMethod{source: nil},
		removeMenuItem: &MenuRemoveMenuItemMethod{source: nil},
		setLabel:       &MenuSetLabelMethod{source: nil},
		update:         &MenuUpdateMethod{source: nil},
	}
}

func NewMenuClassFrom(source *applicationsrc.Menu) data.ClassStmt {
	return &MenuClass{
		source:         source,
		construct:      &MenuConstructMethod{source: source},
		add:            &MenuAddMethod{source: source},
		addCheckbox:    &MenuAddCheckboxMethod{source: source},
		addRadio:       &MenuAddRadioMethod{source: source},
		addRole:        &MenuAddRoleMethod{source: source},
		addSeparator:   &MenuAddSeparatorMethod{source: source},
		addSubmenu:     &MenuAddSubmenuMethod{source: source},
		append:         &MenuAppendMethod{source: source},
		clear:          &MenuClearMethod{source: source},
		clone:          &MenuCloneMethod{source: source},
		destroy:        &MenuDestroyMethod{source: source},
		findByLabel:    &MenuFindByLabelMethod{source: source},
		findByRole:     &MenuFindByRoleMethod{source: source},
		itemAt:         &MenuItemAtMethod{source: source},
		prepend:        &MenuPrependMethod{source: source},
		removeMenuItem: &MenuRemoveMenuItemMethod{source: source},
		setLabel:       &MenuSetLabelMethod{source: source},
		update:         &MenuUpdateMethod{source: source},
	}
}

type MenuClass struct {
	node.Node
	source         *applicationsrc.Menu
	construct      data.Method
	add            data.Method
	addCheckbox    data.Method
	addRadio       data.Method
	addRole        data.Method
	addSeparator   data.Method
	addSubmenu     data.Method
	append         data.Method
	clear          data.Method
	clone          data.Method
	destroy        data.Method
	findByLabel    data.Method
	findByRole     data.Method
	itemAt         data.Method
	prepend        data.Method
	removeMenuItem data.Method
	setLabel       data.Method
	update         data.Method
}

func (s *MenuClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMenuClassFrom(&applicationsrc.Menu{}), ctx.CreateBaseContext()), nil
}
func (s *MenuClass) GetName() string                            { return "wails\\application\\Menu" }
func (s *MenuClass) GetExtend() *string                         { return nil }
func (s *MenuClass) GetImplements() []string                    { return nil }
func (s *MenuClass) AsString() string                           { return "Menu{}" }
func (s *MenuClass) GetSource() any                             { return s.source }
func (s *MenuClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *MenuClass) GetProperties() map[string]data.Property    { return nil }
func (s *MenuClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "add":
		return s.add, true
	case "addCheckbox":
		return s.addCheckbox, true
	case "addRadio":
		return s.addRadio, true
	case "addRole":
		return s.addRole, true
	case "addSeparator":
		return s.addSeparator, true
	case "addSubmenu":
		return s.addSubmenu, true
	case "append":
		return s.append, true
	case "clear":
		return s.clear, true
	case "clone":
		return s.clone, true
	case "destroy":
		return s.destroy, true
	case "findByLabel":
		return s.findByLabel, true
	case "findByRole":
		return s.findByRole, true
	case "itemAt":
		return s.itemAt, true
	case "prepend":
		return s.prepend, true
	case "removeMenuItem":
		return s.removeMenuItem, true
	case "setLabel":
		return s.setLabel, true
	case "update":
		return s.update, true
	}
	return nil, false
}

func (s *MenuClass) GetMethods() []data.Method {
	return []data.Method{
		s.add,
		s.addCheckbox,
		s.addRadio,
		s.addRole,
		s.addSeparator,
		s.addSubmenu,
		s.append,
		s.clear,
		s.clone,
		s.destroy,
		s.findByLabel,
		s.findByRole,
		s.itemAt,
		s.prepend,
		s.removeMenuItem,
		s.setLabel,
		s.update,
	}
}

func (s *MenuClass) GetConstruct() data.Method { return s.construct }

type MenuConstructMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MenuConstructMethod) GetName() string               { return "construct" }
func (h *MenuConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MenuConstructMethod) GetIsStatic() bool             { return false }
func (h *MenuConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MenuConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MenuConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
