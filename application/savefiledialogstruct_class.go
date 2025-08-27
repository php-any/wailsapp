package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewSaveFileDialogStructClass() data.ClassStmt {
	return &SaveFileDialogStructClass{
		source:                          nil,
		construct:                       &SaveFileDialogStructConstructMethod{source: nil},
		addFilter:                       &SaveFileDialogStructAddFilterMethod{source: nil},
		allowsOtherFileTypes:            &SaveFileDialogStructAllowsOtherFileTypesMethod{source: nil},
		attachToWindow:                  &SaveFileDialogStructAttachToWindowMethod{source: nil},
		canCreateDirectories:            &SaveFileDialogStructCanCreateDirectoriesMethod{source: nil},
		canSelectHiddenExtension:        &SaveFileDialogStructCanSelectHiddenExtensionMethod{source: nil},
		hideExtension:                   &SaveFileDialogStructHideExtensionMethod{source: nil},
		promptForSingleSelection:        &SaveFileDialogStructPromptForSingleSelectionMethod{source: nil},
		setButtonText:                   &SaveFileDialogStructSetButtonTextMethod{source: nil},
		setDirectory:                    &SaveFileDialogStructSetDirectoryMethod{source: nil},
		setFilename:                     &SaveFileDialogStructSetFilenameMethod{source: nil},
		setMessage:                      &SaveFileDialogStructSetMessageMethod{source: nil},
		setOptions:                      &SaveFileDialogStructSetOptionsMethod{source: nil},
		showHiddenFiles:                 &SaveFileDialogStructShowHiddenFilesMethod{source: nil},
		treatsFilePackagesAsDirectories: &SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod{source: nil},
	}
}

func NewSaveFileDialogStructClassFrom(source *applicationsrc.SaveFileDialogStruct) data.ClassStmt {
	return &SaveFileDialogStructClass{
		source:                          source,
		construct:                       &SaveFileDialogStructConstructMethod{source: source},
		addFilter:                       &SaveFileDialogStructAddFilterMethod{source: source},
		allowsOtherFileTypes:            &SaveFileDialogStructAllowsOtherFileTypesMethod{source: source},
		attachToWindow:                  &SaveFileDialogStructAttachToWindowMethod{source: source},
		canCreateDirectories:            &SaveFileDialogStructCanCreateDirectoriesMethod{source: source},
		canSelectHiddenExtension:        &SaveFileDialogStructCanSelectHiddenExtensionMethod{source: source},
		hideExtension:                   &SaveFileDialogStructHideExtensionMethod{source: source},
		promptForSingleSelection:        &SaveFileDialogStructPromptForSingleSelectionMethod{source: source},
		setButtonText:                   &SaveFileDialogStructSetButtonTextMethod{source: source},
		setDirectory:                    &SaveFileDialogStructSetDirectoryMethod{source: source},
		setFilename:                     &SaveFileDialogStructSetFilenameMethod{source: source},
		setMessage:                      &SaveFileDialogStructSetMessageMethod{source: source},
		setOptions:                      &SaveFileDialogStructSetOptionsMethod{source: source},
		showHiddenFiles:                 &SaveFileDialogStructShowHiddenFilesMethod{source: source},
		treatsFilePackagesAsDirectories: &SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod{source: source},
	}
}

type SaveFileDialogStructClass struct {
	node.Node
	source                          *applicationsrc.SaveFileDialogStruct
	construct                       data.Method
	addFilter                       data.Method
	allowsOtherFileTypes            data.Method
	attachToWindow                  data.Method
	canCreateDirectories            data.Method
	canSelectHiddenExtension        data.Method
	hideExtension                   data.Method
	promptForSingleSelection        data.Method
	setButtonText                   data.Method
	setDirectory                    data.Method
	setFilename                     data.Method
	setMessage                      data.Method
	setOptions                      data.Method
	showHiddenFiles                 data.Method
	treatsFilePackagesAsDirectories data.Method
}

func (s *SaveFileDialogStructClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewSaveFileDialogStructClassFrom(&applicationsrc.SaveFileDialogStruct{}), ctx.CreateBaseContext()), nil
}
func (s *SaveFileDialogStructClass) GetName() string {
	return "wails\\application\\SaveFileDialogStruct"
}
func (s *SaveFileDialogStructClass) GetExtend() *string                         { return nil }
func (s *SaveFileDialogStructClass) GetImplements() []string                    { return nil }
func (s *SaveFileDialogStructClass) AsString() string                           { return "SaveFileDialogStruct{}" }
func (s *SaveFileDialogStructClass) GetSource() any                             { return s.source }
func (s *SaveFileDialogStructClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *SaveFileDialogStructClass) GetProperties() map[string]data.Property    { return nil }
func (s *SaveFileDialogStructClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "addFilter":
		return s.addFilter, true
	case "allowsOtherFileTypes":
		return s.allowsOtherFileTypes, true
	case "attachToWindow":
		return s.attachToWindow, true
	case "canCreateDirectories":
		return s.canCreateDirectories, true
	case "canSelectHiddenExtension":
		return s.canSelectHiddenExtension, true
	case "hideExtension":
		return s.hideExtension, true
	case "promptForSingleSelection":
		return s.promptForSingleSelection, true
	case "setButtonText":
		return s.setButtonText, true
	case "setDirectory":
		return s.setDirectory, true
	case "setFilename":
		return s.setFilename, true
	case "setMessage":
		return s.setMessage, true
	case "setOptions":
		return s.setOptions, true
	case "showHiddenFiles":
		return s.showHiddenFiles, true
	case "treatsFilePackagesAsDirectories":
		return s.treatsFilePackagesAsDirectories, true
	}
	return nil, false
}

func (s *SaveFileDialogStructClass) GetMethods() []data.Method {
	return []data.Method{
		s.addFilter,
		s.allowsOtherFileTypes,
		s.attachToWindow,
		s.canCreateDirectories,
		s.canSelectHiddenExtension,
		s.hideExtension,
		s.promptForSingleSelection,
		s.setButtonText,
		s.setDirectory,
		s.setFilename,
		s.setMessage,
		s.setOptions,
		s.showHiddenFiles,
		s.treatsFilePackagesAsDirectories,
	}
}

func (s *SaveFileDialogStructClass) GetConstruct() data.Method { return s.construct }

type SaveFileDialogStructConstructMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *SaveFileDialogStructConstructMethod) GetName() string            { return "construct" }
func (h *SaveFileDialogStructConstructMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SaveFileDialogStructConstructMethod) GetIsStatic() bool          { return false }
func (h *SaveFileDialogStructConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *SaveFileDialogStructConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *SaveFileDialogStructConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
