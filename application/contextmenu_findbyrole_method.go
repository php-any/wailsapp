package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuFindByRoleMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuFindByRoleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.FindByRole 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Role
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.Role(arg0Int)

	ret0 := h.source.FindByRole(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *ContextMenuFindByRoleMethod) GetName() string            { return "findByRole" }
func (h *ContextMenuFindByRoleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuFindByRoleMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuFindByRoleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuFindByRoleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuFindByRoleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
