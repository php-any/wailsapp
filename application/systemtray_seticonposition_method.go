package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTraySetIconPositionMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTraySetIconPositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetIconPosition 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.IconPosition
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.IconPosition(arg0Int)

	ret0 := h.source.SetIconPosition(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTraySetIconPositionMethod) GetName() string            { return "setIconPosition" }
func (h *SystemTraySetIconPositionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTraySetIconPositionMethod) GetIsStatic() bool          { return true }
func (h *SystemTraySetIconPositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "iconPosition", 0, nil, nil),
	}
}

func (h *SystemTraySetIconPositionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "iconPosition", 0, nil),
	}
}

func (h *SystemTraySetIconPositionMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
