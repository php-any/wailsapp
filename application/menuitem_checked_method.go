package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemCheckedMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemCheckedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Checked()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemCheckedMethod) GetName() string            { return "checked" }
func (h *MenuItemCheckedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemCheckedMethod) GetIsStatic() bool          { return true }
func (h *MenuItemCheckedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemCheckedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemCheckedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
