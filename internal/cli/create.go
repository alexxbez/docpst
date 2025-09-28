package cli

import (
	"fmt"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

func createDirectory(name string) {
	err := os.MkdirAll(name, 0755)
	if err != nil {
		fmt.Printf("Couldn't create directory %v: %v\n", name, err)
		os.Exit(1)
	}
}

func copyTemplate(name string) {
	template_dir := "/home/alexx/bin/docs/" // temporal
	template_source := filepath.Join(template_dir, "template.typ")
	img_source := filepath.Join(template_dir, "img")

	err := cp.Copy(template_source, name+"/main.typ")
	if err != nil {
		fmt.Printf("Couldn't copy template.typ: %v\n", err)
		os.Exit(1)
	}
	err = cp.Copy(img_source, name+"/img")
	if err != nil {
		fmt.Printf("Couldn't copy img/: %v\n", err)
		os.Exit(1)
	}
}

func create(name string) {
	createDirectory(name)
	copyTemplate(name)
	fmt.Printf("Document %v created\n", name)
}
