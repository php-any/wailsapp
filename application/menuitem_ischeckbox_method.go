package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemIsCheckboxMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemIsCheckboxMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsCheckbox()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemIsCheckboxMethod) GetName() string            { return "isCheckbox" }
func (h *MenuItemIsCheckboxMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemIsCheckboxMethod) GetIsStatic() bool          { return true }
func (h *MenuItemIsCheckboxMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemIsCheckboxMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemIsCheckboxMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
