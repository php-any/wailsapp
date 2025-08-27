package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type KeyBindingManagerGetAllMethod struct {
	source *applicationsrc.KeyBindingManager
}

func (h *KeyBindingManagerGetAllMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetAll()
	return data.NewAnyValue(ret0), nil
}

func (h *KeyBindingManagerGetAllMethod) GetName() string            { return "getAll" }
func (h *KeyBindingManagerGetAllMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *KeyBindingManagerGetAllMethod) GetIsStatic() bool          { return true }
func (h *KeyBindingManagerGetAllMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *KeyBindingManagerGetAllMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *KeyBindingManagerGetAllMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
