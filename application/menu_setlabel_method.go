package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuSetLabelMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuSetLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.SetLabel 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.SetLabel(arg0)
	return nil, nil
}

func (h *MenuSetLabelMethod) GetName() string            { return "setLabel" }
func (h *MenuSetLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuSetLabelMethod) GetIsStatic() bool          { return true }
func (h *MenuSetLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "label", 0, nil, nil),
	}
}

func (h *MenuSetLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "label", 0, nil),
	}
}

func (h *MenuSetLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
