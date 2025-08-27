package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppConfigMethod struct {
	source *applicationsrc.App
}

func (h *AppConfigMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Config()
	return data.NewClassValue(NewOptionsClassFrom(&ret0), ctx), nil
}

func (h *AppConfigMethod) GetName() string            { return "config" }
func (h *AppConfigMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppConfigMethod) GetIsStatic() bool          { return true }
func (h *AppConfigMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *AppConfigMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *AppConfigMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
