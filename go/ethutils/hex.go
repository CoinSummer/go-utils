package ethutils

import (
	"fmt"
)

func HexZeroPad(value string, length int) (string, error) {
	if len(value) > 2*length+2 || len(value) <= 2 {
		return "", fmt.Errorf("%s value out of range", value)
	}

	for {
		if len(value) < 2*length+2 {
			value = "0x0" + value[2:]
		} else {
			return value, nil
		}
	}
}
