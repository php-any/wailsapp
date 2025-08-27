package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIsVisibleMethod struct {
	source applicationsrc.Window
}

func (h *WindowIsVisibleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsVisible()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowIsVisibleMethod) GetName() string            { return "isVisible" }
func (h *WindowIsVisibleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIsVisibleMethod) GetIsStatic() bool          { return true }
func (h *WindowIsVisibleMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIsVisibleMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIsVisibleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
