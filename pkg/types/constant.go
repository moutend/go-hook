//go:generate stringer -type=Hook
//go:generate stringer -type=Message
//go:generate stringer -type=VKCode

package types

// Hook represents Windows hook types.
type Hook uintptr

const (
	WH_JOURNALRECORD   Hook = 0
	WH_JOURNALPLAYBACK Hook = 1
	WH_KEYBOARD        Hook = 2
	WH_GETMESSAGE      Hook = 3
	WH_CALLWNDPROC     Hook = 4
	WH_CBT             Hook = 5
	WH_SYSMSGFILTER    Hook = 6
	WH_MOUSE           Hook = 7
	WH_DEBUG           Hook = 9
	WH_SHELL           Hook = 10
	WH_FOREGROUNDIDLE  Hook = 11
	WH_CALLWNDPROCRET  Hook = 12
	WH_KEYBOARD_LL     Hook = 13
	WH_MOUSE_LL        Hook = 14
)

// Message represents Windows events.
type Message uintptr

const (
	WM_LBUTTONDOWN Message = 0x0201
	WM_LBUTTONUP   Message = 0x0202
	WM_MOUSEMOVE   Message = 0x0200
	WM_MOUSEWHEEL  Message = 0x020A
	WM_MOUSEHWHEEL Message = 0x020E
	WM_RBUTTONDOWN Message = 0x0204
	WM_RBUTTONUP   Message = 0x0205
	WM_KEYDOWN     Message = 0x0100
	WM_KEYUP       Message = 0x0101
	WM_SYSKEYDOWN  Message = 0x0104
	WM_SYSKEYUP    Message = 0x0105
)

// VKCode represents Microsoft defined virtual key codes.
//
// For more details, see the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
type VKCode uint32

const (
	VK_LBUTTON             VKCode = 0x01 // Left mouse button
	VK_RBUTTON             VKCode = 0x02 // Right mouse button
	VK_CANCEL              VKCode = 0x03 // Control-break processing
	VK_MBUTTON             VKCode = 0x04 // Middle mouse button (three-button mouse)
	VK_XBUTTON1            VKCode = 0x05 // X1 mouse button
	VK_XBUTTON2            VKCode = 0x06 // X2 mouse button
	VK_BACK                VKCode = 0x08 // BACKSPACE key
	VK_TAB                 VKCode = 0x09 // TAB key
	VK_CLEAR               VKCode = 0x0C // CLEAR key
	VK_RETURN              VKCode = 0x0D // ENTER key
	VK_SHIFT               VKCode = 0x10 // SHIFT key
	VK_CONTROL             VKCode = 0x11 // CTRL key
	VK_MENU                VKCode = 0x12 // ALT key
	VK_PAUSE               VKCode = 0x13 // PAUSE key
	VK_CAPITAL             VKCode = 0x14 // CAPS LOCK key
	VK_KANA                VKCode = 0x15 // IME Kana mode
	VK_HANGUEL             VKCode = 0x15 // IME Hanguel mode (maintained for compatibility; use VK_HANGUL)
	VK_HANGUL              VKCode = 0x15 // IME Hangul mode
	VK_IME_ON              VKCode = 0x16 // IME On
	VK_JUNJA               VKCode = 0x17 // IME Junja mode
	VK_FINAL               VKCode = 0x18 // IME final mode
	VK_HANJA               VKCode = 0x19 // IME Hanja mode
	VK_KANJI               VKCode = 0x19 // IME Kanji mode
	VK_IME_OFF             VKCode = 0x1A // IME Off
	VK_ESCAPE              VKCode = 0x1B // ESC key
	VK_CONVERT             VKCode = 0x1C // IME convert
	VK_NONCONVERT          VKCode = 0x1D // IME nonconvert
	VK_ACCEPT              VKCode = 0x1E // IME accept
	VK_MODECHANGE          VKCode = 0x1F // IME mode change request
	VK_SPACE               VKCode = 0x20 // SPACEBAR
	VK_PRIOR               VKCode = 0x21 // PAGE UP key
	VK_NEXT                VKCode = 0x22 // PAGE DOWN key
	VK_END                 VKCode = 0x23 // END key
	VK_HOME                VKCode = 0x24 // HOME key
	VK_LEFT                VKCode = 0x25 // LEFT ARROW key
	VK_UP                  VKCode = 0x26 // UP ARROW key
	VK_RIGHT               VKCode = 0x27 // RIGHT ARROW key
	VK_DOWN                VKCode = 0x28 // DOWN ARROW key
	VK_SELECT              VKCode = 0x29 // SELECT key
	VK_PRINT               VKCode = 0x2A // PRINT key
	VK_EXECUTE             VKCode = 0x2B // EXECUTE key
	VK_SNAPSHOT            VKCode = 0x2C // PRINT SCREEN key
	VK_INSERT              VKCode = 0x2D // INS key
	VK_DELETE              VKCode = 0x2E // DEL key
	VK_HELP                VKCode = 0x2F // HELP key
	VK_0                   VKCode = 0x30 // 0 key
	VK_1                   VKCode = 0x31 // 1 key
	VK_2                   VKCode = 0x32 // 2 key
	VK_3                   VKCode = 0x33 // 3 key
	VK_4                   VKCode = 0x34 // 4 key
	VK_5                   VKCode = 0x35 // 5 key
	VK_6                   VKCode = 0x36 // 6 key
	VK_7                   VKCode = 0x37 // 7 key
	VK_8                   VKCode = 0x38 // 8 key
	VK_9                   VKCode = 0x39 // 9 key
	VK_A                   VKCode = 0x41 // A key
	VK_B                   VKCode = 0x42 // B key
	VK_C                   VKCode = 0x43 // C key
	VK_D                   VKCode = 0x44 // D key
	VK_E                   VKCode = 0x45 // E key
	VK_F                   VKCode = 0x46 // F key
	VK_G                   VKCode = 0x47 // G key
	VK_H                   VKCode = 0x48 // H key
	VK_I                   VKCode = 0x49 // I key
	VK_J                   VKCode = 0x4A // J key
	VK_K                   VKCode = 0x4B // K key
	VK_L                   VKCode = 0x4C // L key
	VK_M                   VKCode = 0x4D // M key
	VK_N                   VKCode = 0x4E // N key
	VK_O                   VKCode = 0x4F // O key
	VK_P                   VKCode = 0x50 // P key
	VK_Q                   VKCode = 0x51 // Q key
	VK_R                   VKCode = 0x52 // R key
	VK_S                   VKCode = 0x53 // S key
	VK_T                   VKCode = 0x54 // T key
	VK_U                   VKCode = 0x55 // U key
	VK_V                   VKCode = 0x56 // V key
	VK_W                   VKCode = 0x57 // W key
	VK_X                   VKCode = 0x58 // X key
	VK_Y                   VKCode = 0x59 // Y key
	VK_Z                   VKCode = 0x5A // Z key
	VK_LWIN                VKCode = 0x5B // Left Windows key (Natural keyboard)
	VK_RWIN                VKCode = 0x5C // Right Windows key (Natural keyboard)
	VK_APPS                VKCode = 0x5D // Applications key (Natural keyboard)
	VK_SLEEP               VKCode = 0x5F // Computer Sleep key
	VK_NUMPAD0             VKCode = 0x60 // Numeric keypad 0 key
	VK_NUMPAD1             VKCode = 0x61 // Numeric keypad 1 key
	VK_NUMPAD2             VKCode = 0x62 // Numeric keypad 2 key
	VK_NUMPAD3             VKCode = 0x63 // Numeric keypad 3 key
	VK_NUMPAD4             VKCode = 0x64 // Numeric keypad 4 key
	VK_NUMPAD5             VKCode = 0x65 // Numeric keypad 5 key
	VK_NUMPAD6             VKCode = 0x66 // Numeric keypad 6 key
	VK_NUMPAD7             VKCode = 0x67 // Numeric keypad 7 key
	VK_NUMPAD8             VKCode = 0x68 // Numeric keypad 8 key
	VK_NUMPAD9             VKCode = 0x69 // Numeric keypad 9 key
	VK_MULTIPLY            VKCode = 0x6A // Multiply key
	VK_ADD                 VKCode = 0x6B // Add key
	VK_SEPARATOR           VKCode = 0x6C // Separator key
	VK_SUBTRACT            VKCode = 0x6D // Subtract key
	VK_DECIMAL             VKCode = 0x6E // Decimal key
	VK_DIVIDE              VKCode = 0x6F // Divide key
	VK_F1                  VKCode = 0x70 // F1 key
	VK_F2                  VKCode = 0x71 // F2 key
	VK_F3                  VKCode = 0x72 // F3 key
	VK_F4                  VKCode = 0x73 // F4 key
	VK_F5                  VKCode = 0x74 // F5 key
	VK_F6                  VKCode = 0x75 // F6 key
	VK_F7                  VKCode = 0x76 // F7 key
	VK_F8                  VKCode = 0x77 // F8 key
	VK_F9                  VKCode = 0x78 // F9 key
	VK_F10                 VKCode = 0x79 // F10 key
	VK_F11                 VKCode = 0x7A // F11 key
	VK_F12                 VKCode = 0x7B // F12 key
	VK_F13                 VKCode = 0x7C // F13 key
	VK_F14                 VKCode = 0x7D // F14 key
	VK_F15                 VKCode = 0x7E // F15 key
	VK_F16                 VKCode = 0x7F // F16 key
	VK_F17                 VKCode = 0x80 // F17 key
	VK_F18                 VKCode = 0x81 // F18 key
	VK_F19                 VKCode = 0x82 // F19 key
	VK_F20                 VKCode = 0x83 // F20 key
	VK_F21                 VKCode = 0x84 // F21 key
	VK_F22                 VKCode = 0x85 // F22 key
	VK_F23                 VKCode = 0x86 // F23 key
	VK_F24                 VKCode = 0x87 // F24 key
	VK_NUMLOCK             VKCode = 0x90 // NUM LOCK key
	VK_SCROLL              VKCode = 0x91 // SCROLL LOCK key
	VK_LSHIFT              VKCode = 0xA0 // Left SHIFT key
	VK_RSHIFT              VKCode = 0xA1 // Right SHIFT key
	VK_LCONTROL            VKCode = 0xA2 // Left CONTROL key
	VK_RCONTROL            VKCode = 0xA3 // Right CONTROL key
	VK_LMENU               VKCode = 0xA4 // Left MENU key
	VK_RMENU               VKCode = 0xA5 // Right MENU key
	VK_BROWSER_BACK        VKCode = 0xA6 // Browser Back key
	VK_BROWSER_FORWARD     VKCode = 0xA7 // Browser Forward key
	VK_BROWSER_REFRESH     VKCode = 0xA8 // Browser Refresh key
	VK_BROWSER_STOP        VKCode = 0xA9 // Browser Stop key
	VK_BROWSER_SEARCH      VKCode = 0xAA // Browser Search key
	VK_BROWSER_FAVORITES   VKCode = 0xAB // Browser Favorites key
	VK_BROWSER_HOME        VKCode = 0xAC // Browser Start and Home key
	VK_VOLUME_MUTE         VKCode = 0xAD // Volume Mute key
	VK_VOLUME_DOWN         VKCode = 0xAE // Volume Down key
	VK_VOLUME_UP           VKCode = 0xAF // Volume Up key
	VK_MEDIA_NEXT_TRACK    VKCode = 0xB0 // Next Track key
	VK_MEDIA_PREV_TRACK    VKCode = 0xB1 // Previous Track key
	VK_MEDIA_STOP          VKCode = 0xB2 // Stop Media key
	VK_MEDIA_PLAY_PAUSE    VKCode = 0xB3 // Play/Pause Media key
	VK_LAUNCH_MAIL         VKCode = 0xB4 // Start Mail key
	VK_LAUNCH_MEDIA_SELECT VKCode = 0xB5 // Select Media key
	VK_LAUNCH_APP1         VKCode = 0xB6 // Start Application 1 key
	VK_LAUNCH_APP2         VKCode = 0xB7 // Start Application 2 key
	VK_OEM_1               VKCode = 0xBA // Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the ';:' key
	VK_OEM_PLUS            VKCode = 0xBB // For any country/region, the '+' key
	VK_OEM_COMMA           VKCode = 0xBC // For any country/region, the ',' key
	VK_OEM_MINUS           VKCode = 0xBD // For any country/region, the '-' key
	VK_OEM_PERIOD          VKCode = 0xBE // For any country/region, the '.' key
	VK_OEM_2               VKCode = 0xBF // Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the '/?' key
	VK_OEM_3               VKCode = 0xC0 // Used for miscellaneous characters; it can vary by keyboard.  For the US standard keyboard, the '`~' key
	VK_OEM_4               VKCode = 0xDB // Used for miscellaneous characters; it can vary by keyboard.  For the US standard keyboard, the '[{' key
	VK_OEM_5               VKCode = 0xDC // Used for miscellaneous characters; it can vary by keyboard.  For the US standard keyboard, the '\|' key
	VK_OEM_6               VKCode = 0xDD // Used for miscellaneous characters; it can vary by keyboard. For the US standard keyboard, the ']}' key
	VK_OEM_7               VKCode = 0xDE // Used for miscellaneous characters; it can vary by keyboard.  For the US standard keyboard, the 'single-quote/double-quote' key
	VK_OEM_8               VKCode = 0xDF // Used for miscellaneous characters; it can vary by keyboard.
	VK_OEM_102             VKCode = 0xE2 // Either the angle bracket key or the backslash key on the RT 102-key keyboard
	VK_PROCESSKEY          VKCode = 0xE5 // IME PROCESS key
	VK_PACKET              VKCode = 0xE7 // Used to pass Unicode characters as if they were keystrokes. The VK_PACKET key is the low word of a 32-bit Virtual Key value used for non-keyboard input methods. For more information, see Remark in KEYBDINPUT, SendInput, WM_KEYDOWN, and WM_KEYUP
	VK_ATTN                VKCode = 0xF6 // Attn key
	VK_CRSEL               VKCode = 0xF7 // CrSel key
	VK_EXSEL               VKCode = 0xF8 // ExSel key
	VK_EREOF               VKCode = 0xF9 // Erase EOF key
	VK_PLAY                VKCode = 0xFA // Play key
	VK_ZOOM                VKCode = 0xFB // Zoom key
	VK_NONAME              VKCode = 0xFC // Reserved
	VK_PA1                 VKCode = 0xFD // PA1 key
	VK_OEM_CLEAR           VKCode = 0xFE // Clear key
)
