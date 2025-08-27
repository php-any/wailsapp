package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemIsRadioMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemIsRadioMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsRadio()
	return data.NewBoolValue(ret0), nil
}

func (h *MenuItemIsRadioMethod) GetName() string            { return "isRadio" }
func (h *MenuItemIsRadioMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemIsRadioMethod) GetIsStatic() bool          { return true }
func (h *MenuItemIsRadioMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemIsRadioMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemIsRadioMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
