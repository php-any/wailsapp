package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ClipboardManagerSetTextMethod struct {
	source *applicationsrc.ClipboardManager
}

func (h *ClipboardManagerSetTextMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ClipboardManager.SetText 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetText(arg0)
	return data.NewBoolValue(ret0), nil
}

func (h *ClipboardManagerSetTextMethod) GetName() string            { return "setText" }
func (h *ClipboardManagerSetTextMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ClipboardManagerSetTextMethod) GetIsStatic() bool          { return true }
func (h *ClipboardManagerSetTextMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "text", 0, nil, nil),
	}
}

func (h *ClipboardManagerSetTextMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "text", 0, nil),
	}
}

func (h *ClipboardManagerSetTextMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
