package validation

import (
	"fmt"
	"os"
	"path/filepath"

	"l1/internal/constants"
)

func ValidateFilePath(path string) error {
	if path == "" {
		return fmt.Errorf(constants.NO_PATH_GIVEN)
	}

	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return fmt.Errorf(constants.FILE_DOES_NOT_EXIST, path)
	}

	if err != nil {
		return fmt.Errorf(constants.FILE_INFO_ERROR, err)
	}

	if info.IsDir() {
		return fmt.Errorf(constants.DIRECTORY_ERROR, path)
	}

	if filepath.Ext(path) != constants.TXT {
		return fmt.Errorf(constants.ONLY_TXT, path)
	}

	return nil
}
