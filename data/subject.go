package data

import "log"

type Subject struct {
	Id   int    `db:"id"`
	Type string `db:"type"`
}

func (subject *Subject) Create() (err error) {
	err = Db.QueryRow("insert into subjects (type) values ($1) returning id", subject.Type).Scan(&subject.Id)
	return
}

func GetSubjectById(subjectId int) (subject Subject, err error) {
	err = Db.QueryRow("select id from subjects where id = $1", subjectId).Scan(&subject.Type)
	return
}

func GetSubjectByType(subjectType string) (subject Subject, err error) {
	err = Db.QueryRowx("select id, type from subjects where type = $1", subjectType).StructScan(&subject)
	return
}

func GetAllSubject() (subjects []Subject, err error) {
	rows, err := Db.Queryx("select id, type from subjects")
	if err != nil {
		return
	}
	for rows.Next() {
		subject := Subject{}
		err = rows.StructScan(&subject)
		if err != nil {
			return
		}
		subjects = append(subjects, subject)
	}
	rows.Close()
	return
}

func (subject *Subject) Update() (err error) {
	_, err = Db.Exec(`update subjects set type = $2 where id = $1`, &subject.Id, &subject.Type)
	return
}

func (subject *Subject) Delete() (err error) {
	_, err = Db.Exec("delete from subjects where id = $1", &subject.Id)
	return
}

func InitSubject() (err error) {
	subjects := []Subject{
		{Type: "Toán Học"},
		{Type: "Vật Lý"},
		{Type: "Hóa Học"},
		{Type: "Sinh Học"},
		{Type: "Tiếng Anh"},
	}

	for _, subject := range subjects {
		err = subject.Create()
		if err != nil {
			return
		}

	}

	return
}

func GetAllSubjectInit() (mapSubject map[string]Subject, err error) {
	mapSubject = map[string]Subject{}
	subjectInit := []string{"Toán Học", "Vật Lý", "Hóa Học", "Sinh Học", "Tiếng Anh"}

	for _, subjectType := range subjectInit {
		subject, err := GetSubjectByType(subjectType)
		if err != nil {
			log.Fatal(err)
		}
		mapSubject[subjectType] = subject
	}

	return
}
