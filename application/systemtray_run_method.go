package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTrayRunMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTrayRunMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Run()
	return nil, nil
}

func (h *SystemTrayRunMethod) GetName() string            { return "run" }
func (h *SystemTrayRunMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTrayRunMethod) GetIsStatic() bool          { return true }
func (h *SystemTrayRunMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SystemTrayRunMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SystemTrayRunMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
