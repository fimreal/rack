package serve

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/config"
)

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
	ezap.Println("rack Version " + config.VERSION)
	ezap.Println("rack Mod: ")
	for _, mv := range config.MODVERSION {
		ezap.Print(mv + " ")
	}
}
