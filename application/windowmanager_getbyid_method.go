package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerGetByIDMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerGetByIDMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.GetByID 缺少参数, index: 0"))
	}

	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := uint(arg0Int)

	ret0, ret1 := h.source.GetByID(arg0)
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WindowManagerGetByIDMethod) GetName() string            { return "getByID" }
func (h *WindowManagerGetByIDMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerGetByIDMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerGetByIDMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "id", 0, nil, nil),
	}
}

func (h *WindowManagerGetByIDMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "id", 0, nil),
	}
}

func (h *WindowManagerGetByIDMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
