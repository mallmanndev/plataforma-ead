package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/value-objects"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (ur *UsersRepository) Create(user *entities.User) error {
	fmt.Println("Chegou aqui!")
	_, err := ur.db.Exec("INSERT INTO users (id, name, email, phone, password, type_id) VALUES ($1, $2, $3, $4, $5, $6);",
		user.Id, user.Name, user.Email.Email, user.Phone.Phone, user.GetPassword(), user.Type.Id)

	return err
}

func (ur *UsersRepository) Update(user *entities.User) error {
	_, err := ur.db.Exec("UPDATE users SET name = $1, email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4;",
		user.Name, user.Email, user.Phone, user.Id)

	return err
}

func (ur *UsersRepository) Delete(id string) error {
	_, err := ur.db.Exec("UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1;", id)

	return err
}

func (ur *UsersRepository) FindByEmail(email *value_objects.EmailAddress) (*entities.User, error) {
	row := ur.db.QueryRow("SELECT * FROM users WHERE deleted_at IS NULL AND email = $1;",
		email.Email)

	model := &UserModel{}
	err := row.Scan(&model.Id, &model.Name, &model.Email,
		&model.Phone, &model.Password, &model.Type,
		&model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return modelToEntity(model), nil
}

func (ur *UsersRepository) FindById(id string) (*entities.User, error) {
	row := ur.db.QueryRow("SELECT * FROM users WHERE deleted_at IS NULL AND id = $1;", id)

	if row == nil {
		return nil, nil
	}

	model := &UserModel{}
	err := row.Scan(&model.Id, &model.Name, &model.Password,
		&model.Phone, &model.Password, &model.Type,
		&model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return modelToEntity(model), nil
}

type UserModel struct {
	Id        string
	Name      string
	Email     string
	Phone     string
	Password  string
	Type      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

func modelToEntity(m *UserModel) *entities.User {
	userType := &entities.UserType{Id: m.Type}
	email := &value_objects.EmailAddress{Email: m.Email}
	phone := &value_objects.Phone{Phone: m.Phone}
	user := &entities.User{
		Id:        m.Id,
		Name:      m.Name,
		Email:     email,
		Phone:     phone,
		Type:      userType,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt.Time,
	}

	user.SetPassword(m.Password)

	return user
}
