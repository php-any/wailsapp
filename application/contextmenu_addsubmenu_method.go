package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuAddSubmenuMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuAddSubmenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.AddSubmenu 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.AddSubmenu(arg0)
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *ContextMenuAddSubmenuMethod) GetName() string            { return "addSubmenu" }
func (h *ContextMenuAddSubmenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuAddSubmenuMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuAddSubmenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuAddSubmenuMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuAddSubmenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
