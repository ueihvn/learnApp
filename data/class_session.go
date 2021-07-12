package data

type ClassSession struct {
	Id          int    `db:"id"`
	ClassId     int    `db:"class_id"`
	DayOfWeekId int    `db:"dayofweek_id"`
	StartTime   string `db:"start_time"` //hh:mm:ss
	EndTime     string `db:"end_time"`   //hh:mm:ss
	IsUse       bool   `db:"is_use"`
}

func (classSession *ClassSession) Create() (err error) {
	err = Db.QueryRow("insert into class_session (class_id, dayofweek_id, start_time, end_time) values ($1, $2, $3, $4) returning id",
		classSession.ClassId,
		classSession.DayOfWeekId,
		classSession.StartTime,
		classSession.EndTime).Scan(&classSession.Id)
	return
}

func GetClassSessionByClassId(classId int) (classSessions []ClassSession, err error) {
	rows, err := Db.Queryx("select id, class_id, dayofweek_id, is_use start_time, end_time from class_session where id = $1", classId)
	if err != nil {
		return
	}

	for rows.Next() {
		classSession := ClassSession{}
		err = rows.StructScan(&classSession)
		if err != nil {
			return
		}
		classSessions = append(classSessions, classSession)
	}
	defer rows.Close()
	return
}

func GetClassSessionByDayOfWeekId(dayOfWeekId int) (classSessions []ClassSession, err error) {
	rows, err := Db.Queryx("select id, class_id, dayofweek_id, is_use, start_time, end_time from class_session where id = $1", dayOfWeekId)
	if err != nil {
		return
	}

	for rows.Next() {
		classSession := ClassSession{}
		err = rows.StructScan(&classSession)
		if err != nil {
			return
		}
		classSessions = append(classSessions, classSession)
	}
	defer rows.Close()
	return
}

func (classSession *ClassSession) Update() (err error) {
	_, err = Db.NamedExec(`update class_session set class_id= :class_id, start_time= :start_time, end_time= :end_time, is_use= :is_use where id = :id`, classSession)
	return
}

func (classSession *ClassSession) Delete() (err error) {
	_, err = Db.NamedExec(`delete from class_session where id= :id`, classSession)
	return
}
