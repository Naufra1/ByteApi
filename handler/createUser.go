package handler

import (
	"net/http"
	"time"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userInfo struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Email     string
	Address   string
	Birthday  string
	Number    int64
}

// @BasePath /api/v1

// @Summary Create user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /signup [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("erro na validação: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), cost)
	if err != nil {
		logger.Errorf("erro ao encriptar a senha: %v", err)
		return
	}

	users := schemas.User{
		Name:          request.Name,
		Email:         request.Email,
		CriptPassword: string(hash),
		Birthday:      request.Birthday,
		Number:        request.Number,
		Address:       request.Address,
	}

	if err := db.Create(&users).Error; err != nil {
		logger.Errorf("erro adicionando usuários: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}

	userResponse := userInfo{
		Id:        users.ID,
		Name:      users.Name,
		Email:     users.Email,
		Birthday:  users.Birthday,
		Number:    users.Number,
		Address:   users.Address,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
		DeletedAt: users.DeletedAt,
	}
	SendSuccess(ctx, "create-user", userResponse)
}
