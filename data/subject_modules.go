package data

type SubjectModule struct {
	Id        int    `db:"id"`
	SubjectId int    `db:"subject_id"`
	Type      string `db:"type"`
}

func (subjectModule *SubjectModule) Create() (err error) {
	err = Db.QueryRow("insert into subject_modules (subject_id, type) values ($1, $2) returning id",
		subjectModule.SubjectId,
		subjectModule.Type).Scan(&subjectModule.Id)
	return
}

func GetSubjectModulesBySubjectId(subjectId int) (subjectModules []SubjectModule, err error) {
	rows, err := Db.Queryx("select id, subject_id, type from subject_modules where subject_id = $1", subjectId)

	if err != nil {
		return
	}

	for rows.Next() {
		subjectModule := SubjectModule{}
		err = rows.StructScan(&subjectModule)
		if err != nil {
			return
		}
		subjectModules = append(subjectModules, subjectModule)
	}
	rows.Close()
	return
}

func GetSubjectModuleById(subjectModuleId int) (subjectModule SubjectModule, err error) {
	err = Db.QueryRowx("select id, subject_id, type from subject_modules where id= $1", &subjectModule.Id).StructScan(&subjectModule)
	return
}

func (subjectModule *SubjectModule) Update() (err error) {
	_, err = Db.NamedExec(`update subject_modules set subject_id=:subject_id, type=:type where id = :id`, subjectModule)
	return
}

func (subjectModule *SubjectModule) Delete() (err error) {
	_, err = Db.NamedExec(`delete from subject_modules where id = :id`, subjectModule)
	return
}

func InitSubjectModules() (err error) {
	subjectInit, err := GetAllSubjectInit()
	if err != nil {
		return err
	}

	subjectModules := []SubjectModule{
		{
			SubjectId: subjectInit["Toán Học"].Id,
			Type:      "Lượng Giác",
		},
		{
			SubjectId: subjectInit["Toán Học"].Id,
			Type:      "Xác Xuất",
		},
		{
			SubjectId: subjectInit["Vật Lý"].Id,
			Type:      "Chuyển động cơ",
		},
		{
			SubjectId: subjectInit["Vật Lý"].Id,
			Type:      "Điện xoay chiều",
		},
		{
			SubjectId: subjectInit["Hóa Học"].Id,
			Type:      "Hóa vô cơ",
		},
		{
			SubjectId: subjectInit["Hóa Học"].Id,
			Type:      "Hóa hữu cơ",
		},
		{
			SubjectId: subjectInit["Sinh Học"].Id,
			Type:      "Di truyền",
		},
		{
			SubjectId: subjectInit["Sinh Học"].Id,
			Type:      "ADN",
		},
		{
			SubjectId: subjectInit["Tiếng Anh"].Id,
			Type:      "Idiom",
		},
		{
			SubjectId: subjectInit["Tiếng Anh"].Id,
			Type:      "Pronunciation",
		},
	}

	for _, sm := range subjectModules {
		err = sm.Create()
		if err != nil {
			return err
		}
	}
	return
}
