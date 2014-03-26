// +build !linux !amd64
// +build !linux !arm

package lxc

func setHostname(hostname string) error {
	panic("Not supported on darwin")
}
