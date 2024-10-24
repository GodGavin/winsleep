package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"syscall"
)

var ch = make(chan int)
var chs = make(chan int)

func main() {
	//#s::;;
	//	Sleep / Suspend:
	//	DllCall("PowrProf\SetSuspendState", "int", 0, "int", 0, "int", 0)
	go GuiInit()
	fmt.Println("hotkeysleep running")
	deteckeyboard()

	//#p::;;
	//Hibernate:
	//	DllCall("PowrProf\SetSuspendState", "int", 1, "int", 0, "int", 0)

}
func deteckeyboard() {

	//sleep := robotgo.AddEvents("s", "command")
	//if sleep {
	//	fmt.Println("sleep")
	//	//sleepDLLImplementation()
	//}
	fmt.Println("keyboard detecing")
	hook.Register(hook.KeyDown, []string{"command", "s"}, func(e hook.Event) {
		sleepDLLImplementation()
		fmt.Println("sleep execued")
	}) //这个实现不阻塞 good!!!

	hook.Register(hook.KeyDown, []string{"command", "p"}, func(e hook.Event) {
		hibernateDLLImplementation()
		fmt.Println("hibernate execued")
	})

	hook.Register(hook.KeyDown, []string{"command", "t"}, func(e hook.Event) {
		timeImplementation()
		fmt.Println("time execued")
	})

	hook.Register(hook.KeyDown, []string{"command", "m"}, func(e hook.Event) {
		timespeakImplementation()
		fmt.Println("timespeak execued")
	})

	s := hook.Start()

	<-hook.Process(s)

}

//	func deteckeyboard2() {
//		for {
//			hibernate := robotgo.AddEvents("p", "command")
//			if hibernate {
//				fmt.Println("hibernate")
//				//sleepDLLImplementation2()
//			}
//		}
//	}
func sleepDLLImplementation() {
	var mod = syscall.NewLazyDLL("Powrprof.dll")
	var proc = mod.NewProc("SetSuspendState")

	// DLL API : public static extern bool SetSuspendState(bool hiberate, bool forceCritical, bool disableWakeEvent);
	// ex. : uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
	ret, _, _ := proc.Call(
		uintptr(0),
		uintptr(0),
		uintptr(0))

	fmt.Println("Command executed, result code [" + fmt.Sprint(ret) + "]")
}
func hibernateDLLImplementation() {
	var mod = syscall.NewLazyDLL("Powrprof.dll")
	var proc = mod.NewProc("SetSuspendState")

	// DLL API : public static extern bool SetSuspendState(bool hiberate, bool forceCritical, bool disableWakeEvent);
	// ex. : uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
	ret, _, _ := proc.Call(
		uintptr(1),
		uintptr(0),
		uintptr(0))

	fmt.Println("Command executed, result code [" + fmt.Sprint(ret) + "]")
}

func timeImplementation() {
	ch <- 1
}

func timespeakImplementation() {
	chs <- 1
}
