package ipquery

import (
	"github.com/spf13/cobra"
)

// var cmd = &cobra.Command{
// 	Use:   ID,
// 	Short: Comment,
// 	Long:  Comment,
// 	Run: func(cmd *cobra.Command, args []string) {

// 	},
// }

func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool("ipqury", false, "是否启用 ipquery 查询")
}
