package data

type Subject struct {
	Id   int
	Type string
}

func (subject *Subject) Create() (err error) {
	err = Db.QueryRow("insert into subjects (type) values ($1) returning id", subject.Type).Scan(&subject.Id)
	if err != nil {
		return
	}
	return
}

func GetSubjectByType(subject_type string) (subject Subject, err error) {
	subject = Subject{Type: subject_type}
	_, err = Db.Exec("select id, type from subjects where type = $1", subject_type)

	if err != nil {
		return
	}
	return subject, err
}

func GetAllSubject() (subjects []Subject, err error) {
	rows, err := Db.Query("select id, type from subjects")
	if err != nil {
		return
	}
	for rows.Next() {
		subject := Subject{}
		err = rows.Scan(&subject.Id, &subject.Type)
		if err != nil {
			return
		}
		subjects = append(subjects, subject)
	}
	rows.Close()
	return
}

func (subject *Subject) Update() (err error) {
	_, err = Db.Exec("update subjects set type = $2 where id = $1", &subject.Id, &subject.Type)
	return
}

func (subject *Subject) Delete() (err error) {
	_, err = Db.Exec("delete from subjects where id = $1", &subject.Id)
	return
}
