package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerGetMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerGetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.Get 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0, ret1 := h.source.Get(arg0)
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WindowManagerGetMethod) GetName() string            { return "get" }
func (h *WindowManagerGetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerGetMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerGetMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
	}
}

func (h *WindowManagerGetMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
	}
}

func (h *WindowManagerGetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
