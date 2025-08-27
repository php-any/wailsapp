package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowExecJSMethod struct {
	source applicationsrc.Window
}

func (h *WindowExecJSMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.ExecJS 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.ExecJS(arg0)
	return nil, nil
}

func (h *WindowExecJSMethod) GetName() string            { return "execJS" }
func (h *WindowExecJSMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowExecJSMethod) GetIsStatic() bool          { return true }
func (h *WindowExecJSMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowExecJSMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowExecJSMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
