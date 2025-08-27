package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemAtMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuItemAtMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.ItemAt 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.ItemAt(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemAtMethod) GetName() string            { return "itemAt" }
func (h *MenuItemAtMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemAtMethod) GetIsStatic() bool          { return true }
func (h *MenuItemAtMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "index", 0, nil, nil),
	}
}

func (h *MenuItemAtMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "index", 0, nil),
	}
}

func (h *MenuItemAtMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
