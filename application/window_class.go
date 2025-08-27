package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewWindowClassFrom(source applicationsrc.Window) data.ClassStmt {
	return &WindowClass{
		source:                         source,
		construct:                      &WindowConstructMethod{source: source},
		bounds:                         &WindowBoundsMethod{source: source},
		callError:                      &WindowCallErrorMethod{source: source},
		callResponse:                   &WindowCallResponseMethod{source: source},
		center:                         &WindowCenterMethod{source: source},
		close:                          &WindowCloseMethod{source: source},
		dialogError:                    &WindowDialogErrorMethod{source: source},
		dialogResponse:                 &WindowDialogResponseMethod{source: source},
		disableSizeConstraints:         &WindowDisableSizeConstraintsMethod{source: source},
		dispatchWailsEvent:             &WindowDispatchWailsEventMethod{source: source},
		emitEvent:                      &WindowEmitEventMethod{source: source},
		enableSizeConstraints:          &WindowEnableSizeConstraintsMethod{source: source},
		error:                          &WindowErrorMethod{source: source},
		execJS:                         &WindowExecJSMethod{source: source},
		flash:                          &WindowFlashMethod{source: source},
		focus:                          &WindowFocusMethod{source: source},
		forceReload:                    &WindowForceReloadMethod{source: source},
		fullscreen:                     &WindowFullscreenMethod{source: source},
		getBorderSizes:                 &WindowGetBorderSizesMethod{source: source},
		getScreen:                      &WindowGetScreenMethod{source: source},
		getZoom:                        &WindowGetZoomMethod{source: source},
		handleDragAndDropMessage:       &WindowHandleDragAndDropMessageMethod{source: source},
		handleKeyEvent:                 &WindowHandleKeyEventMethod{source: source},
		handleMessage:                  &WindowHandleMessageMethod{source: source},
		handleWindowEvent:              &WindowHandleWindowEventMethod{source: source},
		height:                         &WindowHeightMethod{source: source},
		hide:                           &WindowHideMethod{source: source},
		hideMenuBar:                    &WindowHideMenuBarMethod{source: source},
		iD:                             &WindowIDMethod{source: source},
		info:                           &WindowInfoMethod{source: source},
		initiateFrontendDropProcessing: &WindowInitiateFrontendDropProcessingMethod{source: source},
		isFocused:                      &WindowIsFocusedMethod{source: source},
		isFullscreen:                   &WindowIsFullscreenMethod{source: source},
		isIgnoreMouseEvents:            &WindowIsIgnoreMouseEventsMethod{source: source},
		isMaximised:                    &WindowIsMaximisedMethod{source: source},
		isMinimised:                    &WindowIsMinimisedMethod{source: source},
		isVisible:                      &WindowIsVisibleMethod{source: source},
		maximise:                       &WindowMaximiseMethod{source: source},
		minimise:                       &WindowMinimiseMethod{source: source},
		name:                           &WindowNameMethod{source: source},
		nativeWindow:                   &WindowNativeWindowMethod{source: source},
		onWindowEvent:                  &WindowOnWindowEventMethod{source: source},
		openContextMenu:                &WindowOpenContextMenuMethod{source: source},
		openDevTools:                   &WindowOpenDevToolsMethod{source: source},
		position:                       &WindowPositionMethod{source: source},
		print:                          &WindowPrintMethod{source: source},
		registerHook:                   &WindowRegisterHookMethod{source: source},
		relativePosition:               &WindowRelativePositionMethod{source: source},
		reload:                         &WindowReloadMethod{source: source},
		resizable:                      &WindowResizableMethod{source: source},
		restore:                        &WindowRestoreMethod{source: source},
		run:                            &WindowRunMethod{source: source},
		setAlwaysOnTop:                 &WindowSetAlwaysOnTopMethod{source: source},
		setBackgroundColour:            &WindowSetBackgroundColourMethod{source: source},
		setBounds:                      &WindowSetBoundsMethod{source: source},
		setCloseButtonState:            &WindowSetCloseButtonStateMethod{source: source},
		setContentProtection:           &WindowSetContentProtectionMethod{source: source},
		setEnabled:                     &WindowSetEnabledMethod{source: source},
		setFrameless:                   &WindowSetFramelessMethod{source: source},
		setHTML:                        &WindowSetHTMLMethod{source: source},
		setIgnoreMouseEvents:           &WindowSetIgnoreMouseEventsMethod{source: source},
		setMaxSize:                     &WindowSetMaxSizeMethod{source: source},
		setMaximiseButtonState:         &WindowSetMaximiseButtonStateMethod{source: source},
		setMenu:                        &WindowSetMenuMethod{source: source},
		setMinSize:                     &WindowSetMinSizeMethod{source: source},
		setMinimiseButtonState:         &WindowSetMinimiseButtonStateMethod{source: source},
		setPosition:                    &WindowSetPositionMethod{source: source},
		setRelativePosition:            &WindowSetRelativePositionMethod{source: source},
		setResizable:                   &WindowSetResizableMethod{source: source},
		setSize:                        &WindowSetSizeMethod{source: source},
		setTitle:                       &WindowSetTitleMethod{source: source},
		setURL:                         &WindowSetURLMethod{source: source},
		setZoom:                        &WindowSetZoomMethod{source: source},
		show:                           &WindowShowMethod{source: source},
		showMenuBar:                    &WindowShowMenuBarMethod{source: source},
		size:                           &WindowSizeMethod{source: source},
		snapAssist:                     &WindowSnapAssistMethod{source: source},
		toggleFrameless:                &WindowToggleFramelessMethod{source: source},
		toggleFullscreen:               &WindowToggleFullscreenMethod{source: source},
		toggleMaximise:                 &WindowToggleMaximiseMethod{source: source},
		toggleMenuBar:                  &WindowToggleMenuBarMethod{source: source},
		unFullscreen:                   &WindowUnFullscreenMethod{source: source},
		unMaximise:                     &WindowUnMaximiseMethod{source: source},
		unMinimise:                     &WindowUnMinimiseMethod{source: source},
		width:                          &WindowWidthMethod{source: source},
		zoom:                           &WindowZoomMethod{source: source},
		zoomIn:                         &WindowZoomInMethod{source: source},
		zoomOut:                        &WindowZoomOutMethod{source: source},
		zoomReset:                      &WindowZoomResetMethod{source: source},
	}
}

type WindowClass struct {
	node.Node
	source                         applicationsrc.Window
	construct                      data.Method
	bounds                         data.Method
	callError                      data.Method
	callResponse                   data.Method
	center                         data.Method
	close                          data.Method
	dialogError                    data.Method
	dialogResponse                 data.Method
	disableSizeConstraints         data.Method
	dispatchWailsEvent             data.Method
	emitEvent                      data.Method
	enableSizeConstraints          data.Method
	error                          data.Method
	execJS                         data.Method
	flash                          data.Method
	focus                          data.Method
	forceReload                    data.Method
	fullscreen                     data.Method
	getBorderSizes                 data.Method
	getScreen                      data.Method
	getZoom                        data.Method
	handleDragAndDropMessage       data.Method
	handleKeyEvent                 data.Method
	handleMessage                  data.Method
	handleWindowEvent              data.Method
	height                         data.Method
	hide                           data.Method
	hideMenuBar                    data.Method
	iD                             data.Method
	info                           data.Method
	initiateFrontendDropProcessing data.Method
	isFocused                      data.Method
	isFullscreen                   data.Method
	isIgnoreMouseEvents            data.Method
	isMaximised                    data.Method
	isMinimised                    data.Method
	isVisible                      data.Method
	maximise                       data.Method
	minimise                       data.Method
	name                           data.Method
	nativeWindow                   data.Method
	onWindowEvent                  data.Method
	openContextMenu                data.Method
	openDevTools                   data.Method
	position                       data.Method
	print                          data.Method
	registerHook                   data.Method
	relativePosition               data.Method
	reload                         data.Method
	resizable                      data.Method
	restore                        data.Method
	run                            data.Method
	setAlwaysOnTop                 data.Method
	setBackgroundColour            data.Method
	setBounds                      data.Method
	setCloseButtonState            data.Method
	setContentProtection           data.Method
	setEnabled                     data.Method
	setFrameless                   data.Method
	setHTML                        data.Method
	setIgnoreMouseEvents           data.Method
	setMaxSize                     data.Method
	setMaximiseButtonState         data.Method
	setMenu                        data.Method
	setMinSize                     data.Method
	setMinimiseButtonState         data.Method
	setPosition                    data.Method
	setRelativePosition            data.Method
	setResizable                   data.Method
	setSize                        data.Method
	setTitle                       data.Method
	setURL                         data.Method
	setZoom                        data.Method
	show                           data.Method
	showMenuBar                    data.Method
	size                           data.Method
	snapAssist                     data.Method
	toggleFrameless                data.Method
	toggleFullscreen               data.Method
	toggleMaximise                 data.Method
	toggleMenuBar                  data.Method
	unFullscreen                   data.Method
	unMaximise                     data.Method
	unMinimise                     data.Method
	width                          data.Method
	zoom                           data.Method
	zoomIn                         data.Method
	zoomOut                        data.Method
	zoomReset                      data.Method
}

func (s *WindowClass) GetValue(_ data.Context) (data.GetValue, data.Control) {
	clone := *s
	return &clone, nil
}
func (s *WindowClass) GetName() string                            { return "wails\\application\\Window" }
func (s *WindowClass) GetExtend() *string                         { return nil }
func (s *WindowClass) GetImplements() []string                    { return nil }
func (s *WindowClass) AsString() string                           { return "Window{}" }
func (s *WindowClass) GetSource() any                             { return s.source }
func (s *WindowClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *WindowClass) GetProperties() map[string]data.Property    { return nil }
func (s *WindowClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "bounds":
		return s.bounds, true
	case "callError":
		return s.callError, true
	case "callResponse":
		return s.callResponse, true
	case "center":
		return s.center, true
	case "close":
		return s.close, true
	case "dialogError":
		return s.dialogError, true
	case "dialogResponse":
		return s.dialogResponse, true
	case "disableSizeConstraints":
		return s.disableSizeConstraints, true
	case "dispatchWailsEvent":
		return s.dispatchWailsEvent, true
	case "emitEvent":
		return s.emitEvent, true
	case "enableSizeConstraints":
		return s.enableSizeConstraints, true
	case "error":
		return s.error, true
	case "execJS":
		return s.execJS, true
	case "flash":
		return s.flash, true
	case "focus":
		return s.focus, true
	case "forceReload":
		return s.forceReload, true
	case "fullscreen":
		return s.fullscreen, true
	case "getBorderSizes":
		return s.getBorderSizes, true
	case "getScreen":
		return s.getScreen, true
	case "getZoom":
		return s.getZoom, true
	case "handleDragAndDropMessage":
		return s.handleDragAndDropMessage, true
	case "handleKeyEvent":
		return s.handleKeyEvent, true
	case "handleMessage":
		return s.handleMessage, true
	case "handleWindowEvent":
		return s.handleWindowEvent, true
	case "height":
		return s.height, true
	case "hide":
		return s.hide, true
	case "hideMenuBar":
		return s.hideMenuBar, true
	case "iD":
		return s.iD, true
	case "info":
		return s.info, true
	case "initiateFrontendDropProcessing":
		return s.initiateFrontendDropProcessing, true
	case "isFocused":
		return s.isFocused, true
	case "isFullscreen":
		return s.isFullscreen, true
	case "isIgnoreMouseEvents":
		return s.isIgnoreMouseEvents, true
	case "isMaximised":
		return s.isMaximised, true
	case "isMinimised":
		return s.isMinimised, true
	case "isVisible":
		return s.isVisible, true
	case "maximise":
		return s.maximise, true
	case "minimise":
		return s.minimise, true
	case "name":
		return s.name, true
	case "nativeWindow":
		return s.nativeWindow, true
	case "onWindowEvent":
		return s.onWindowEvent, true
	case "openContextMenu":
		return s.openContextMenu, true
	case "openDevTools":
		return s.openDevTools, true
	case "position":
		return s.position, true
	case "print":
		return s.print, true
	case "registerHook":
		return s.registerHook, true
	case "relativePosition":
		return s.relativePosition, true
	case "reload":
		return s.reload, true
	case "resizable":
		return s.resizable, true
	case "restore":
		return s.restore, true
	case "run":
		return s.run, true
	case "setAlwaysOnTop":
		return s.setAlwaysOnTop, true
	case "setBackgroundColour":
		return s.setBackgroundColour, true
	case "setBounds":
		return s.setBounds, true
	case "setCloseButtonState":
		return s.setCloseButtonState, true
	case "setContentProtection":
		return s.setContentProtection, true
	case "setEnabled":
		return s.setEnabled, true
	case "setFrameless":
		return s.setFrameless, true
	case "setHTML":
		return s.setHTML, true
	case "setIgnoreMouseEvents":
		return s.setIgnoreMouseEvents, true
	case "setMaxSize":
		return s.setMaxSize, true
	case "setMaximiseButtonState":
		return s.setMaximiseButtonState, true
	case "setMenu":
		return s.setMenu, true
	case "setMinSize":
		return s.setMinSize, true
	case "setMinimiseButtonState":
		return s.setMinimiseButtonState, true
	case "setPosition":
		return s.setPosition, true
	case "setRelativePosition":
		return s.setRelativePosition, true
	case "setResizable":
		return s.setResizable, true
	case "setSize":
		return s.setSize, true
	case "setTitle":
		return s.setTitle, true
	case "setURL":
		return s.setURL, true
	case "setZoom":
		return s.setZoom, true
	case "show":
		return s.show, true
	case "showMenuBar":
		return s.showMenuBar, true
	case "size":
		return s.size, true
	case "snapAssist":
		return s.snapAssist, true
	case "toggleFrameless":
		return s.toggleFrameless, true
	case "toggleFullscreen":
		return s.toggleFullscreen, true
	case "toggleMaximise":
		return s.toggleMaximise, true
	case "toggleMenuBar":
		return s.toggleMenuBar, true
	case "unFullscreen":
		return s.unFullscreen, true
	case "unMaximise":
		return s.unMaximise, true
	case "unMinimise":
		return s.unMinimise, true
	case "width":
		return s.width, true
	case "zoom":
		return s.zoom, true
	case "zoomIn":
		return s.zoomIn, true
	case "zoomOut":
		return s.zoomOut, true
	case "zoomReset":
		return s.zoomReset, true
	}
	return nil, false
}

func (s *WindowClass) GetMethods() []data.Method {
	return []data.Method{
		s.bounds,
		s.callError,
		s.callResponse,
		s.center,
		s.close,
		s.dialogError,
		s.dialogResponse,
		s.disableSizeConstraints,
		s.dispatchWailsEvent,
		s.emitEvent,
		s.enableSizeConstraints,
		s.error,
		s.execJS,
		s.flash,
		s.focus,
		s.forceReload,
		s.fullscreen,
		s.getBorderSizes,
		s.getScreen,
		s.getZoom,
		s.handleDragAndDropMessage,
		s.handleKeyEvent,
		s.handleMessage,
		s.handleWindowEvent,
		s.height,
		s.hide,
		s.hideMenuBar,
		s.iD,
		s.info,
		s.initiateFrontendDropProcessing,
		s.isFocused,
		s.isFullscreen,
		s.isIgnoreMouseEvents,
		s.isMaximised,
		s.isMinimised,
		s.isVisible,
		s.maximise,
		s.minimise,
		s.name,
		s.nativeWindow,
		s.onWindowEvent,
		s.openContextMenu,
		s.openDevTools,
		s.position,
		s.print,
		s.registerHook,
		s.relativePosition,
		s.reload,
		s.resizable,
		s.restore,
		s.run,
		s.setAlwaysOnTop,
		s.setBackgroundColour,
		s.setBounds,
		s.setCloseButtonState,
		s.setContentProtection,
		s.setEnabled,
		s.setFrameless,
		s.setHTML,
		s.setIgnoreMouseEvents,
		s.setMaxSize,
		s.setMaximiseButtonState,
		s.setMenu,
		s.setMinSize,
		s.setMinimiseButtonState,
		s.setPosition,
		s.setRelativePosition,
		s.setResizable,
		s.setSize,
		s.setTitle,
		s.setURL,
		s.setZoom,
		s.show,
		s.showMenuBar,
		s.size,
		s.snapAssist,
		s.toggleFrameless,
		s.toggleFullscreen,
		s.toggleMaximise,
		s.toggleMenuBar,
		s.unFullscreen,
		s.unMaximise,
		s.unMinimise,
		s.width,
		s.zoom,
		s.zoomIn,
		s.zoomOut,
		s.zoomReset,
	}
}

func (s *WindowClass) GetConstruct() data.Method { return s.construct }

type WindowConstructMethod struct {
	source applicationsrc.Window
}

func (h *WindowConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *WindowConstructMethod) GetName() string               { return "construct" }
func (h *WindowConstructMethod) GetModifier() data.Modifier    { return data.ModifierPublic }
func (h *WindowConstructMethod) GetIsStatic() bool             { return false }
func (h *WindowConstructMethod) GetParams() []data.GetValue    { return []data.GetValue{} }
func (h *WindowConstructMethod) GetVariables() []data.Variable { return []data.Variable{} }
func (h *WindowConstructMethod) GetReturnType() data.Types     { return data.NewBaseType("void") }
