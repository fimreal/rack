package config

import (
	"github.com/fimreal/goutils/ezap"
)

func ShowInfo() {
	PrintVersion()
	PrintMods()
	PrintRack()
}

func PrintRack() {
	if AppName != "rack" {
		return
	}
	return
	ezap.Print(
		` ██▀███   ▄▄▄       ▄████▄   ██ ▄█▀
▓██ ▒ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ 
▓██ ░▄█ ▒▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ 
▒██▀▀█▄  ░██▄▄▄▄██ ▒▓▓▄ ▄██▒▓██ █▄ 
░██▓ ▒██▒ ▓█   ▓██▒▒ ▓███▀ ░▒██▒ █▄
░ ▒▓ ░▒▓░ ▒▒   ▓▒█░░ ░▒ ▒  ░▒ ▒▒ ▓▒
  ░▒ ░ ▒░  ▒   ▒▒ ░  ░  ▒   ░ ░▒ ▒░
  ░░   ░   ░   ▒   ░        ░ ░░ ░ 
   ░           ░  ░░ ░      ░  ░   
                   ░    
`)
}

func PrintMods() {
	ezap.Printf("build with Module: %+v\n", GetModVer())
}

func PrintVersion() {
	ezap.Println(AppName + " Version: " + Version)
	ezap.Println("Build Time: " + BuildTime)
}
