package data

type Role struct {
	Id         int    `db:"id"`
	Type       string `db:"type"`
	Created_At string `db:"created_at"` //timestampz
}

func (role *Role) Create() (err error) {
	stmtx, err := Db.Preparex(createRole)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(role.Type).StructScan(role)
	return
}

func GetRoleById(roleId int) (role Role, err error) {
	stmtx, err := Db.Preparex(getRoleById)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(roleId).StructScan(&role)
	return
}

func GetAllRole() (roles []Role, err error) {
	stmtx, err := Db.Preparex(getAllRole)
	if err != nil {
		return
	}
	defer stmtx.Close()

	rowxs, err := stmtx.Queryx()
	if err != nil {
		return
	}

	for rowxs.Next() {
		role := Role{}
		err = rowxs.StructScan(&role)
		if err != nil {
			return
		}
		roles = append(roles, role)
	}
	defer rowxs.Close()
	return
}

func (role *Role) Delete() (err error) {
	stmtx, err := Db.Preparex(deleteRole)
	if err != nil {
		return
	}
	defer stmtx.Close()

	_, err = stmtx.Exec(role.Id)
	return
}
