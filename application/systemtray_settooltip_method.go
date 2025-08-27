package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTraySetTooltipMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTraySetTooltipMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetTooltip 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.SetTooltip(arg0)
	return nil, nil
}

func (h *SystemTraySetTooltipMethod) GetName() string            { return "setTooltip" }
func (h *SystemTraySetTooltipMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTraySetTooltipMethod) GetIsStatic() bool          { return true }
func (h *SystemTraySetTooltipMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "tooltip", 0, nil, nil),
	}
}

func (h *SystemTraySetTooltipMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "tooltip", 0, nil),
	}
}

func (h *SystemTraySetTooltipMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
