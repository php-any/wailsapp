package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayShowMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayShowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Show()
	return nil, nil
}

func (h *SystemTrayShowMethod) GetName() string            { return "show" }
func (h *SystemTrayShowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayShowMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayShowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayShowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayShowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
