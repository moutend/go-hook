package types

// KBDLLHOOKSTRUCT represents KBDLLHOOKSTRUCT structure.
//
// For more details, see the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-kbdllhookstruct
type KBDLLHOOKSTRUCT struct {
	VKCode      VKCode
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
}

// MSLLHOOKSTRUCT represents MSLLHOOKSTRUCT structure.
//
// For more details, see the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msllhookstruct
type MSLLHOOKSTRUCT struct {
	POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
}

// MSG represents MSG structure.
//
// For more details, refer to the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uint32
	LParam  uint32
	Time    uint32
	POINT
}

// Point represents Point structure.
//
// For more details, refer to the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
type POINT struct {
	X int32
	Y int32
}

// HOOKPROC represents HOOKPROC callback function type.
//
// For more details, see the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/winmsg/using-hooks
//
// Note: you don't have to care about this function unless customize the default the mouse / keyboard hook behavior.
type HOOKPROC func(code int32, wParam, lParam uintptr) uintptr
