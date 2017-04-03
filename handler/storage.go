package handler

import (
	"fmt"
	"path/filepath"
)

type Storage struct {
	root string
}

func (s *Storage) CreateStorage(root string) {
	abs_root, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Storage create failed")
	}
	s.root = abs_root
	fmt.Println("root storage: ", s.root)
}
