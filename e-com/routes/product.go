package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ilyasbabu/e-com/database"
	"github.com/ilyasbabu/e-com/models"
)

type ProductSerializer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	SerialNo string `json:"serialNo"`
}

func SerializeProduct(productModel *models.Product) ProductSerializer {
	return ProductSerializer{ID: productModel.ID, Name: productModel.Name, SerialNo: productModel.SerialNo}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.Database.Db.Create(&product)
	response := SerializeProduct(&product)
	return c.Status(200).JSON(response)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	response := []ProductSerializer{}
	for _, product := range products {
		serializedPrdouct := SerializeProduct(&product)
		response = append(response, serializedPrdouct)
	}
	return c.Status(200).JSON(response)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("invalid ID")
	}
	return nil
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}
	if err := findProduct(id, &product); err != nil {
		return c.SendStatus(404)
	}
	// type updateProductSerializer struct {
	// 	Name     string `json:"name"`
	// 	SerialNo string `json:"serialNo"`
	// }
	// var updateData updateProductSerializer
	// if err := c.BodyParser(&updateData); err != nil {
	// 	return c.Status(400).SendString(err.Error())
	// }
	// user.Name = updateData.Name
	// user.SerialNo = updateData.SerialNo
	name := c.FormValue("name")
	serialNo := c.FormValue("serialNo")
	product.Name = name
	product.SerialNo = serialNo
	database.Database.Db.Save(&product)
	response := SerializeProduct(&product)
	return c.Status(200).JSON(response)
}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}
	if err := findProduct(id, &product); err != nil {
		return c.SendStatus(404)
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("Deleted succesfully")
}
