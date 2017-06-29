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
