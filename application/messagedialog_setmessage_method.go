package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogSetMessageMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogSetMessageMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetMessage 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetMessage(arg0)
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *MessageDialogSetMessageMethod) GetName() string            { return "setMessage" }
func (h *MessageDialogSetMessageMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogSetMessageMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogSetMessageMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "message", 0, nil, nil),
	}
}

func (h *MessageDialogSetMessageMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "message", 0, nil),
	}
}

func (h *MessageDialogSetMessageMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
