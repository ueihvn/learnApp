package data

type Permission struct {
	Id         int    `db:"id"`
	Type       string `db:"type"`
	Created_At string `db:"created_at"` //timestampz
}

func (permission *Permission) Create() (err error) {
	stmtx, err := Db.Preparex(createPermission)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(permission.Type).StructScan(permission)
	return
}

func GetPermissionById(permissionId int) (permission Permission, err error) {
	stmtx, err := Db.Preparex(getPermissionById)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(permissionId).StructScan(&permission)
	return
}

func GetAllPermission() (permissions []Permission, err error) {
	stmtx, err := Db.Preparex(getAllPermission)
	if err != nil {
		return
	}
	defer stmtx.Close()

	rowxs, err := stmtx.Queryx()
	if err != nil {
		return
	}

	for rowxs.Next() {
		permission := Permission{}
		err = rowxs.StructScan(&permission)
		if err != nil {
			return
		}
		permissions = append(permissions, permission)
	}
	defer rowxs.Close()
	return
}

func (permission *Permission) Delete() (err error) {
	stmtx, err := Db.Preparex(deletePermission)
	if err != nil {
		return
	}
	defer stmtx.Close()

	_, err = stmtx.Exec(permission.Id)
	return
}
