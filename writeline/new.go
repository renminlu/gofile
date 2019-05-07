package writeline

import (
	"os"
)

func New(pFilename string) (*WriteLine, error) {
	sw, ok := single[pFilename]
	if ok {
		return sw, nil
	}
	// fmt.Println("调试：单列")
	w := WriteLine{}

	fh, err := os.OpenFile(pFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &w, err
	}
	w.osFile = fh
	single[pFilename] = &w
	return &w, nil
}
