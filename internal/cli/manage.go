package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func compileForFirstTime(name string) {
	if _, err := os.Stat(name + ".pdf"); os.IsNotExist(err) {
		cmd := exec.Command("typst", "compile", "main.typ")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error compiling: %v\n", err)
			os.Exit(1)
		}
	}
}

func openZellij() {
	exec.Command("zellij", "action", "new-tab", "-n", "main").Run()
	exec.Command("zellij", "action", "move-tab", "left").Run()
}

func typstWatch(name string) *exec.Cmd {
	typst_cmd := exec.Command("typst", "watch", "main.typ", name+".pdf")
	typst_cmd.Stdout = os.Stdout
	typst_cmd.Stderr = os.Stderr

	if err := typst_cmd.Start(); err != nil {
		fmt.Printf("Error starting typst watch: %v\n", err)
		os.Exit(1)
	}

	return typst_cmd
}

func zathuraOpen(name string) *exec.Cmd {
	zathura_cmd := exec.Command("zathura", name+".pdf")
	zathura_cmd.Stdout = os.Stdout
	zathura_cmd.Stderr = os.Stderr

	if err := zathura_cmd.Start(); err != nil {
		fmt.Printf("Error starting zathura: %v\n", err)
		os.Exit(1)
	}

	return zathura_cmd
}

func closeZellij() {
	exec.Command("zellij", "action", "close-tab").Run()
}

func watch() {
	if _, err := os.Stat("main.typ"); os.IsNotExist(err) {
		fmt.Println("Not in a project directory. Run from project root.")
		os.Exit(1)
	}

	name := "main"
	compileForFirstTime(name)
	openZellij()
	var zathura_cmd *exec.Cmd = zathuraOpen(name)
	var typst_cmd *exec.Cmd = typstWatch(name)

	zathura_cmd.Wait()
	typst_cmd.Process.Kill()
	typst_cmd.Wait()

	closeZellij()

	fmt.Printf("Finished watching file %v", name)
}
