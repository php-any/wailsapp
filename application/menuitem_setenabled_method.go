package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetEnabledMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetEnabledMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetEnabled 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetEnabled(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetEnabledMethod) GetName() string            { return "setEnabled" }
func (h *MenuItemSetEnabledMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetEnabledMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetEnabledMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "enabled", 0, nil, nil),
	}
}

func (h *MenuItemSetEnabledMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "enabled", 0, nil),
	}
}

func (h *MenuItemSetEnabledMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
