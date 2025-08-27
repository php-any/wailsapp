package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayOnClickMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayOnClickMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnClick 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnClick 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnClick(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayOnClickMethod) GetName() string            { return "onClick" }
func (h *SystemTrayOnClickMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayOnClickMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayOnClickMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "handler", 0, nil, nil),
	}
}

func (h *SystemTrayOnClickMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "handler", 0, nil),
	}
}

func (h *SystemTrayOnClickMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
