package config

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module"
)

func ShowInfo() {
	PrintVersion()
	PrintMods()
	PrintRack()
}

func PrintRack() {
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
	ezap.Printf("build with Mod: %+v\n", module.ModVersion)
	for _, mv := range MODVERSION {
		ezap.Print(mv + " ")
	}
}

func PrintVersion() {
	ezap.Println("Rack Version: " + VERSION)
}
