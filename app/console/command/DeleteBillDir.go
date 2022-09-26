package command

import (
	"fmt"
	"os"
)

func DeleteBillDir() {
	path := "./storage/download/fapp"
	_err := os.RemoveAll(path)
	fmt.Println("error", _err)
}
