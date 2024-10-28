package main

import (
	"context"
	"fmt"
	hook "github.com/robotn/gohook"
	"os"
	"os/signal"
	"syscall"
)

var ch = make(chan int)
var chs = make(chan int)

func main() {
	//#s::;;
	//	Sleep / Suspend:
	//	DllCall("PowrProf\SetSuspendState", "int", 0, "int", 0, "int", 0)
	ctx, cancel := context.WithCancel(context.Background())
	go GuiInit(ctx)
	fmt.Println("hotkeysleep running")
	go deteckeyboard()

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt)
	select {
	case <-done:
		fmt.Println("exit")
		hook.End()
		cancel()
	}
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
		//sleepDLLImplementation()
		fmt.Println("sleep execued")
	}) //这个实现不阻塞 good!!!

	hook.Register(hook.KeyDown, []string{"command", "p"}, func(e hook.Event) {
		//hibernateDLLImplementation()
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

	<-hook.Process(debugprintchan(s))

}

func debugprintchan(s chan hook.Event) (out chan hook.Event) {
	out = make(chan hook.Event)
	var kindtostrng = map[uint8]string{
		hook.KeyDown: "KeyDown",
		hook.KeyUp:   "KeyUp",
		hook.KeyHold: "KeyHold",
	}
	go func() {
		for {
			e, ok := <-s
			if !ok {
				break
			}
			if e.Keycode == 3675 {
				fmt.Println(kindtostrng[e.Kind] + " command")
			}
			if e.Keycode == 29 {
				fmt.Println(kindtostrng[e.Kind] + " control")
			}
			if e.Keycode == 56 {
				fmt.Println(kindtostrng[e.Kind] + " alt")
			}
			if e.Keycode == 42 {
				fmt.Println(kindtostrng[e.Kind] + " shift")
			}
			if e.Keycode == 20 {
				fmt.Println(kindtostrng[e.Kind] + " t")
			}
			if e.Keycode == 31 {
				fmt.Println(kindtostrng[e.Kind] + " s")
			}
			if e.Keycode == 25 {
				fmt.Println(kindtostrng[e.Kind] + " p")
			}
			if e.Keycode == 50 {
				fmt.Println(kindtostrng[e.Kind] + " m")
			}

			//fmt.Println(e.String())
			out <- e
		}
		close(out)
	}()
	return out
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
