package types

// KeyboardEvent contains information about keyboard input event.
type KeyboardEvent struct {
	Message Message
	KBDLLHOOKSTRUCT
}

// MouseEvent contains information about mouse input event.
type MouseEvent struct {
	Message Message
	MSLLHOOKSTRUCT
}
