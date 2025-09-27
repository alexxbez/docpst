package cli

import (
	"fmt"
	cp "github.com/otiai10/copy"
	// "io"
	"os"
	"path/filepath"
)

func create_directory(name string) {
	err := os.MkdirAll(name, 0755)
	if err != nil {
		fmt.Errorf("Couldn't create directory %v: %v", name, err)
		os.Exit(1)
	}
}

func copy_template() {
	template_dir := "/home/alexx/bin/docs/" // temporal
	template_source := filepath.Join(template_dir, "template.typ")
	img_source := filepath.Join(template_dir, "img")

	err := cp.Copy(template_source, "./main.typ")
	if err != nil {
		fmt.Errorf("Couldn't copy template.typ: %v", err)
	}
	err = cp.Copy(img_source, ".")
	if err != nil {
		fmt.Errorf("Couldn't copy img/: %v", err)
	}
}

func create(name string) {
	create_directory(name)
	copy_template()
}
