package winhook

type HMODULE uintptr
type HHOOK uintptr
type HWND uintptr
type POINT struct {
	X int32
	Y int32
}
type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uint32
	LParam  uint32
	Time    uint32
	POINT
}
type KBDLLHOOKSTRUCT struct {
	VKCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
}
type MSLLHOOKSTRUCT struct {
	POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
	uint32
}
