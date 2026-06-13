package uievents

import (
	"errors"

	"github.com/gost-dom/browser/internal/uievents"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// The decoders and attribute getters below are hand-written because the Web-IDL
// code generator emits no-op stubs for UIEventInit/MouseEventInit (their
// dictionaries are not fully derived from IDL, e.g. UIEventInit.view is an
// unexported model field). The constructors in the *_generated.go files call
// these decoders, and InitializeUIEvent/InitializeMouseEvent reference the
// getters by name, so the implementations live in this same package.

func decodePointerEventInit[T any](
	scope js.Scope[T], options js.Object[T], init *uievents.PointerEventInit,
) error {
	return decodeMouseEventInit(scope, options, &init.MouseEventInit)
}

func decodeMouseEventInit[T any](
	scope js.Scope[T], options js.Object[T], init *uievents.MouseEventInit,
) error {
	return errors.Join(
		decodeUIEventInit(scope, options, &init.UIEventInit),
		js.DecodeInto(scope, &init.ScreenX, options, "screenX", codec.DecodeInt),
		js.DecodeInto(scope, &init.ScreenY, options, "screenY", codec.DecodeInt),
	)
}

func decodeUIEventInit[T any](
	scope js.Scope[T], options js.Object[T], init *uievents.UIEventInit,
) error {
	return js.DecodeInto(scope, &init.Detail, options, "detail", codec.DecodeInt)
}

// UIEvent_detail implements the UIEvent.detail getter.
func UIEvent_detail[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	eventInit, err := codec.RetrieveEventInit[uievents.UIEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	return codec.EncodeInt(cbCtx, eventInit.Detail)
}

// MouseEvent_screenX implements the MouseEvent.screenX getter.
func MouseEvent_screenX[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	eventInit, err := codec.RetrieveEventInit[uievents.MouseEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	return codec.EncodeInt(cbCtx, eventInit.ScreenX)
}

// MouseEvent_screenY implements the MouseEvent.screenY getter.
func MouseEvent_screenY[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	eventInit, err := codec.RetrieveEventInit[uievents.MouseEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	return codec.EncodeInt(cbCtx, eventInit.ScreenY)
}
