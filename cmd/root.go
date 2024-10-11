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
		createDir := exec.Command("mkdir", args[0])
		createVite := exec.Command("bun", "create", "vite", "frontend", "--template", "react")
		installFrontend := exec.Command("bun", "install")
		createBackend := exec.Command("mkdir", "backend")
		installExpress := exec.Command("bun", "add", "express")

		_, err := createDir.Output()
		fmt.Println(err)

		os.Chdir(args[0])

		_, err = createVite.Output()
		fmt.Println(err)

		os.Chdir("frontend")

		_, err = installFrontend.Output()
		fmt.Println(err)

		os.Chdir("..")

		_, err = createBackend.Output()
		fmt.Println(err)

		os.Chdir("backend")

		_, err = installExpress.Output()
		fmt.Println(err)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
