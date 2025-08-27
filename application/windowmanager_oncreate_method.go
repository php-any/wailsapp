package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowManagerOnCreateMethod struct {
	source *applicationsrc.WindowManager
}

func (h *WindowManagerOnCreateMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.OnCreate 缺少参数, index: 0"))
	}

	var arg0 func(applicationsrc.Window)
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func(p0 applicationsrc.Window) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("WindowManager.OnCreate 参数类型不支持, index: 0"))
	}

	h.source.OnCreate(arg0)
	return nil, nil
}

func (h *WindowManagerOnCreateMethod) GetName() string            { return "onCreate" }
func (h *WindowManagerOnCreateMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowManagerOnCreateMethod) GetIsStatic() bool          { return true }
func (h *WindowManagerOnCreateMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "callback", 0, nil, nil),
	}
}

func (h *WindowManagerOnCreateMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "callback", 0, nil),
	}
}

func (h *WindowManagerOnCreateMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
