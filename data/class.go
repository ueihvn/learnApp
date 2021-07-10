package data

import "log"

type Class struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	TeacherId int    `db:"teacher_id"`
	IsUse     bool   `db:"is_use"`
}

func (class *Class) Create() (err error) {
	err = Db.QueryRow("insert into class (name, teacher_id) values($1, $2) returning id", class.Name, class.TeacherId).Scan(&class.Id)
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
	_, err = Db.NamedExec(`delete from class where id =:id`, class)
	return
}

func InitClass() (err error) {
	mapTeacher, err := GetAllTeacherInit()
	if err != nil {
		return
	}

	classes := []Class{
		{
			Name:      "class 1 cua thay1",
			TeacherId: mapTeacher["thay1"].Id,
		},
		{
			Name:      "class 1 cua thay2",
			TeacherId: mapTeacher["thay2"].Id,
		},
		{
			Name:      "class 1 cua thay3",
			TeacherId: mapTeacher["thay3"].Id,
		},
		{
			Name:      "class 1 cua thay4",
			TeacherId: mapTeacher["thay4"].Id,
		},
		{
			Name:      "class 1 cua thay5",
			TeacherId: mapTeacher["thay5"].Id,
		},
	}

	for _, class := range classes {
		err = class.Create()
		if err != nil {
			return
		}
	}
	return
}

func GetAllClassInit() (mapClass map[int]Class, err error) {
	mapClass = map[int]Class{}
	classInit := []int{1, 2, 3, 4, 5}

	for _, classId := range classInit {
		class, err := GetClassById(classId)
		if err != nil {
			log.Fatal(err)
		}
		mapClass[classId] = class
	}
	return
}
