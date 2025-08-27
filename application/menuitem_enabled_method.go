package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemEnabledMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemEnabledMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Enabled()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemEnabledMethod) GetName() string            { return "enabled" }
func (h *MenuItemEnabledMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemEnabledMethod) GetIsStatic() bool          { return true }
func (h *MenuItemEnabledMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemEnabledMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemEnabledMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
