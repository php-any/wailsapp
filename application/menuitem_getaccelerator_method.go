package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemGetAcceleratorMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemGetAcceleratorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetAccelerator()
	return data.NewStringValue(ret0), nil
}

func (h *MenuItemGetAcceleratorMethod) GetName() string            { return "getAccelerator" }
func (h *MenuItemGetAcceleratorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemGetAcceleratorMethod) GetIsStatic() bool          { return true }
func (h *MenuItemGetAcceleratorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemGetAcceleratorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemGetAcceleratorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
