package installation

import (
	"debug/buildinfo"
	"errors"
	"os"
)

func RegularTarget(path string) (os.FileInfo, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, errors.New("refusing a symlink or non-regular target")
	}
	return info, nil
}

func Identify(path string) error {
	info, err := buildinfo.ReadFile(path)
	if err != nil || info.Path != "opnay/jev" {
		return errors.New("target is not an identifiable Jev binary; leaving it untouched")
	}
	return nil
}
