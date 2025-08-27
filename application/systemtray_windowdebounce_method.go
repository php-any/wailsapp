package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	"time"
)

type SystemTrayWindowDebounceMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayWindowDebounceMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.WindowDebounce 缺少参数, index: 0"))
	}

	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := time.Duration(arg0Int)

	ret0 := h.source.WindowDebounce(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayWindowDebounceMethod) GetName() string            { return "windowDebounce" }
func (h *SystemTrayWindowDebounceMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayWindowDebounceMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayWindowDebounceMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "debounce", 0, nil, nil),
	}
}

func (h *SystemTrayWindowDebounceMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "debounce", 0, nil),
	}
}

func (h *SystemTrayWindowDebounceMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
