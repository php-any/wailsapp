package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowRestoreMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowRestoreMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Restore()
	return nil, nil
}

func (h *WebviewWindowRestoreMethod) GetName() string            { return "restore" }
func (h *WebviewWindowRestoreMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowRestoreMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowRestoreMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowRestoreMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowRestoreMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
