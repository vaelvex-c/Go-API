package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/chwlr/golang-api/helper"
	"github.com/chwlr/golang-api/model/domain"
	"time"
)

type UserRepositoryImpl struct {

}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(id_role, name, email, password, created_at) values(?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		user.IdRole,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
	)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update user set id_role = ?, name = ?, email = ?, password = ? where id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.IdRole,
		user.Name,
		user.Email,
		user.Password,
		user.Id,
	)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from user where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error ){
	SQL := "select id, id_role, name, email from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next(){
		err := rows.Scan(&user.Id, &user.IdRole, &user.Name, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	}else {
		return user, errors.New("user no found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, id_role, name, email from user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err  := rows.Scan(&user.Id, &user.IdRole, &user.Name, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}

