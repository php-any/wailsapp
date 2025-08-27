package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppOnShutdownMethod struct {
	source *applicationsrc.App
}

func (h *AppOnShutdownMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("App.OnShutdown 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("App.OnShutdown 参数类型不支持, index: 0"))
	}

	h.source.OnShutdown(arg0)
	return nil, nil
}

func (h *AppOnShutdownMethod) GetName() string            { return "onShutdown" }
func (h *AppOnShutdownMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppOnShutdownMethod) GetIsStatic() bool          { return true }
func (h *AppOnShutdownMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "f", 0, nil, nil),
	}
}

func (h *AppOnShutdownMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "f", 0, nil),
	}
}

func (h *AppOnShutdownMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
