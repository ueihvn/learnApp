package data

type Resource struct {
	Id         int    `db:"id"`
	Type       string `db:"type"`
	Created_At string `db:"created_at"` //timestampz
}

func (resource *Resource) Create() (err error) {
	stmtx, err := Db.Preparex(createResource)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(resource.Type).StructScan(resource)
	return
}

func GetResourceById(resourceId int) (resource Resource, err error) {
	stmtx, err := Db.Preparex(getResourceById)
	if err != nil {
		return
	}
	defer stmtx.Close()

	err = stmtx.QueryRowx(resourceId).StructScan(&resource)
	return
}

func GetAllResource() (resources []Resource, err error) {
	stmtx, err := Db.Preparex(getAllResource)
	if err != nil {
		return
	}
	defer stmtx.Close()

	rowxs, err := stmtx.Queryx()
	if err != nil {
		return
	}

	for rowxs.Next() {
		resource := Resource{}
		err = rowxs.StructScan(&resource)
		if err != nil {
			return
		}
		resources = append(resources, resource)
	}
	defer rowxs.Close()
	return
}

func (resource *Resource) Delete() (err error) {
	stmtx, err := Db.Preparex(deleteResource)
	if err != nil {
		return
	}
	defer stmtx.Close()

	_, err = stmtx.Exec(resource.Id)
	return
}
