package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogAddButtonMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogAddButtonMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.AddButton 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.AddButton(arg0)
	return data.NewClassValue(NewButtonClassFrom(ret0), ctx), nil
}

func (h *MessageDialogAddButtonMethod) GetName() string            { return "addButton" }
func (h *MessageDialogAddButtonMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogAddButtonMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogAddButtonMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "s", 0, nil, nil),
	}
}

func (h *MessageDialogAddButtonMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "s", 0, nil),
	}
}

func (h *MessageDialogAddButtonMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
