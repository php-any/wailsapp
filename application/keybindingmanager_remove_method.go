package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type KeyBindingManagerRemoveMethod struct {
	source *applicationsrc.KeyBindingManager
}

func (h *KeyBindingManagerRemoveMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Remove 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.Remove(arg0)
	return nil, nil
}

func (h *KeyBindingManagerRemoveMethod) GetName() string            { return "remove" }
func (h *KeyBindingManagerRemoveMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *KeyBindingManagerRemoveMethod) GetIsStatic() bool          { return true }
func (h *KeyBindingManagerRemoveMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "accelerator", 0, nil, nil),
	}
}

func (h *KeyBindingManagerRemoveMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "accelerator", 0, nil),
	}
}

func (h *KeyBindingManagerRemoveMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
