package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewSystemTrayClass() data.ClassStmt {
	return &SystemTrayClass{
		source:             nil,
		construct:          &SystemTrayConstructMethod{source: nil},
		attachWindow:       &SystemTrayAttachWindowMethod{source: nil},
		destroy:            &SystemTrayDestroyMethod{source: nil},
		hide:               &SystemTrayHideMethod{source: nil},
		label:              &SystemTrayLabelMethod{source: nil},
		onClick:            &SystemTrayOnClickMethod{source: nil},
		onDoubleClick:      &SystemTrayOnDoubleClickMethod{source: nil},
		onMouseEnter:       &SystemTrayOnMouseEnterMethod{source: nil},
		onMouseLeave:       &SystemTrayOnMouseLeaveMethod{source: nil},
		onRightClick:       &SystemTrayOnRightClickMethod{source: nil},
		onRightDoubleClick: &SystemTrayOnRightDoubleClickMethod{source: nil},
		openMenu:           &SystemTrayOpenMenuMethod{source: nil},
		positionWindow:     &SystemTrayPositionWindowMethod{source: nil},
		run:                &SystemTrayRunMethod{source: nil},
		setDarkModeIcon:    &SystemTraySetDarkModeIconMethod{source: nil},
		setIcon:            &SystemTraySetIconMethod{source: nil},
		setIconPosition:    &SystemTraySetIconPositionMethod{source: nil},
		setLabel:           &SystemTraySetLabelMethod{source: nil},
		setMenu:            &SystemTraySetMenuMethod{source: nil},
		setTemplateIcon:    &SystemTraySetTemplateIconMethod{source: nil},
		setTooltip:         &SystemTraySetTooltipMethod{source: nil},
		show:               &SystemTrayShowMethod{source: nil},
		windowDebounce:     &SystemTrayWindowDebounceMethod{source: nil},
		windowOffset:       &SystemTrayWindowOffsetMethod{source: nil},
	}
}

func NewSystemTrayClassFrom(source *applicationsrc.SystemTray) data.ClassStmt {
	return &SystemTrayClass{
		source:             source,
		construct:          &SystemTrayConstructMethod{source: source},
		attachWindow:       &SystemTrayAttachWindowMethod{source: source},
		destroy:            &SystemTrayDestroyMethod{source: source},
		hide:               &SystemTrayHideMethod{source: source},
		label:              &SystemTrayLabelMethod{source: source},
		onClick:            &SystemTrayOnClickMethod{source: source},
		onDoubleClick:      &SystemTrayOnDoubleClickMethod{source: source},
		onMouseEnter:       &SystemTrayOnMouseEnterMethod{source: source},
		onMouseLeave:       &SystemTrayOnMouseLeaveMethod{source: source},
		onRightClick:       &SystemTrayOnRightClickMethod{source: source},
		onRightDoubleClick: &SystemTrayOnRightDoubleClickMethod{source: source},
		openMenu:           &SystemTrayOpenMenuMethod{source: source},
		positionWindow:     &SystemTrayPositionWindowMethod{source: source},
		run:                &SystemTrayRunMethod{source: source},
		setDarkModeIcon:    &SystemTraySetDarkModeIconMethod{source: source},
		setIcon:            &SystemTraySetIconMethod{source: source},
		setIconPosition:    &SystemTraySetIconPositionMethod{source: source},
		setLabel:           &SystemTraySetLabelMethod{source: source},
		setMenu:            &SystemTraySetMenuMethod{source: source},
		setTemplateIcon:    &SystemTraySetTemplateIconMethod{source: source},
		setTooltip:         &SystemTraySetTooltipMethod{source: source},
		show:               &SystemTrayShowMethod{source: source},
		windowDebounce:     &SystemTrayWindowDebounceMethod{source: source},
		windowOffset:       &SystemTrayWindowOffsetMethod{source: source},
	}
}

type SystemTrayClass struct {
	node.Node
	source             *applicationsrc.SystemTray
	construct          data.Method
	attachWindow       data.Method
	destroy            data.Method
	hide               data.Method
	label              data.Method
	onClick            data.Method
	onDoubleClick      data.Method
	onMouseEnter       data.Method
	onMouseLeave       data.Method
	onRightClick       data.Method
	onRightDoubleClick data.Method
	openMenu           data.Method
	positionWindow     data.Method
	run                data.Method
	setDarkModeIcon    data.Method
	setIcon            data.Method
	setIconPosition    data.Method
	setLabel           data.Method
	setMenu            data.Method
	setTemplateIcon    data.Method
	setTooltip         data.Method
	show               data.Method
	windowDebounce     data.Method
	windowOffset       data.Method
}

func (s *SystemTrayClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewSystemTrayClassFrom(&applicationsrc.SystemTray{}), ctx.CreateBaseContext()), nil
}
func (s *SystemTrayClass) GetName() string                            { return "wails\\application\\SystemTray" }
func (s *SystemTrayClass) GetExtend() *string                         { return nil }
func (s *SystemTrayClass) GetImplements() []string                    { return nil }
func (s *SystemTrayClass) AsString() string                           { return "SystemTray{}" }
func (s *SystemTrayClass) GetSource() any                             { return s.source }
func (s *SystemTrayClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *SystemTrayClass) GetProperties() map[string]data.Property    { return nil }
func (s *SystemTrayClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "attachWindow":
		return s.attachWindow, true
	case "destroy":
		return s.destroy, true
	case "hide":
		return s.hide, true
	case "label":
		return s.label, true
	case "onClick":
		return s.onClick, true
	case "onDoubleClick":
		return s.onDoubleClick, true
	case "onMouseEnter":
		return s.onMouseEnter, true
	case "onMouseLeave":
		return s.onMouseLeave, true
	case "onRightClick":
		return s.onRightClick, true
	case "onRightDoubleClick":
		return s.onRightDoubleClick, true
	case "openMenu":
		return s.openMenu, true
	case "positionWindow":
		return s.positionWindow, true
	case "run":
		return s.run, true
	case "setDarkModeIcon":
		return s.setDarkModeIcon, true
	case "setIcon":
		return s.setIcon, true
	case "setIconPosition":
		return s.setIconPosition, true
	case "setLabel":
		return s.setLabel, true
	case "setMenu":
		return s.setMenu, true
	case "setTemplateIcon":
		return s.setTemplateIcon, true
	case "setTooltip":
		return s.setTooltip, true
	case "show":
		return s.show, true
	case "windowDebounce":
		return s.windowDebounce, true
	case "windowOffset":
		return s.windowOffset, true
	}
	return nil, false
}

func (s *SystemTrayClass) GetMethods() []data.Method {
	return []data.Method{
		s.attachWindow,
		s.destroy,
		s.hide,
		s.label,
		s.onClick,
		s.onDoubleClick,
		s.onMouseEnter,
		s.onMouseLeave,
		s.onRightClick,
		s.onRightDoubleClick,
		s.openMenu,
		s.positionWindow,
		s.run,
		s.setDarkModeIcon,
		s.setIcon,
		s.setIconPosition,
		s.setLabel,
		s.setMenu,
		s.setTemplateIcon,
		s.setTooltip,
		s.show,
		s.windowDebounce,
		s.windowOffset,
	}
}

func (s *SystemTrayClass) GetConstruct() data.Method { return s.construct }

type SystemTrayConstructMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *SystemTrayConstructMethod) GetName() string               { return "construct" }
func (h *SystemTrayConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *SystemTrayConstructMethod) GetIsStatic() bool             { return false }
func (h *SystemTrayConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *SystemTrayConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *SystemTrayConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
