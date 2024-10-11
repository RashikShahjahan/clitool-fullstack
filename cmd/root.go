package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "magic",
	Short: "Magic creates you boilerplate fullstack app",
	Long:  "Magic creates you boilerplate fullstack app",
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		var needDB string

		fmt.Print("Please enter project path: ")
		fmt.Scanln(&path)
		createDir := exec.Command("mkdir", path)
		createVite := exec.Command("bun", "create", "vite", "frontend", "--template", "react-swc-ts")
		installFrontend := exec.Command("bun", "install")
		createBackend := exec.Command("mkdir", "backend")
		installExpress := exec.Command("bun", "add", "express")

		_, err := createDir.Output()
		fmt.Println(err)

		os.Chdir(path)

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

		fmt.Println("Do you need a Database? (Y/N)")
		fmt.Scanln(&needDB)

		if strings.ToLower(needDB) == strings.ToLower("Y") {
			var dockerpath string
			installPrisma := exec.Command("bun", "add", "prisma")
			_, err = installPrisma.Output()
			fmt.Println(err)

			prismaInit := exec.Command("npx", "prisma", "init")
			_, err = prismaInit.Output()
			fmt.Print("Please enter docker file path: ")
			fmt.Scanln(&dockerpath)
			dockerPull := exec.Command("docker-compose", "-f", dockerpath, "pull")
			_, err = dockerPull.Output()
			fmt.Println(err)
			dockerUp := exec.Command("docker-compose", "-f", dockerpath, "up", "-d")
			_, err = dockerUp.Output()
			fmt.Println(err)

			newContent := `DATABASE_URL="postgresql://postgres:postgres@localhost:10011"`

			err = os.WriteFile(".env", []byte(newContent), 0644)

		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
