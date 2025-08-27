package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetContentProtectionMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetContentProtectionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetContentProtection 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetContentProtection(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetContentProtectionMethod) GetName() string { return "setContentProtection" }
func (h *WebviewWindowSetContentProtectionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowSetContentProtectionMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowSetContentProtectionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "b", 0, nil, nil),
	}
}

func (h *WebviewWindowSetContentProtectionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "b", 0, nil),
	}
}

func (h *WebviewWindowSetContentProtectionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
