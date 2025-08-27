package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetHiddenMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetHiddenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetHidden 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetHidden(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetHiddenMethod) GetName() string            { return "setHidden" }
func (h *MenuItemSetHiddenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetHiddenMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetHiddenMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "hidden", 0, nil, nil),
	}
}

func (h *MenuItemSetHiddenMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "hidden", 0, nil),
	}
}

func (h *MenuItemSetHiddenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
