package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTraySetLabelMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTraySetLabelMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetLabel 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.SetLabel(arg0)
	return nil, nil
}

func (h *SystemTraySetLabelMethod) GetName() string            { return "setLabel" }
func (h *SystemTraySetLabelMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTraySetLabelMethod) GetIsStatic() bool          { return true }
func (h *SystemTraySetLabelMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "label", 0, nil, nil),
	}
}

func (h *SystemTraySetLabelMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "label", 0, nil),
	}
}

func (h *SystemTraySetLabelMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
