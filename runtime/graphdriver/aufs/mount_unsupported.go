// +build !linux !amd64
// +build !linux !arm

package aufs

import "errors"

const MsRemount = 0

func mount(source string, target string, fstype string, flags uintptr, data string) (err error) {
	return errors.New("mount is not implemented on darwin")
}
