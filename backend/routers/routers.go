package routers

import (
	"backend/orm"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var hmacSampleSecret []byte

func Router(e *echo.Echo) {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:N@czaaja2537@mssql_server:1433?database=master"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&orm.UserDB{}, &orm.TodoDB{})

	handler := newHandler(db)

	e.POST("/login", handler.Login)

	todo_authorized := e.Group("/todos", JWTAuthen())
	user_authorized := e.Group("/users", JWTAuthen())

	user_authorized.GET("user", handler.ReadUsersAll)
	todo_authorized.GET("/readall", handler.ReadTodosAll)
	todo_authorized.GET("/readtodo", handler.ReadTodos)
	todo_authorized.POST("/create", handler.CreateTodo)
	// todo_authorized.DELETE("/delete", handler.DeleteTodo)
}

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Login(c echo.Context) error {
	var json orm.User
	if err := c.Bind(&json); err != nil {
		return err
	}
	//check user exists
	var findUser = orm.UserDB{}
	result := h.db.Find(&findUser, "email = ?", json.Email)

	if result.RowsAffected == 0 {
		findUser = orm.UserDB{
			Email: json.Email,
		}
		h.db.Save(&findUser)
	}

	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": findUser.ID,
		"exp":    time.Now().Add(time.Minute * 10).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
		"email": findUser.Email,
	})
}

func JWTAuthen() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			hmacSampleSecret := []byte(os.Getenv("my_secret_key"))

			header := c.Request().Header.Get("Authorization")
			tokenString := strings.Replace(header, "Bearer ", "", 1)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return hmacSampleSecret, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("userId", claims["userId"])
			} else {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": err.Error(),
				})
			}
			return next(c)
		}

	}
}

//-----------------function------------------------

func (h *Handler) ReadUsersAll(c echo.Context) error {
	var users []orm.UserDB
	result := h.db.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c.JSON(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "user read success",
		"users":   users,
	})
}

func (h *Handler) ReadTodosAll(c echo.Context) error {
	var todos = []orm.TodoDB{}
	result := h.db.Find(&todos)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c.JSON(http.StatusBadRequest, result.Error)
	}
	log.Print("ReadTodoAll function is already completed.")
	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	userId := c.Get("userId").(float64)

	todo := orm.Todo{}
	if err := c.Bind(&todo); err != nil {
		fmt.Println(err)
		return err
	}
	newTodo := orm.TodoDB{
		Message:  todo.Message,
		CreateBy: int(userId),
	}
	result := h.db.Save(&newTodo)
	if result.Error != nil {
		fmt.Println(result.RowsAffected)
		return c.JSON(http.StatusBadRequest, result.RowsAffected)
	}
	log.Print("CreateTodo function is already completed.")
	return c.JSON(http.StatusOK, newTodo)
}

func (h *Handler) ReadTodos(c echo.Context) error {
	userId := c.Get("userId").(float64)
	var userIdInt int = int(userId)

	var todos = []orm.TodoDB{}
	result := h.db.Find(&todos, "create_by = ?", userIdInt)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c.JSON(http.StatusBadRequest, result.Error)
	}
	log.Print("ReadTodo function is already completed.")
	return c.JSON(http.StatusOK, todos)
}

// func (h *Handler) DeleteTodo(c echo.Context) error {
// 	log.Print("DeleteTodo function is already completed.")
// }
