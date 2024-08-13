// 内置基础命令，方便 scratch 容器查看文件
package cmd

import (
	"os"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "use rack like unix shell",
	Long: `use rack like unix shell.
avalible commands:
    ls
    pwd`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("tty") {
			// 进入简易交互式 shell
			ezap.Println("enter shell")
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")

	shellCmdsInit()

	viper.BindPFlags(shellCmd.Flags())
}

// 添加可用命令

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list directory contents",
	Long:  `list directory contents`,
	Run: func(cmd *cobra.Command, args []string) {
		var dir string
		if len(args) > 0 {
			dir = args[0]
		} else {
			dir = "." // 默认当前目录
		}
		showHideFile := viper.GetBool("all")
		showLongFormat := viper.GetBool("long")
		ezap.Info("show hide file:", showHideFile)
		ezap.Info("show long format:", showLongFormat)
		utils.ListDirectory(dir, showHideFile, showLongFormat)
	},
}

// pwdCmd 表示 pwd 命令
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "print current working directory",
	Long:  `print current working directory`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取当前工作目录
		dir, err := os.Getwd()
		if err != nil {
			ezap.Printf("Error getting current working directory: %v\n", err)
			return
		}
		ezap.Println(dir) // 打印当前工作目录
	},
}

func shellCmdsInit() {
	shellCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("all", "a", false, "show all file, include hide file starting with .")
	lsCmd.Flags().BoolP("long", "l", false, "use a long listing format")
	// lsCmd.Flags().BoolP("directory", "d", false, "list directories themselves, not their contents")

	viper.BindPFlags(lsCmd.Flags())

	shellCmd.AddCommand(pwdCmd)
}
