package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetURLMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetURLMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetURL 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetURL(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetURLMethod) GetName() string            { return "setURL" }
func (h *WebviewWindowSetURLMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetURLMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetURLMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "s", 0, nil, nil),
	}
}

func (h *WebviewWindowSetURLMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "s", 0, nil),
	}
}

func (h *WebviewWindowSetURLMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
