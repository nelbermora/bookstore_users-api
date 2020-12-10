package users

import (
	"fmt"

	"github.com/nelbermora/bookstore_users-api/datasources/mysql/users_db"
	"github.com/nelbermora/bookstore_users-api/utils/date_utils"
	"github.com/nelbermora/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created from users where id = ?;"
	queryUpdateUser       = "UPDATE users set  first_name = ?, last_name = ?, email = ? where id = ?;"
	queryDeleteUser       = "DELETE FROM users where id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status from users where status = ?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalErr("Error at prepare Statements")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	err = result.Scan(&user.Id, &user.FirstName, &user.Lastanme, &user.Email, &user.DateCreated)
	if err != nil {
		return errors.NewNotFoundErr("Registro no encontrado")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalErr("Error at prepare Statements")
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertRes, err := stmt.Exec(user.FirstName, user.Lastanme, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalErr(fmt.Sprintf("Error al Crear User: %s", err.Error()))
	}
	userId, err := insertRes.LastInsertId()
	if err != nil {
		return errors.NewInternalErr(fmt.Sprintf("Ultimo Id no encontrado %s", err.Error()))
	}
	user.Id = userId

	return nil

}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalErr("Error at prepare Statements")
	}
	defer stmt.Close()
	_, erro := stmt.Exec(user.FirstName, user.Lastanme, user.Email, user.Id)
	if erro != nil {
		return errors.NewInternalErr("Error al actualziar")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalErr("Error at prepare Statements")
	}
	defer stmt.Close()
	_, errDel := stmt.Exec(user.Id)
	if errDel != nil {
		return errors.NewInternalErr("Error al borrar")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.UsersDB.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalErr("Error at prepare Statements")
	}
	defer stmt.Close()

	rows, errQ := stmt.Query(status)
	if errQ != nil {
		return nil, errors.NewInternalErr("Error at query")
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.FirstName, &user.Lastanme, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			return nil, errors.NewInternalErr("Error on Parsing")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundErr("No encontrados")
	}

	return results, nil

}
