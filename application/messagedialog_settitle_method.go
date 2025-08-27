package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogSetTitleMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogSetTitleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetTitle 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetTitle(arg0)
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *MessageDialogSetTitleMethod) GetName() string            { return "setTitle" }
func (h *MessageDialogSetTitleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogSetTitleMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogSetTitleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "title", 0, nil, nil),
	}
}

func (h *MessageDialogSetTitleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "title", 0, nil),
	}
}

func (h *MessageDialogSetTitleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
