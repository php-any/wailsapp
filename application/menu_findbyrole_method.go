package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuFindByRoleMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuFindByRoleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.FindByRole 缺少参数, index: 0"))
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

func (h *MenuFindByRoleMethod) GetName() string            { return "findByRole" }
func (h *MenuFindByRoleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuFindByRoleMethod) GetIsStatic() bool          { return true }
func (h *MenuFindByRoleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "role", 0, nil, nil),
	}
}

func (h *MenuFindByRoleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "role", 0, nil),
	}
}

func (h *MenuFindByRoleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
