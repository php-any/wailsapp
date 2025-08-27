package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetFramelessMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetFramelessMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetFrameless 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetFrameless(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetFramelessMethod) GetName() string            { return "setFrameless" }
func (h *WebviewWindowSetFramelessMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetFramelessMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetFramelessMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "frameless", 0, nil, nil),
	}
}

func (h *WebviewWindowSetFramelessMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "frameless", 0, nil),
	}
}

func (h *WebviewWindowSetFramelessMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
