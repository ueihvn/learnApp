package data

type Class struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	TeacherId int    `db:"teacher_id"`
	IsUse     bool   `db:"is_use"`
}

func (class *Class) Create() (err error) {
	err = Db.QueryRow("insert into class (name, teacher_id, is_use) values($1, $2, $3) returning id").Scan(&class.Id)
	return
}

func GetAllClassByTeacherId(teacherId int) (classes []Class, err error) {
	rows, err := Db.Queryx("select id, name, teacher_id, is_use from class where teacher_id = $1", teacherId)
	if err != nil {
		return
	}
	for rows.Next() {
		class := Class{}
		err = rows.StructScan(&class)
		if err != nil {
			return
		}
		classes = append(classes, class)
	}
	rows.Close()
	return
}

func GetClassById(classId int) (class Class, err error) {
	err = Db.QueryRowx("select id, name, teacher_id, is_use from class where id =$1", classId).StructScan(&class)
	return
}

func (class *Class) Update() (err error) {
	_, err = Db.NamedExec(`update class set name= :name, teacher_id= :teacher_id, is_use= :is_use`, class)
	return
}

func (class *Class) Delete() (err error) {
	_, err = Db.NamedExec(`delete class where id =:id`, class)
	return
}
