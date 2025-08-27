package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemHiddenMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemHiddenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Hidden()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemHiddenMethod) GetName() string            { return "hidden" }
func (h *MenuItemHiddenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemHiddenMethod) GetIsStatic() bool          { return true }
func (h *MenuItemHiddenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemHiddenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemHiddenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
