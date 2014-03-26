// +build !linux !amd64
// +build !linux !arm

package utils

import (
	"errors"
)

type Utsname struct {
	Release [65]byte
}

func uname() (*Utsname, error) {
	return nil, errors.New("Kernel version detection is available only on linux")
}
