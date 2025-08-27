package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetIgnoreMouseEventsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetIgnoreMouseEventsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetIgnoreMouseEvents 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetIgnoreMouseEvents(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetIgnoreMouseEventsMethod) GetName() string { return "setIgnoreMouseEvents" }
func (h *WebviewWindowSetIgnoreMouseEventsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowSetIgnoreMouseEventsMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowSetIgnoreMouseEventsMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "ignore", 0, nil, nil),
	}
}

func (h *WebviewWindowSetIgnoreMouseEventsMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "ignore", 0, nil),
	}
}

func (h *WebviewWindowSetIgnoreMouseEventsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
