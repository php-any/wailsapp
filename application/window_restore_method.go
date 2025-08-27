package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowRestoreMethod struct {
	source applicationsrc.Window
}

func (h *WindowRestoreMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Restore()
	return nil, nil
}

func (h *WindowRestoreMethod) GetName() string            { return "restore" }
func (h *WindowRestoreMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowRestoreMethod) GetIsStatic() bool          { return true }
func (h *WindowRestoreMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowRestoreMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowRestoreMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
