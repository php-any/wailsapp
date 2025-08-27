package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuManagerGetMethod struct {
	source *applicationsrc.ContextMenuManager
}

func (h *ContextMenuManagerGetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Get 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0, ret1 := h.source.Get(arg0)
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *ContextMenuManagerGetMethod) GetName() string            { return "get" }
func (h *ContextMenuManagerGetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuManagerGetMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuManagerGetMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
	}
}

func (h *ContextMenuManagerGetMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
	}
}

func (h *ContextMenuManagerGetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
