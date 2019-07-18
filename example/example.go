package main

type MyDB struct {
	objID int64
	name string
	email string
}

func (md *MyDB) Contains() []map[string]string {
	return []map[string]string{
		map[string]string{
			"objID": eziql.Quint32,
		},
		map[string]string{
			"name": eziql.Qvarchar,
		},
		map[string]string{
			"email": eziql.Qvarchar,
		},
	}
}

func (md *MyDB) Pack(data []interface{}) {
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
