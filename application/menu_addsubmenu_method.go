package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuAddSubmenuMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuAddSubmenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.AddSubmenu 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.AddSubmenu(arg0)
	return data.NewClassValue(NewMenuClassFrom(ret0), ctx), nil
}

func (h *MenuAddSubmenuMethod) GetName() string            { return "addSubmenu" }
func (h *MenuAddSubmenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuAddSubmenuMethod) GetIsStatic() bool          { return true }
func (h *MenuAddSubmenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "s", 0, nil, nil),
	}
}

func (h *MenuAddSubmenuMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "s", 0, nil),
	}
}

func (h *MenuAddSubmenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
