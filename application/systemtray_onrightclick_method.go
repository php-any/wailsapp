package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayOnRightClickMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayOnRightClickMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnRightClick 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnRightClick 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnRightClick(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayOnRightClickMethod) GetName() string            { return "onRightClick" }
func (h *SystemTrayOnRightClickMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayOnRightClickMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayOnRightClickMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "handler", 0, nil, nil),
	}
}

func (h *SystemTrayOnRightClickMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "handler", 0, nil),
	}
}

func (h *SystemTrayOnRightClickMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
