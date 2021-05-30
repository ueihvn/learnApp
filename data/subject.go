package main

type Subject struct {
	Id   int
	Type string
}

func (subject *Subject) Creat() (err error) {
	err = Db.QueryRow("insert into subjects (type) values ($1) returning id", subject.Type).Scan(&subject.Id)
	if err != nil {
		return
	}
	return
}
