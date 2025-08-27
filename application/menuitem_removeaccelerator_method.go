package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemRemoveAcceleratorMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemRemoveAcceleratorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.RemoveAccelerator()
	return nil, nil
}

func (h *MenuItemRemoveAcceleratorMethod) GetName() string            { return "removeAccelerator" }
func (h *MenuItemRemoveAcceleratorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemRemoveAcceleratorMethod) GetIsStatic() bool          { return true }
func (h *MenuItemRemoveAcceleratorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemRemoveAcceleratorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemRemoveAcceleratorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
