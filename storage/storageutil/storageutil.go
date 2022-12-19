package storageutil

import (
	"fmt"
	"time"
)

// GetSubPath return a formatted representation of the current date
// intended to be used as upload subfolders names
func GetSubPath() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d", t.Year(), t.Month())
}
