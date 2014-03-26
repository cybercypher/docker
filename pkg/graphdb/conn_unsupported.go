// +build !linux !arm

package graphdb

func NewSqliteConn(root string) (*Database, error) {
	panic("Not implemented")
}
