package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ClipboardManagerTextMethod struct {
	source *applicationsrc.ClipboardManager
}

func (h *ClipboardManagerTextMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.Text()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *ClipboardManagerTextMethod) GetName() string            { return "text" }
func (h *ClipboardManagerTextMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ClipboardManagerTextMethod) GetIsStatic() bool          { return true }
func (h *ClipboardManagerTextMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *ClipboardManagerTextMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *ClipboardManagerTextMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
