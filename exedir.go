package exedir

import (
	"os"
	"path/filepath"
)

func Get() string {
	if os.Getenv("EXEDIR") != "" {
		return os.Getenv("EXEDIR")
	} else {
		p := filepath.Dir(os.Args[0])
		if p == "." {
			var err error
			p, err = os.Getwd()
			if err != nil {
				panic(err)
			}
		}
		return p
	}
}
