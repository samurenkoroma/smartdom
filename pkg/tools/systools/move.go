package systools

import (
	"fmt"
	"os"
)

func Move(prevPath, newPath string, mode os.FileMode) error {
	err := os.Rename(prevPath, newPath)
	if err != nil {
		byteArr, err2 := os.ReadFile(prevPath)
		if err2 != nil {
			return err2
		}

		err2 = os.WriteFile(newPath, byteArr, mode)
		if err2 == nil {
			// Remove the file iff it was able to be written
			_ = os.Remove(prevPath)
		} else {
			fmt.Printf("unable to write the file out to the new path. Previous: %s --> New: %s\n", prevPath, newPath)
			// Remove any partial file data that may have been written in the case of unfulfilled writes.
			_ = os.Remove(newPath)
		}

		return err2
	}
	return err
}
