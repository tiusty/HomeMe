package models

/*import (
	"github.com/revel/revel"
)*/

type User struct {
	Id                 int
	Created            int64
	Name               string
	Username, Password string
	HashedPassword     []byte
}
