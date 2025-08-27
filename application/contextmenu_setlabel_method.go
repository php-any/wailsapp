package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuSetLabelMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuSetLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.SetLabel 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.SetLabel(arg0)
	return nil, nil
}

func (h *ContextMenuSetLabelMethod) GetName() string            { return "setLabel" }
func (h *ContextMenuSetLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuSetLabelMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuSetLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuSetLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuSetLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
