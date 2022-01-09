package funcs

import (
	"database/sql"
)

var DB *sql.DB

type User struct {
	ID          int
	FirstName   string
	LastName    string
	NameOfGroup string
	Status      bool
}

func (*User) Get() []User {
	q, err := DB.Query(`SELECT id,firstname,lastname,nameofgroup,status
        FROM firsttable`)
	if err != nil {
		panic(err)
	}

	var users []User
	var user User
	for q.Next() {
		q.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.NameOfGroup,
			&user.Status,
		)
		users = append(users, user)
	}
	return users
}

func (u *User) Update(usr User) bool {
	result, err := DB.Exec(`UPDATE firsttable
		SET firstname = $2, lastname = $3, nameofgroup = $4, status = $5 
		WHERE ID = $1
	`, usr.ID, usr.FirstName, usr.LastName, usr.NameOfGroup, usr.Status)
	if err != nil {
		return false
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return false
	}
	return true
}

func (User) Delete(id int) error {
	result, err := DB.Exec(`delete from firsttable
	where id = $1`, id)
	if err != nil {
		panic(err)
	}
	if i, _ := result.RowsAffected(); i == 0 {
		panic("Delete method was failed")
	}
	return nil
}

func (*User) Insert(usr User) bool {
	result, _ := DB.Exec(`insert into firsttable (id,firstname,lastname,nameofgroup,status)
		values($1,$2,$3,$4,$5)
	`, usr.ID, usr.FirstName, usr.LastName, usr.NameOfGroup, usr.Status)
	if i, _ := result.RowsAffected(); i == 0 {
		panic("Insert method was failed")
	}
	return true
}

func (*User) GetByID(id int) User {
	row := DB.QueryRow(`
	SELECT id,firstname,lastname,nameofgroup,status
        FROM firsttable WHERE id = $1`, id)
	var user User
	row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NameOfGroup,
		&user.Status,
	)
	return user
}
