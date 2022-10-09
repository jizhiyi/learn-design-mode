package sample

var personMgr *PersonMgr

type PersonMgr struct {
	PlayerCount int
}

func NewPersonMgr() *PersonMgr {
	if personMgr == nil {
		personMgr = &PersonMgr{}
	}
	return personMgr
}
