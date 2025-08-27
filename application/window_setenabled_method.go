package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetEnabledMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetEnabledMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetEnabled 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.SetEnabled(arg0)
	return nil, nil
}

func (h *WindowSetEnabledMethod) GetName() string            { return "setEnabled" }
func (h *WindowSetEnabledMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetEnabledMethod) GetIsStatic() bool          { return true }
func (h *WindowSetEnabledMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetEnabledMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetEnabledMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
