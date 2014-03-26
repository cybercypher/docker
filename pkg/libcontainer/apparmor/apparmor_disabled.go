// +build !apparmor !linux !amd64 !arm

package apparmor

import ()

func IsEnabled() bool {
	return false
}

func ApplyProfile(pid int, name string) error {
	return nil
}
