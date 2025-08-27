package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemIsSeparatorMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemIsSeparatorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsSeparator()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemIsSeparatorMethod) GetName() string            { return "isSeparator" }
func (h *MenuItemIsSeparatorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemIsSeparatorMethod) GetIsStatic() bool          { return true }
func (h *MenuItemIsSeparatorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemIsSeparatorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemIsSeparatorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
