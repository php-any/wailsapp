package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type KeyBindingManagerAddMethod struct {
	source *applicationsrc.KeyBindingManager
}

func (h *KeyBindingManagerAddMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Add 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Add 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	var arg1 func(applicationsrc.Window)
	switch fnv := a1.(type) {
	case *data.FuncValue:
		arg1 = func(p0 applicationsrc.Window) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Add 参数类型不支持, index: 1"))
	}

	h.source.Add(arg0, arg1)
	return nil, nil
}

func (h *KeyBindingManagerAddMethod) GetName() string            { return "add" }
func (h *KeyBindingManagerAddMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *KeyBindingManagerAddMethod) GetIsStatic() bool          { return true }
func (h *KeyBindingManagerAddMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "accelerator", 0, nil, nil),
		node.NewParameter(nil, "callback", 1, nil, nil),
	}
}

func (h *KeyBindingManagerAddMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "accelerator", 0, nil),
		node.NewVariable(nil, "callback", 1, nil),
	}
}

func (h *KeyBindingManagerAddMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
