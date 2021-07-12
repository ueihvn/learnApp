package data

import "log"

type Course struct {
	Id        int    `db:"id"`
	ClassId   int    `db:"class_id"`
	StartDate string `db:"start_date"` //yyyy-mm-dd
	EndDate   string `db:"end_date"`   //yyyy-mm-dd
	Grade     int    `db:"grade"`
}

func (course *Course) Create() (err error) {
	err = Db.QueryRow("insert into courses (class_id, start_date, end_date, grade) values ($1, $2, $3, $4) returning id",
		course.ClassId,
		course.StartDate,
		course.EndDate,
		course.Grade).Scan(&course.Id)
	return
}

func GetCoursesByGrade(grade int) (courses []Course, err error) {
	rows, err := Db.Queryx("select id, class_id, start_date, end_date, grade from courses where grade = $1", grade)
	if err != nil {
		return
	}

	for rows.Next() {
		course := Course{}
		err = rows.StructScan(&course)
		if err != nil {
			return
		}
		courses = append(courses, course)
	}
	defer rows.Close()
	return
}

func GetCoursesByClassId(classId int) (courses []Course, err error) {
	rows, err := Db.Queryx("select id, class_id, start_date, end_date, grade from courses where grade = $1", classId)
	if err != nil {
		return
	}

	for rows.Next() {
		course := Course{}
		err = rows.StructScan(&course)
		if err != nil {
			return
		}
		courses = append(courses, course)
	}
	defer rows.Close()
	return
}

func GetCourseById(courseId int) (course Course, err error) {
	err = Db.QueryRowx("select id, class_id, start_date, end_date, grade from courses where id = $1", courseId).StructScan(&course)
	return
}

func (course *Course) Update() (err error) {
	_, err = Db.NamedExec(`update courses set class_id= :class_id, start_date= :start_date, end_date= :end_date, grade= :grade where id= :id`, course)
	return
}

func (course *Course) Delete() (err error) {
	_, err = Db.NamedExec(`delete from courses where id=:id`, course)
	return
}

func InitCourse() (err error) {
	mapClass, err := GetAllClassInit()
	if err != nil {
		return
	}

	courses := []Course{
		{
			ClassId:   mapClass[1].Id,
			StartDate: "2021-08-20",
			EndDate:   "2021-08-20",
			Grade:     10,
		},
		{
			ClassId:   mapClass[2].Id,
			StartDate: "2021-08-20",
			EndDate:   "2021-08-20",
			Grade:     10,
		},
		{
			ClassId:   mapClass[3].Id,
			StartDate: "2021-08-20",
			EndDate:   "2021-08-20",
			Grade:     10,
		},
		{
			ClassId:   mapClass[4].Id,
			StartDate: "2021-08-20",
			EndDate:   "2021-08-20",
			Grade:     10,
		},
		{
			ClassId:   mapClass[5].Id,
			StartDate: "2021-08-20",
			EndDate:   "2021-08-20",
			Grade:     10,
		},
		{
			ClassId:   mapClass[1].Id,
			StartDate: "2021-08-20",
			EndDate:   "2021-08-20",
			Grade:     10,
		},
	}

	for _, course := range courses {
		err = course.Create()
		if err != nil {
			return
		}
	}
	return
}

func GetAllCourseInit() (mapCourse map[int]Course, err error) {
	mapCourse = map[int]Course{}
	initCourse := []int{1, 2, 3, 4, 5, 6}

	for _, courseId := range initCourse {
		course, err := GetCourseById(courseId)
		if err != nil {
			log.Fatal(err)
		}
		mapCourse[courseId] = course
	}
	return
}
