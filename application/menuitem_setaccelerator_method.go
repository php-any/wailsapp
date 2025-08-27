package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetAcceleratorMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetAcceleratorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetAccelerator 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetAccelerator(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetAcceleratorMethod) GetName() string            { return "setAccelerator" }
func (h *MenuItemSetAcceleratorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetAcceleratorMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetAcceleratorMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "shortcut", 0, nil, nil),
	}
}

func (h *MenuItemSetAcceleratorMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "shortcut", 0, nil),
	}
}

func (h *MenuItemSetAcceleratorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
