package data

type User struct {
	Id         int    `db:"id"`
	Email      string `db:"email"`
	FullName   string `db:"full_name"`
	UserName   string `db:"user_name"`
	Password   string `db:"password"`
	Created_At string `db:"created_at"`
	Updated_At string `db:"updated_at"`
}

func (user *User) Create() (err error) {
	stmtx, err := Db.Preparex(createUser)
	if err != nil {
		return err
	}
	defer stmtx.Close()
	err = stmtx.QueryRowx(
		user.Email,
		user.FullName,
		user.UserName,
		user.Password,
	).StructScan(user)
	return
}

func GetUserById(userId int) (user User, err error) {
	stmtx, err := Db.Preparex(getUserById)
	if err != nil {
		return
	}
	defer stmtx.Close()
	err = stmtx.QueryRowx(userId).StructScan(&user)
	return
}

func GetAllUser() (users []User, err error) {
	stmtx, err := Db.Preparex(getAllUser)
	if err != nil {
		return
	}
	defer stmtx.Close()
	rowx, err := stmtx.Queryx(stmtx)

	for rowx.Next() {
		user := User{}
		err = rowx.StructScan(&user)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	defer rowx.Close()
	return
}

func (user *User) Update() (err error) {
	return
}

func (user *User) Delete() (err error) {
	_, err = Db.Exec(deleteUser, user.Id)
	return
}
