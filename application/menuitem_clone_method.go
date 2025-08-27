package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemCloneMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemCloneMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Clone()
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemCloneMethod) GetName() string            { return "clone" }
func (h *MenuItemCloneMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemCloneMethod) GetIsStatic() bool          { return true }
func (h *MenuItemCloneMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemCloneMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemCloneMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
