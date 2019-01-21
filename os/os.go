package os

// OS system type
type OS int

const (
	// Win windows
	Win OS = iota
	// Win7 windows7
	Win7
	// Win10 windows10
	Win10
	// MacOS mac os
	MacOS
	// IOS iphone syetem
	IOS
	// Android Android
	Android
	// Linux linux
	Linux
)
