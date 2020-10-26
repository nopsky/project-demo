/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/22 16:54
 */
package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nopsky/project-demo/internal/user"
	"github.com/nopsky/project-demo/pkg/util/wrap"
	pb "github.com/nopsky/project-demo/proto/user"
)

type userController struct {
	uuc user.IUserUseCase
}

func NewUserController(uuc user.IUserUseCase) *userController {
	return &userController{uuc: uuc}
}

func (uc *userController) RegisterUser(c *gin.Context) {

	req := new(pb.RegisterUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, wrap.Fail(err))
		return
	}

	resp, err := uc.uuc.RegisterUser(c, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, wrap.Fail(err))
		return
	}

	c.JSON(http.StatusOK, wrap.Success(resp))
}
