package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemIsSubmenuMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemIsSubmenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsSubmenu()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemIsSubmenuMethod) GetName() string            { return "isSubmenu" }
func (h *MenuItemIsSubmenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemIsSubmenuMethod) GetIsStatic() bool          { return true }
func (h *MenuItemIsSubmenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemIsSubmenuMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemIsSubmenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
