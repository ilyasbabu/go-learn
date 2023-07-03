package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilyasbabu/e-com/database"
	"github.com/ilyasbabu/e-com/models"
)

type OrderSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time         `json:"ordered_date"`
}

func SerializeOrder(order models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(201).JSON(err.Error())
	}
	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(201).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := SerializeUser(user)
	responseProduct := SerializeProduct(&product)
	response := SerializeOrder(order, responseUser, responseProduct)
	return c.Status(201).JSON(response)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)

	response := []OrderSerializer{}
	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id=?", order.UserRefer)
		database.Database.Db.Find(&product, "id=?", order.ProductRefer)
		res_order := SerializeOrder(order, SerializeUser(user), SerializeProduct(&product))
		response = append(response, res_order)
	}
	return c.Status(201).JSON(response)

}
