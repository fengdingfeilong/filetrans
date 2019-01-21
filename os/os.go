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
	// MacOS
	MacOS
	// IOS iphone syetem
	IOS
	// Android 安卓手机系统
	Android
	// Linux
	Linux
)
