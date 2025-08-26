package validation

import (
	"fmt"
	"os"
	"path/filepath"

	"l1/internal/constants"
)

func ValidateFilePath(path string) error {
	if path == "" {
		return fmt.Errorf(constants.ErrorNoPath)
	}

	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return fmt.Errorf(constants.ErrorFileNotExist, path)
	}

	if err != nil {
		return fmt.Errorf(constants.ErrorFileInfo, err)
	}

	if info.IsDir() {
		return fmt.Errorf(constants.ErrorDirectory, path)
	}

	if filepath.Ext(path) != constants.ExtensionTxt {
		return fmt.Errorf(constants.ErrorOnlyTxt, path)
	}

	return nil
}
