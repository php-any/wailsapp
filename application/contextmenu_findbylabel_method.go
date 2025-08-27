package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuFindByLabelMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuFindByLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.FindByLabel 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.FindByLabel(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *ContextMenuFindByLabelMethod) GetName() string            { return "findByLabel" }
func (h *ContextMenuFindByLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuFindByLabelMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuFindByLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuFindByLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuFindByLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
