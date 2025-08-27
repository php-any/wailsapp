package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuManagerRemoveMethod struct {
	source *applicationsrc.ContextMenuManager
}

func (h *ContextMenuManagerRemoveMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Remove 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.Remove(arg0)
	return nil, nil
}

func (h *ContextMenuManagerRemoveMethod) GetName() string            { return "remove" }
func (h *ContextMenuManagerRemoveMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuManagerRemoveMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuManagerRemoveMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
	}
}

func (h *ContextMenuManagerRemoveMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
	}
}

func (h *ContextMenuManagerRemoveMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
