package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemTooltipMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemTooltipMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Tooltip()
	return data.NewStringValue(ret0), nil
}

func (h *MenuItemTooltipMethod) GetName() string            { return "tooltip" }
func (h *MenuItemTooltipMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemTooltipMethod) GetIsStatic() bool          { return true }
func (h *MenuItemTooltipMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemTooltipMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemTooltipMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
