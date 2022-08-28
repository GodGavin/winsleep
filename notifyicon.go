package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/lxn/walk"
	"image"
	_ "image/png"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"
)

//go:embed static/favicon.png
var iconfile []byte

func GuiInit() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}
	//托盘图标文件
	img, _, err := image.Decode(bytes.NewReader(iconfile))
	icon, _ := walk.NewIconFromImageForDPI(img, 64)
	if err != nil {
		log.Fatal(err)
	}
	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("hotkeysleep \n使用方法:win+s睡眠 win+p休眠\n作者:GavinGao \nQQ:3042752146"); err != nil {
		log.Fatal(err)
	}
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		if err := ni.ShowCustom(
			"hotkeysleep",
			"hotkeysleep 已运行",
			icon); err != nil {
			log.Fatal(err)
		}
	})

	go func() {
		for {
			if 1 == <-ch {
				if err := ni.ShowCustom(
					"现在时间",
					time.Now().Format("15:04:05"),
					icon); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	go func() {
		for {
			if 1 == <-chs {
				Timespeak()
			}
		}
	}()

	aboutAction := walk.NewAction()
	if err := aboutAction.SetText("关于"); err != nil {
		log.Fatal(err)
	}
	//about 实现的功能
	aboutAction.Triggered().Attach(func() {
		walk.MsgBox(mw, "关于",
			"hotkeysleep \n使用方法:win+s睡眠 win+p休眠\n作者:GavinGao \nQQ:3042752146",
			walk.MsgBoxOK)
	}) //如何能够复制提示信息呢？？？

	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出"); err != nil {
		log.Fatal(err)
	}
	//Exit 实现的功能
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	exitAction.Triggered().Attach(func() { os.Exit(1) })

	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}
	if err := ni.ContextMenu().Actions().Add(aboutAction); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}
	if err := ni.ShowInfo("hotkeysleep", "hotkeysleep 运行中"); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}

func Timespeak() {
	fmt.Println("语音报时提示")
	SpeakText("你好,世界!")
}

func SpeakText(text string){
	ttsdll:=syscall.NewLazyDLL("tts.dll")
	speak:=ttsdll.NewProc("rapidSpeakText")
	speak.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}
