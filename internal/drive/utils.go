package drive

import (
	"fmt"
	"path/filepath"
)

func (d Drive) fileName(name string) string {
	fmt.Println(filepath.Join(d.Path, name))
	return filepath.Join(d.Path, name)
}
