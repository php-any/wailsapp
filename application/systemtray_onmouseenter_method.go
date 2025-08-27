package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayOnMouseEnterMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayOnMouseEnterMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnMouseEnter 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.OnMouseEnter 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnMouseEnter(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTrayOnMouseEnterMethod) GetName() string            { return "onMouseEnter" }
func (h *SystemTrayOnMouseEnterMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayOnMouseEnterMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayOnMouseEnterMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "handler", 0, nil, nil),
	}
}

func (h *SystemTrayOnMouseEnterMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "handler", 0, nil),
	}
}

func (h *SystemTrayOnMouseEnterMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
