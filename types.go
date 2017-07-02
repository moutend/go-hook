package hook

// HMODULE
type HMODULE uintptr

// HHOOK
type HHOOK uintptr

// HWND
type HWND uintptr

// Point corresponds to Point structure.
// For more information, see the documentation on MSDN.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/dd162805(v=vs.85).aspx
type POINT struct {
	X int32
	Y int32
}

// MSG corresponds to MSG structure.
// For more information, see the documentation on MSDN.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644958(v=vs.85).aspx
type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uint32
	LParam  uint32
	Time    uint32
	POINT
}
