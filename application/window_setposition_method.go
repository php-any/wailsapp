package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetPositionMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetPositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetPosition 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetPosition 缺少参数, index: 1"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg1, err := a1.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.SetPosition(arg0, arg1)
	return nil, nil
}

func (h *WindowSetPositionMethod) GetName() string            { return "setPosition" }
func (h *WindowSetPositionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetPositionMethod) GetIsStatic() bool          { return true }
func (h *WindowSetPositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
	}
}

func (h *WindowSetPositionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
	}
}

func (h *WindowSetPositionMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
