package repository

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-product/entity"
	"github.com/MuhammadIbraAlfathar/go-product/helper"
	"log"
)

type UserRepository interface {
	RegisterUser(tx *sql.Tx, req *entity.Users) error
	FindUsers(tx *sql.Tx) ([]entity.Users, error)
	//FindUserByEmail(tx *sql.Tx, email string) (*entity.Users, error)
	//EmailExist(email string) bool
}

type userRepositoryImpl struct {
}

func NewUserRepositoryImpl() *userRepositoryImpl {
	return &userRepositoryImpl{}
}

func (r userRepositoryImpl) RegisterUser(tx *sql.Tx, req *entity.Users) error {

	query := "insert into users(name, user_name, email, password, gender, address) values (?,?,?,?,?,?)"
	exec, err := tx.Exec(query, req.Name, req.UserName, req.Email, req.Password, req.Gender, req.Address)
	if err != nil {
		return &helper.InternalServerError{
			Message: "Internal Server Error",
		}
		log.Println("Error Exec Query")
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return &helper.InternalServerError{
			Message: "Internal Server Error",
		}
		log.Println("Error insert id users")

	}

	req.Id = int(id)

	return nil

}

func (r userRepositoryImpl) FindUsers(tx *sql.Tx) ([]entity.Users, error) {
	query := "select id, name, user_name, email, gender, address from users"
	rows, err := tx.Query(query)
	if err != nil {
		log.Println("Erorr Query get User")
		errorHelper := &helper.InternalServerError{
			Message: "Internal Server Error",
		}
		return nil, errorHelper

	}

	defer rows.Close()

	//if err != nil {
	//	panic(err)
	//}

	var users []entity.Users

	for rows.Next() {
		user := entity.Users{}
		err = rows.Scan(&user.Id, &user.Name, &user.UserName, &user.Email, &user.Gender, &user.Address)
		if err != nil {
			errorHelper := &helper.InternalServerError{
				Message: "Internal Server Error",
			}
			return nil, errorHelper
			log.Println("Erorr Scan data User")
		}

		users = append(users, user)
	}

	return users, nil
}
