package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowFlashMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowFlashMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.Flash 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.Flash(arg0)
	return nil, nil
}

func (h *WebviewWindowFlashMethod) GetName() string            { return "flash" }
func (h *WebviewWindowFlashMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowFlashMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowFlashMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "enabled", 0, nil, nil),
	}
}

func (h *WebviewWindowFlashMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "enabled", 0, nil),
	}
}

func (h *WebviewWindowFlashMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
