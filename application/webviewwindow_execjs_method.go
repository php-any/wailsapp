package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowExecJSMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowExecJSMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.ExecJS 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.ExecJS(arg0)
	return nil, nil
}

func (h *WebviewWindowExecJSMethod) GetName() string            { return "execJS" }
func (h *WebviewWindowExecJSMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowExecJSMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowExecJSMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "js", 0, nil, nil),
	}
}

func (h *WebviewWindowExecJSMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "js", 0, nil),
	}
}

func (h *WebviewWindowExecJSMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
