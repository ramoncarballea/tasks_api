package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	baseDto "tasks.com/modules/base/dto"
	"tasks.com/modules/identity/domain/services"
	"tasks.com/modules/identity/dto"
)

type UserHandler struct {
	userService services.UserService
	roleService services.RoleService
	log         *zap.Logger
}

// swagger:model CreatedOK
type CreatedOK struct {
	Id any `json:"id"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Status  int
	Code    string
	Message string
}

func NewUserHandler(userService services.UserService, log *zap.Logger, roleService services.RoleService) *UserHandler {
	return &UserHandler{userService: userService, log: log, roleService: roleService}
}

// SignUp registers a new user
// @Summary Register a new user
// @Description Register a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.SignUpDto true "User object that needs to be added"
// @Success 201 {object} CreatedOK
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/user/signup [post]
func (uh *UserHandler) SignUp(c *gin.Context) {
	var payload dto.SignUpDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, baseDto.BadRequest("Invalid payload"))
		return
	}

	uh.log.Info("signing up user", zap.Any("payload", payload))

	userRole, err := uh.roleService.GetDetails("user")
	if err != nil {
		c.JSON(http.StatusBadRequest, baseDto.BadRequest("User role not found"))
		return
	}

	if id, err := uh.userService.Create(dto.CreateUserDto{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
		Roles: []dto.RoleDto{
			{
				ID:   userRole.ID,
				Name: userRole.Name,
			},
		},
	}); err != nil {
		c.JSON(http.StatusInternalServerError, baseDto.ServerError(err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, baseDto.CreatedOK(id))
	}
}
