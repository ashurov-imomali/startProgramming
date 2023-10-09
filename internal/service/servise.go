package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"main/internal/repository"
	"main/pkg/models"
	"time"
	"unicode"
)

type Service struct {
	Repository *repository.Repository
	Logger     *logrus.Logger
}

func GetService(repos *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{Repository: repos, Logger: logger}
}

func (srv *Service) CheckNewUser(login, password string) error {
	currentUser, err := srv.Repository.GetCurrentUser(login)
	if err != nil {
		return err
	}
	if currentUser.Login != "" {
		return errors.New("this login is used:(")
	}
	if len([]rune(login)) < 5 {
		return errors.New("this login is too short")
	}
	err = CheckPassword(password)
	if err != nil {
		return err
	}
	return nil
}
func (srv *Service) CreateNewUser(employer *models.Employer) error {
	hash, err := GetHash(employer.Password)
	if err != nil {
		return err
	}
	employer.Password = hash
	err = srv.Repository.InsertNewUser(employer)
	if err != nil {
		return err
	}
	return nil
}
func GetHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(password string) error {
	if len([]rune(password)) < 8 {
		return errors.New("Пароль должен содержать не менее 8 символов")
	}

	hasNumber := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return errors.New("Пароль должен содержать хотя бы одну цифру")
	}

	return nil
}

func (srv *Service) AddNewToken(token *models.Token) (error, *models.Token) {
	err, token := srv.Repository.InsertToken(token)
	if err != nil {
		return err, nil
	}
	return nil, token
}

func (srv *Service) CheckUser(login, password string) (error, *models.Employer) {
	err, user := srv.Repository.VerificationLoginAndPass(login, password)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (srv *Service) GetUser(token *models.Token) (error, *models.Employer) {
	err, userId := srv.Repository.CheckToken(token)
	if err != nil {
		return err, nil
	}
	useFullToken, err := jwt.Parse(token.StrToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return err, nil
	}
	claims, ok := useFullToken.Claims.(jwt.MapClaims) //todo
	if ok && useFullToken.Valid {
		roleId := claims["roleId"].(float64)
		user := models.Employer{Id: userId, RoleId: int(roleId)}
		return nil, &user
	}
	return errors.New("failed to retrieve information"), nil
}

func (srv *Service) GetToken(user *models.Employer) (*models.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	calms := token.Claims.(jwt.MapClaims)
	calms["userId"] = user.Id
	calms["roleId"] = user.RoleId
	calms["time"] = time.Now()
	secretKey := []byte("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	var NewToken models.Token
	NewToken.StrToken = tokenString
	NewToken.EmployerId = user.Id
	return &NewToken, nil
}

func (srv *Service) ReturnRoles() ([]models.Role, error) {
	roles, err := srv.Repository.SelectRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
