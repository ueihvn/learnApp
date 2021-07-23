package data

type Course struct {
	Id          int    `db:"id"`
	Resource_id int    `db:"resource_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Subject_id  string `db:"subject_id"`
	IsDelete    bool   `db:"is_delete"`
	IsOpen      bool   `db:"is_open"`
}

func (course *Course) Create() (err error) {
	stmtx, err := Db.Preparex(createCourse)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(
		course.Resource_id,
		course.Name,
		course.Description,
		course.Subject_id,
	).StructScan(course)
	return
}

/*
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
*/
