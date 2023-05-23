package config

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module"
)

func version() {
	PrintVersion()
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

func PrintVersion() {
	ezap.Println("rack Version " + VERSION)
	ezap.Printf("build with Mod: %+v\n", module.ModVersion)
	for _, mv := range MODVERSION {
		ezap.Print(mv + " ")
	}
}
