package cli

import (
	"fmt"
	"os"

	"github.com/1sl4nds/zenas/internal/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zenas",
	Short: "A Platform for Computational Law.",
	Long:  "A Platform for Computational Law.",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer()

		s.ListenAndServe()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
