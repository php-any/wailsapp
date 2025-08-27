package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	"github.com/php-any/origami/runtime"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewContextMenuClass() data.ClassStmt {
	return &ContextMenuClass{
		source:         nil,
		construct:      &ContextMenuConstructMethod{source: nil},
		add:            &ContextMenuAddMethod{source: nil},
		addCheckbox:    &ContextMenuAddCheckboxMethod{source: nil},
		addRadio:       &ContextMenuAddRadioMethod{source: nil},
		addRole:        &ContextMenuAddRoleMethod{source: nil},
		addSeparator:   &ContextMenuAddSeparatorMethod{source: nil},
		addSubmenu:     &ContextMenuAddSubmenuMethod{source: nil},
		append:         &ContextMenuAppendMethod{source: nil},
		clear:          &ContextMenuClearMethod{source: nil},
		clone:          &ContextMenuCloneMethod{source: nil},
		destroy:        &ContextMenuDestroyMethod{source: nil},
		findByLabel:    &ContextMenuFindByLabelMethod{source: nil},
		findByRole:     &ContextMenuFindByRoleMethod{source: nil},
		itemAt:         &ContextMenuItemAtMethod{source: nil},
		prepend:        &ContextMenuPrependMethod{source: nil},
		removeMenuItem: &ContextMenuRemoveMenuItemMethod{source: nil},
		setLabel:       &ContextMenuSetLabelMethod{source: nil},
		update:         &ContextMenuUpdateMethod{source: nil},
	}
}

func NewContextMenuClassFrom(source *applicationsrc.ContextMenu) data.ClassStmt {
	return &ContextMenuClass{
		source:         source,
		construct:      &ContextMenuConstructMethod{source: source},
		add:            &ContextMenuAddMethod{source: source},
		addCheckbox:    &ContextMenuAddCheckboxMethod{source: source},
		addRadio:       &ContextMenuAddRadioMethod{source: source},
		addRole:        &ContextMenuAddRoleMethod{source: source},
		addSeparator:   &ContextMenuAddSeparatorMethod{source: source},
		addSubmenu:     &ContextMenuAddSubmenuMethod{source: source},
		append:         &ContextMenuAppendMethod{source: source},
		clear:          &ContextMenuClearMethod{source: source},
		clone:          &ContextMenuCloneMethod{source: source},
		destroy:        &ContextMenuDestroyMethod{source: source},
		findByLabel:    &ContextMenuFindByLabelMethod{source: source},
		findByRole:     &ContextMenuFindByRoleMethod{source: source},
		itemAt:         &ContextMenuItemAtMethod{source: source},
		prepend:        &ContextMenuPrependMethod{source: source},
		removeMenuItem: &ContextMenuRemoveMenuItemMethod{source: source},
		setLabel:       &ContextMenuSetLabelMethod{source: source},
		update:         &ContextMenuUpdateMethod{source: source},
	}
}

type ContextMenuClass struct {
	node.Node
	source         *applicationsrc.ContextMenu
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

func (s *ContextMenuClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewContextMenuClassFrom(&applicationsrc.ContextMenu{}), ctx.CreateBaseContext()), nil
}
func (s *ContextMenuClass) GetName() string         { return "wails\\application\\ContextMenu" }
func (s *ContextMenuClass) GetExtend() *string      { return nil }
func (s *ContextMenuClass) GetImplements() []string { return nil }
func (s *ContextMenuClass) AsString() string        { return "ContextMenu{}" }
func (s *ContextMenuClass) GetSource() any          { return s.source }
func (s *ContextMenuClass) GetProperty(name string) (data.Property, bool) {
	switch name {
	case "Menu":
		return node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(NewMenuClassFrom(s.source.Menu), runtime.NewContextToDo())), true
	}
	return nil, false
}

func (s *ContextMenuClass) GetProperties() map[string]data.Property {
	return map[string]data.Property{
		"Menu": node.NewProperty(nil, "Menu", "public", true, data.NewClassValue(nil, runtime.NewContextToDo())),
	}
}

func (s *ContextMenuClass) SetProperty(name string, value data.Value) {
	switch name {
	case "Menu":
		switch v := value.(type) {
		case data.GetSource:
			if src := v.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Menu); ok {
					s.source.Menu = ptr
				}
			}
		case *data.AnyValue:
			s.source.Menu = v.Value.(*applicationsrc.Menu)
		}
	}
}

func (s *ContextMenuClass) GetMethod(name string) (data.Method, bool) {
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

func (s *ContextMenuClass) GetMethods() []data.Method {
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

func (s *ContextMenuClass) GetConstruct() data.Method { return s.construct }

type ContextMenuConstructMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *ContextMenuConstructMethod) GetName() string               { return "construct" }
func (h *ContextMenuConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *ContextMenuConstructMethod) GetIsStatic() bool             { return false }
func (h *ContextMenuConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *ContextMenuConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *ContextMenuConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
