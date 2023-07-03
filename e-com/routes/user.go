package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ilyasbabu/e-com/database"
	"github.com/ilyasbabu/e-com/models"
)

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func SerializeUser(userModel models.User) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := SerializeUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}
	for _, user := range users {
		responseUser := SerializeUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.First(&user, id)
	if user.ID == 0 {
		return errors.New("invalid user ID")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := SerializeUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type updateUserSerializer struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var updateData updateUserSerializer
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)
	responseUser := SerializeUser(user)
	type Response struct {
		Msg  string         `json:"msg"`
		Data UserSerializer `json:"data"`
	}
	var response Response
	response.Msg = "Updated Successfully!"
	response.Data = responseUser
	return c.Status(200).JSON(response)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("Deleted succesfully")
}
