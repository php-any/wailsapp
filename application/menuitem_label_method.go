package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemLabelMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Label()
	return data.NewStringValue(ret0), nil
}

func (h *MenuItemLabelMethod) GetName() string            { return "label" }
func (h *MenuItemLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemLabelMethod) GetIsStatic() bool          { return true }
func (h *MenuItemLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *MenuItemLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *MenuItemLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
