package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetRoleMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetRoleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetRole 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Role
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.Role(arg0Int)

	ret0 := h.source.SetRole(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetRoleMethod) GetName() string            { return "setRole" }
func (h *MenuItemSetRoleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetRoleMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetRoleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "role", 0, nil, nil),
	}
}

func (h *MenuItemSetRoleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "role", 0, nil),
	}
}

func (h *MenuItemSetRoleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
