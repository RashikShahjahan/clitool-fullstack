package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "magic",
	Short: "Magic creates you boilerplate fullstack app",
	Long:  "Magic creates you boilerplate fullstack app",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("bun", "create", "vite", args[0], "--template", "react")
		output, err := command.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
