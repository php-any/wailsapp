package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerQuestionMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerQuestionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Question()
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *DialogManagerQuestionMethod) GetName() string            { return "question" }
func (h *DialogManagerQuestionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *DialogManagerQuestionMethod) GetIsStatic() bool          { return true }
func (h *DialogManagerQuestionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *DialogManagerQuestionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *DialogManagerQuestionMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
