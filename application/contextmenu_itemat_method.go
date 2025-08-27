package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuItemAtMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuItemAtMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.ItemAt 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.ItemAt(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *ContextMenuItemAtMethod) GetName() string            { return "itemAt" }
func (h *ContextMenuItemAtMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuItemAtMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuItemAtMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuItemAtMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuItemAtMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
