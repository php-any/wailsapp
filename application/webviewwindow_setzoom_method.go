package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetZoomMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetZoomMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetZoom 缺少参数, index: 0"))
	}

	arg0F, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := float64(arg0F)

	ret0 := h.source.SetZoom(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetZoomMethod) GetName() string            { return "setZoom" }
func (h *WebviewWindowSetZoomMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetZoomMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetZoomMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "magnification", 0, nil, nil),
	}
}

func (h *WebviewWindowSetZoomMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "magnification", 0, nil),
	}
}

func (h *WebviewWindowSetZoomMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
