package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayOnDoubleClickMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayOnDoubleClickMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnDoubleClick 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnDoubleClick 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnDoubleClick(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayOnDoubleClickMethod) GetName() string            { return "onDoubleClick" }
func (h *SystemTrayOnDoubleClickMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayOnDoubleClickMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayOnDoubleClickMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "handler", 0, nil, nil),
	}
}

func (h *SystemTrayOnDoubleClickMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "handler", 0, nil),
	}
}

func (h *SystemTrayOnDoubleClickMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
