package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuAddMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuAddMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.Add 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.Add(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *ContextMenuAddMethod) GetName() string            { return "add" }
func (h *ContextMenuAddMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuAddMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuAddMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuAddMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuAddMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
