package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewMenuItemClass() data.ClassStmt {
	return &MenuItemClass{
		source:            nil,
		construct:         &MenuItemConstructMethod{source: nil},
		checked:           &MenuItemCheckedMethod{source: nil},
		clone:             &MenuItemCloneMethod{source: nil},
		destroy:           &MenuItemDestroyMethod{source: nil},
		enabled:           &MenuItemEnabledMethod{source: nil},
		getAccelerator:    &MenuItemGetAcceleratorMethod{source: nil},
		getSubmenu:        &MenuItemGetSubmenuMethod{source: nil},
		hidden:            &MenuItemHiddenMethod{source: nil},
		isCheckbox:        &MenuItemIsCheckboxMethod{source: nil},
		isRadio:           &MenuItemIsRadioMethod{source: nil},
		isSeparator:       &MenuItemIsSeparatorMethod{source: nil},
		isSubmenu:         &MenuItemIsSubmenuMethod{source: nil},
		label:             &MenuItemLabelMethod{source: nil},
		onClick:           &MenuItemOnClickMethod{source: nil},
		removeAccelerator: &MenuItemRemoveAcceleratorMethod{source: nil},
		setAccelerator:    &MenuItemSetAcceleratorMethod{source: nil},
		setBitmap:         &MenuItemSetBitmapMethod{source: nil},
		setChecked:        &MenuItemSetCheckedMethod{source: nil},
		setEnabled:        &MenuItemSetEnabledMethod{source: nil},
		setHidden:         &MenuItemSetHiddenMethod{source: nil},
		setLabel:          &MenuItemSetLabelMethod{source: nil},
		setRole:           &MenuItemSetRoleMethod{source: nil},
		setTooltip:        &MenuItemSetTooltipMethod{source: nil},
		tooltip:           &MenuItemTooltipMethod{source: nil},
	}
}

func NewMenuItemClassFrom(source *applicationsrc.MenuItem) data.ClassStmt {
	return &MenuItemClass{
		source:            source,
		construct:         &MenuItemConstructMethod{source: source},
		checked:           &MenuItemCheckedMethod{source: source},
		clone:             &MenuItemCloneMethod{source: source},
		destroy:           &MenuItemDestroyMethod{source: source},
		enabled:           &MenuItemEnabledMethod{source: source},
		getAccelerator:    &MenuItemGetAcceleratorMethod{source: source},
		getSubmenu:        &MenuItemGetSubmenuMethod{source: source},
		hidden:            &MenuItemHiddenMethod{source: source},
		isCheckbox:        &MenuItemIsCheckboxMethod{source: source},
		isRadio:           &MenuItemIsRadioMethod{source: source},
		isSeparator:       &MenuItemIsSeparatorMethod{source: source},
		isSubmenu:         &MenuItemIsSubmenuMethod{source: source},
		label:             &MenuItemLabelMethod{source: source},
		onClick:           &MenuItemOnClickMethod{source: source},
		removeAccelerator: &MenuItemRemoveAcceleratorMethod{source: source},
		setAccelerator:    &MenuItemSetAcceleratorMethod{source: source},
		setBitmap:         &MenuItemSetBitmapMethod{source: source},
		setChecked:        &MenuItemSetCheckedMethod{source: source},
		setEnabled:        &MenuItemSetEnabledMethod{source: source},
		setHidden:         &MenuItemSetHiddenMethod{source: source},
		setLabel:          &MenuItemSetLabelMethod{source: source},
		setRole:           &MenuItemSetRoleMethod{source: source},
		setTooltip:        &MenuItemSetTooltipMethod{source: source},
		tooltip:           &MenuItemTooltipMethod{source: source},
	}
}

type MenuItemClass struct {
	node.Node
	source            *applicationsrc.MenuItem
	construct         data.Method
	checked           data.Method
	clone             data.Method
	destroy           data.Method
	enabled           data.Method
	getAccelerator    data.Method
	getSubmenu        data.Method
	hidden            data.Method
	isCheckbox        data.Method
	isRadio           data.Method
	isSeparator       data.Method
	isSubmenu         data.Method
	label             data.Method
	onClick           data.Method
	removeAccelerator data.Method
	setAccelerator    data.Method
	setBitmap         data.Method
	setChecked        data.Method
	setEnabled        data.Method
	setHidden         data.Method
	setLabel          data.Method
	setRole           data.Method
	setTooltip        data.Method
	tooltip           data.Method
}

func (s *MenuItemClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewMenuItemClassFrom(&applicationsrc.MenuItem{}), ctx.CreateBaseContext()), nil
}
func (s *MenuItemClass) GetName() string                            { return "wails\\application\\MenuItem" }
func (s *MenuItemClass) GetExtend() *string                         { return nil }
func (s *MenuItemClass) GetImplements() []string                    { return nil }
func (s *MenuItemClass) AsString() string                           { return "MenuItem{}" }
func (s *MenuItemClass) GetSource() any                             { return s.source }
func (s *MenuItemClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *MenuItemClass) GetProperties() map[string]data.Property    { return nil }
func (s *MenuItemClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "checked":
		return s.checked, true
	case "clone":
		return s.clone, true
	case "destroy":
		return s.destroy, true
	case "enabled":
		return s.enabled, true
	case "getAccelerator":
		return s.getAccelerator, true
	case "getSubmenu":
		return s.getSubmenu, true
	case "hidden":
		return s.hidden, true
	case "isCheckbox":
		return s.isCheckbox, true
	case "isRadio":
		return s.isRadio, true
	case "isSeparator":
		return s.isSeparator, true
	case "isSubmenu":
		return s.isSubmenu, true
	case "label":
		return s.label, true
	case "onClick":
		return s.onClick, true
	case "removeAccelerator":
		return s.removeAccelerator, true
	case "setAccelerator":
		return s.setAccelerator, true
	case "setBitmap":
		return s.setBitmap, true
	case "setChecked":
		return s.setChecked, true
	case "setEnabled":
		return s.setEnabled, true
	case "setHidden":
		return s.setHidden, true
	case "setLabel":
		return s.setLabel, true
	case "setRole":
		return s.setRole, true
	case "setTooltip":
		return s.setTooltip, true
	case "tooltip":
		return s.tooltip, true
	}
	return nil, false
}

func (s *MenuItemClass) GetMethods() []data.Method {
	return []data.Method{
		s.checked,
		s.clone,
		s.destroy,
		s.enabled,
		s.getAccelerator,
		s.getSubmenu,
		s.hidden,
		s.isCheckbox,
		s.isRadio,
		s.isSeparator,
		s.isSubmenu,
		s.label,
		s.onClick,
		s.removeAccelerator,
		s.setAccelerator,
		s.setBitmap,
		s.setChecked,
		s.setEnabled,
		s.setHidden,
		s.setLabel,
		s.setRole,
		s.setTooltip,
		s.tooltip,
	}
}

func (s *MenuItemClass) GetConstruct() data.Method { return s.construct }

type MenuItemConstructMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *MenuItemConstructMethod) GetName() string               { return "construct" }
func (h *MenuItemConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *MenuItemConstructMethod) GetIsStatic() bool             { return false }
func (h *MenuItemConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *MenuItemConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *MenuItemConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
