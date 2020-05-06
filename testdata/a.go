package testdata

import (
	"errors"
	"fmt"
)

func _() {
	var b1, b2, b3, b4 bool
	if b1 { // complexity: 9
		if b2 { // +1
			if b3 { // +2
			}
		}

		if b2 { // +1
			if b3 { // +2
				if b4 { // +3
				}
			}
		}
	}
}

func _() error {
	msg := "msg"
	return errors.New(fmt.Sprintf("error: %s", msg))
}
