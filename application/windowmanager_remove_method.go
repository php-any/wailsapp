package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerRemoveMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerRemoveMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.Remove 缺少参数, index: 0"))
	}

	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := uint(arg0Int)

	h.source.Remove(arg0)
	return nil, nil
}

func (h *WindowManagerRemoveMethod) GetName() string            { return "remove" }
func (h *WindowManagerRemoveMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerRemoveMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerRemoveMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "windowID", 0, nil, nil),
	}
}

func (h *WindowManagerRemoveMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "windowID", 0, nil),
	}
}

func (h *WindowManagerRemoveMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
