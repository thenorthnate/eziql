package main

type MyDB struct {
	objID int64
	name  string
	email string
}

func (md *MyDB) Dump() []interface{} {
	return []interface{}{
		md.objID,
		md.name,
		md.email,
	}
}

func (md *MyDB) Pack(data ...interface{}) {
	objIDt, _ := data[0].(int64)
	namet, _ := data[1].(string)
	emailt, _ := data[2].(string)

	md.objID = objIDt
	md.name = namet
	md.email = emailt
}

func main() {
	// so stuff here

}
