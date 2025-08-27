package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemOnClickMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemOnClickMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.OnClick 缺少参数, index: 0"))
	}

	var arg0 func(*applicationsrc.Context)
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func(p0 *applicationsrc.Context) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.OnClick 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnClick(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemOnClickMethod) GetName() string            { return "onClick" }
func (h *MenuItemOnClickMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemOnClickMethod) GetIsStatic() bool          { return true }
func (h *MenuItemOnClickMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "f", 0, nil, nil),
	}
}

func (h *MenuItemOnClickMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "f", 0, nil),
	}
}

func (h *MenuItemOnClickMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
