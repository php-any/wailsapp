package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetEnabledMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetEnabledMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetEnabled 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.SetEnabled(arg0)
	return nil, nil
}

func (h *WebviewWindowSetEnabledMethod) GetName() string            { return "setEnabled" }
func (h *WebviewWindowSetEnabledMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetEnabledMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetEnabledMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "enabled", 0, nil, nil),
	}
}

func (h *WebviewWindowSetEnabledMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "enabled", 0, nil),
	}
}

func (h *WebviewWindowSetEnabledMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
