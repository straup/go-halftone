package functions

import (
	"os"
	"path/filepath"
	"strings"
)

func DefaultFilterFunc(string) (bool, error) {
	return true, nil
}

func CooperHewittShoeboxFilterFunc(path string) (bool, error) {

	if !strings.HasSuffix(path, "_b.jpg") {
		return false, nil
	}

	root := filepath.Dir(path)
	info := filepath.Join(root, "index.json")

	_, err := os.Stat(info)

	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return true, err
	}

	// get refers_to_uid
	// read refers_to_uid.json
	// check for is_primary

	return true, nil
}
