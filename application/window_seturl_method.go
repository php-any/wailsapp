package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetURLMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetURLMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetURL 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetURL(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowSetURLMethod) GetName() string            { return "setURL" }
func (h *WindowSetURLMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetURLMethod) GetIsStatic() bool          { return true }
func (h *WindowSetURLMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetURLMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetURLMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
