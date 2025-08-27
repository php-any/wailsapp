package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayWindowOffsetMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayWindowOffsetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.WindowOffset 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.WindowOffset(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayWindowOffsetMethod) GetName() string            { return "windowOffset" }
func (h *SystemTrayWindowOffsetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayWindowOffsetMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayWindowOffsetMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "offset", 0, nil, nil),
	}
}

func (h *SystemTrayWindowOffsetMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "offset", 0, nil),
	}
}

func (h *SystemTrayWindowOffsetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
