package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main/pkg/models"
)

type Repository struct {
	Db     *gorm.DB
	Logger *logrus.Logger
}

func GetRepository(db *gorm.DB, log *logrus.Logger) *Repository {
	return &Repository{Db: db, Logger: log}
}

func (r *Repository) GetCurrentUser(login string) (*models.Employer, error) {
	sqlQuery := `select * from employers where login = ? and active = true`
	var NewUser models.Employer
	err := r.Db.Raw(sqlQuery, login).Scan(&NewUser).Error
	if err != nil {
		return nil, err
	}
	return &NewUser, nil
}

func (r *Repository) InsertNewUser(employer *models.Employer) error {
	sqlQuery := `insert into employers (name, role_id, login, password)
values (?,?,?,?)`
	err := r.Db.Exec(sqlQuery, employer.FullName, employer.RoleId, employer.Login, employer.Password).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertToken(newToken *models.Token) (error, *models.Token) {
	sqlQuery := `insert into tokens (token, employer_id) 
values (?,?) returning *`
	var token models.Token
	err := r.Db.Raw(sqlQuery, newToken.StrToken, newToken.EmployerId).Scan(&token).Error
	if err != nil {
		return err, nil
	}
	return nil, &token
}

func (r *Repository) VerificationLoginAndPass(login, password string) (error, *models.Employer) {
	var user models.Employer
	sqlQuery := `select * from employers where login = ? and active = true`
	err := r.Db.Raw(sqlQuery, login).Scan(&user).Error
	if err != nil {
		return err, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err, nil
	}
	if user.Login == "" {
		return errors.New("You are not in the system, please register "), nil
	}
	return nil, &user
}

func (r *Repository) CheckToken(token *models.Token) (error, int) {
	sqlQuery := `select * from tokens where token = ? and expiration_time > current_timestamp`
	var newToken models.Token
	err := r.Db.Raw(sqlQuery, token.StrToken).Scan(&newToken).Error
	if err != nil {
		return err, 0
	}
	if newToken.StrToken == "" {
		return errors.New("your token has expired"), 0
	}
	return nil, newToken.EmployerId
}

func (r *Repository) ReturnRoleId(userId int) (error, int) {
	sqlQuery := `select role_id from employers where id = ?`
	var roleId int
	err := r.Db.Raw(sqlQuery, userId).Scan(&roleId).Error
	if err != nil {
		return err, 0
	}
	return nil, roleId
}

func (r *Repository) SelectRoles() ([]models.Role, error) {
	sqlQuery := `select * from roles where not(id = 3)`
	roles := make([]models.Role, 0)
	err := r.Db.Raw(sqlQuery).Scan(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
