package Utils

import (
	"github.com/rs/xid"
)

func RandomId() string {
	id := xid.New().String()
	return id
}
