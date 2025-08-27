package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuAddCheckboxMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuAddCheckboxMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.AddCheckbox 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.AddCheckbox 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1, err := a1.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.AddCheckbox(arg0, arg1)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *ContextMenuAddCheckboxMethod) GetName() string            { return "addCheckbox" }
func (h *ContextMenuAddCheckboxMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuAddCheckboxMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuAddCheckboxMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
	}
}

func (h *ContextMenuAddCheckboxMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
	}
}

func (h *ContextMenuAddCheckboxMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
