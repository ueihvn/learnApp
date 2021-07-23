package data

type Class struct {
	Id        int    `db:"id"`
	CourseId  int    `db:"course_id"`
	StartDate string `db:"start_date"` //yyyy-mm-dd
	EndDate   string `db:"end_date"`   //yyyy-mm-dd
	Grade     int    `db:"grade"`
}

func (class *Class) Create() (err error) {
	err = Db.QueryRow("insert into classes (course_id, start_date, end_date, grade) values ($1, $2, $3, $4) returning id",
		class.CourseId,
		class.StartDate,
		class.EndDate,
		class.Grade).Scan(&class.Id)
	return
}

func GetClassesByGrade(grade int) (classes []Class, err error) {
	rows, err := Db.Queryx("select id, course_id, start_date, end_date, grade from classes where grade = $1", grade)
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
	defer rows.Close()
	return
}

func GetClassesByClassId(classId int) (classes []Class, err error) {
	rows, err := Db.Queryx("select id, class_id, start_date, end_date, grade from classes where grade = $1", classId)
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
	defer rows.Close()
	defer rows.Close()
	return
}

func GetClassById(classId int) (class Class, err error) {
	err = Db.QueryRowx("select id, course_id, start_date, end_date, grade from classes where id = $1", classId).StructScan(&class)
	return
}

func (class *Class) Update() (err error) {
	_, err = Db.NamedExec(`update classes set course_id= :course_id, start_date= :start_date, end_date= :end_date, grade= :grade where id= :id`, class)
	return
}

func (class *Class) Delete() (err error) {
	_, err = Db.NamedExec(`delete from classes where id=:id`, class)
	return
}

// func InitCourse() (err error) {
// 	mapClass, err := GetAllClassInit()
// 	if err != nil {
// 		return
// 	}

// 	class := []Course{
// 		{
// 			ClassId:   mapClass[1].Id,
// 			StartDate: "2021-08-20",
// 			EndDate:   "2021-08-20",
// 			Grade:     10,
// 		},
// 		{
// 			ClassId:   mapClass[2].Id,
// 			StartDate: "2021-08-20",
// 			EndDate:   "2021-08-20",
// 			Grade:     10,
// 		},
// 		{
// 			ClassId:   mapClass[3].Id,
// 			StartDate: "2021-08-20",
// 			EndDate:   "2021-08-20",
// 			Grade:     10,
// 		},
// 		{
// 			ClassId:   mapClass[4].Id,
// 			StartDate: "2021-08-20",
// 			EndDate:   "2021-08-20",
// 			Grade:     10,
// 		},
// 		{
// 			ClassId:   mapClass[5].Id,
// 			StartDate: "2021-08-20",
// 			EndDate:   "2021-08-20",
// 			Grade:     10,
// 		},
// 		{
// 			ClassId:   mapClass[1].Id,
// 			StartDate: "2021-08-20",
// 			EndDate:   "2021-08-20",
// 			Grade:     10,
// 		},
// 	}

// 	for _, class := range courses {
// 		err = class.Create()
// 		if err != nil {
// 			return
// 		}
// 	}
// 	return
// }

// func GetAllCourseInit() (mapCourse map[int]Course, err error) {
// 	mapCourse = map[int]Course{}
// 	initCourse := []int{1, 2, 3, 4, 5, 6}

// 	for _, classId := range initCourse {
// 		class, err := GetCourseById(courseId)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		mapCourse[classId] = course
// 	}
// 	return
// }
