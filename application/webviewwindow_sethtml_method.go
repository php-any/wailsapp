package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetHTMLMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetHTMLMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetHTML 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetHTML(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetHTMLMethod) GetName() string            { return "setHTML" }
func (h *WebviewWindowSetHTMLMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetHTMLMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetHTMLMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "html", 0, nil, nil),
	}
}

func (h *WebviewWindowSetHTMLMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "html", 0, nil),
	}
}

func (h *WebviewWindowSetHTMLMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
