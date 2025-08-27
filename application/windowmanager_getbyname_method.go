package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerGetByNameMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerGetByNameMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.GetByName 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0, ret1 := h.source.GetByName(arg0)
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WindowManagerGetByNameMethod) GetName() string            { return "getByName" }
func (h *WindowManagerGetByNameMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerGetByNameMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerGetByNameMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
	}
}

func (h *WindowManagerGetByNameMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
	}
}

func (h *WindowManagerGetByNameMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
