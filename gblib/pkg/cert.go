package gblib

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func GenerateCert() {
	cert := [2]string{"certificate-key.crt", "certificate.crt"}

	if _, err := os.Stat("mkcert.exe"); os.IsNotExist(err) {
		fmt.Println("mkcert missing.")
		os.Exit(1)
	} else {
		cmd := exec.Command("mkcert.exe", "0.0.0.0")
		if err := cmd.Run(); err != nil {
			fmt.Println("error: ", err.Error())
			os.Exit(1)
		}
	}

	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if path == "0.0.0.0-key.pem" {
			return os.Rename(path, cert[0])
		}

		if path == "0.0.0.0.pem" {
			return os.Rename(path, cert[1])
		}
		return nil
	})

}
