package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetTitleMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetTitleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetTitle 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetTitle(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowSetTitleMethod) GetName() string            { return "setTitle" }
func (h *WindowSetTitleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetTitleMethod) GetIsStatic() bool          { return true }
func (h *WindowSetTitleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetTitleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetTitleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
