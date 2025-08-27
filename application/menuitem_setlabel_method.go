package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetLabelMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetLabel 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetLabel(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetLabelMethod) GetName() string            { return "setLabel" }
func (h *MenuItemSetLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetLabelMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "s", 0, nil, nil),
	}
}

func (h *MenuItemSetLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "s", 0, nil),
	}
}

func (h *MenuItemSetLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
