package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerRemoveByNameMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerRemoveByNameMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.RemoveByName 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.RemoveByName(arg0)
	return data.NewBoolValue(ret0), nil
}

func (h *WindowManagerRemoveByNameMethod) GetName() string            { return "removeByName" }
func (h *WindowManagerRemoveByNameMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerRemoveByNameMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerRemoveByNameMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
	}
}

func (h *WindowManagerRemoveByNameMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
	}
}

func (h *WindowManagerRemoveByNameMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
