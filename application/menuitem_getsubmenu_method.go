package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemGetSubmenuMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemGetSubmenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetSubmenu()
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *MenuItemGetSubmenuMethod) GetName() string            { return "getSubmenu" }
func (h *MenuItemGetSubmenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemGetSubmenuMethod) GetIsStatic() bool          { return true }
func (h *MenuItemGetSubmenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemGetSubmenuMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemGetSubmenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
