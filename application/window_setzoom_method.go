package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetZoomMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetZoomMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetZoom 缺少参数, index: 0"))
	}

	arg0F, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := float64(arg0F)

	ret0 := h.source.SetZoom(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowSetZoomMethod) GetName() string            { return "setZoom" }
func (h *WindowSetZoomMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetZoomMethod) GetIsStatic() bool          { return true }
func (h *WindowSetZoomMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetZoomMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetZoomMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
