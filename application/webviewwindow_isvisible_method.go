package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIsVisibleMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIsVisibleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsVisible()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowIsVisibleMethod) GetName() string            { return "isVisible" }
func (h *WebviewWindowIsVisibleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowIsVisibleMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowIsVisibleMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIsVisibleMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIsVisibleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
