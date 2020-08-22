package controller

import (
	"strconv"

	"echo-sample/domain"
	"echo-sample/interface/database"
	"echo-sample/usecase"

	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Save(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Update(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{
		ID: id,
	}
	err = controller.Interactor.DeleteById(user)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}
