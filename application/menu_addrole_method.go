package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuAddRoleMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuAddRoleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.AddRole 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Role
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.Role(arg0Int)

	ret0 := h.source.AddRole(arg0)
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *MenuAddRoleMethod) GetName() string            { return "addRole" }
func (h *MenuAddRoleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuAddRoleMethod) GetIsStatic() bool          { return true }
func (h *MenuAddRoleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "role", 0, nil, nil),
	}
}

func (h *MenuAddRoleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "role", 0, nil),
	}
}

func (h *MenuAddRoleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
