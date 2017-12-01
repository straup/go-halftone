package functions

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
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

	fh, err := os.Open(info)

	if err != nil {
		return true, err
	}

	defer fh.Close()

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		return true, err
	}

	var rsp gjson.Result

	rsp = gjson.GetBytes(body, "refers_to_uid")

	if !rsp.Exists() {
		return true, errors.New("Unable to determine refers_to_uid")
	}

	uid := rsp.Int()

	object_fname := fmt.Sprintf("%d.json", uid)
	object_info := filepath.Join(root, object_fname)

	_, err = os.Stat(object_info)

	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return true, err
	}

	object_fh, err := os.Open(object_info)

	if err != nil {
		return true, err
	}

	defer object_fh.Close()

	object_body, err := ioutil.ReadAll(object_fh)

	if err != nil {
		return true, err
	}

	fmt.Println(len(object_body))
	// get refers_to_uid
	// read refers_to_uid.json
	// check for is_primary

	return true, nil
}
