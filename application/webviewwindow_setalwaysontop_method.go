package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetAlwaysOnTopMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetAlwaysOnTopMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetAlwaysOnTop 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetAlwaysOnTop(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetAlwaysOnTopMethod) GetName() string            { return "setAlwaysOnTop" }
func (h *WebviewWindowSetAlwaysOnTopMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetAlwaysOnTopMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetAlwaysOnTopMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "b", 0, nil, nil),
	}
}

func (h *WebviewWindowSetAlwaysOnTopMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "b", 0, nil),
	}
}

func (h *WebviewWindowSetAlwaysOnTopMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
