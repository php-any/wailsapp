package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuAddMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuAddMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.Add 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.Add(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuAddMethod) GetName() string            { return "add" }
func (h *MenuAddMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuAddMethod) GetIsStatic() bool          { return true }
func (h *MenuAddMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "label", 0, nil, nil),
	}
}

func (h *MenuAddMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "label", 0, nil),
	}
}

func (h *MenuAddMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
