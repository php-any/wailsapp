package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetTooltipMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetTooltipMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetTooltip 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetTooltip(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetTooltipMethod) GetName() string            { return "setTooltip" }
func (h *MenuItemSetTooltipMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetTooltipMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetTooltipMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "s", 0, nil, nil),
	}
}

func (h *MenuItemSetTooltipMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "s", 0, nil),
	}
}

func (h *MenuItemSetTooltipMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
