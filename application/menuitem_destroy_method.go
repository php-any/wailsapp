package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemDestroyMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemDestroyMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Destroy()
	return nil, nil
}

func (h *MenuItemDestroyMethod) GetName() string            { return "destroy" }
func (h *MenuItemDestroyMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemDestroyMethod) GetIsStatic() bool          { return true }
func (h *MenuItemDestroyMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemDestroyMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemDestroyMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
