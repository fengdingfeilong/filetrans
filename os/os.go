package os

// OS 系统类型
type OS int

const (
	// Win windows
	Win OS = iota
	// Win7 windows7
	Win7
	// Win10 windows
	Win10
	// MacOS darwin 苹果PC系统
	MacOS
	// iOS 苹果手机系统
	iOS //TODO: 小写，不可导出，包外不可用 建议修改为 'IOS'
	// Android 安卓手机系统
	Android
	// linux linux // TODO: 同 iOS
	linux
)
