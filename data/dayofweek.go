package data

type DayOfWeek struct {
	Id  int
	Day string
}

func (dayOfWeek *DayOfWeek) Create() (err error) {
	err = Db.QueryRow("insert into dayofweek (day) values ($1) returning id", dayOfWeek.Day).Scan(&dayOfWeek.Id)
	return
}

func (dayOfWeek *DayOfWeek) Delete() (err error) {
	_, err = Db.Exec("delete from dayofweek where id = $1", dayOfWeek.Id)
	return
}

func initDayOfWeek() (err error) {
	days := []DayOfWeek{
		{Day: "Monday"},
		{Day: "Tuesday"},
		{Day: "Wednesday"},
		{Day: "Thursday"},
		{Day: "Friday"},
		{Day: "Saturday"},
		{Day: "Sunday"},
	}

	for _, v := range days {
		err = v.Create()
		if err != nil {
			return
		}
	}
	return
}
