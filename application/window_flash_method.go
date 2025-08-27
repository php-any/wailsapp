package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowFlashMethod struct {
	source applicationsrc.Window
}

func (h *WindowFlashMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.Flash 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.Flash(arg0)
	return nil, nil
}

func (h *WindowFlashMethod) GetName() string            { return "flash" }
func (h *WindowFlashMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowFlashMethod) GetIsStatic() bool          { return true }
func (h *WindowFlashMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowFlashMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowFlashMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
