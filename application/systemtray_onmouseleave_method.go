package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayOnMouseLeaveMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayOnMouseLeaveMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnMouseLeave 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnMouseLeave 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnMouseLeave(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayOnMouseLeaveMethod) GetName() string            { return "onMouseLeave" }
func (h *SystemTrayOnMouseLeaveMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayOnMouseLeaveMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayOnMouseLeaveMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "handler", 0, nil, nil),
	}
}

func (h *SystemTrayOnMouseLeaveMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "handler", 0, nil),
	}
}

func (h *SystemTrayOnMouseLeaveMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
