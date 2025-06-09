package Handler

import (
	"auth-api/Database"
	"auth-api/Form"
	"github.com/gofiber/fiber/v2"
)

func ColumnAdd(c *fiber.Ctx) error {
	var reqbody Form.ColumnBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}
 
	if (reqbody.Title != "" &&  reqbody.Description != "" && reqbody.Color != "" &&  reqbody.User != "" ) {
		token, err := Database.CreateColumn(reqbody.Title,reqbody.Description, reqbody.Color, reqbody.User)
		if err != nil {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Column Created successfully",
				"token":   token,
			})
		}
	}
 return nil
}


func ColumnList(c *fiber.Ctx) error {
var reqbody Form.UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	column, err := Database.ListColumn(reqbody.Token)
	if err != nil {
		c.JSON(fiber.Map{
			"status": "error",
		})
		return err
	}

	c.JSON(fiber.Map{
		"status": "OK",
		"data":   column,
	})

	return nil
}

func ColumnDel(c *fiber.Ctx) error {
	var reqbody struct {
		Token string `json:"id"`
	}

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	if reqbody.Token != "" {
		err := Database.DeleteColumn(reqbody.Token)
		if err != nil {
			return c.Status(502).JSON(fiber.Map{
				"status": "ERROR",
				"error":  err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Column deleted successfully",
		})
	}

	return c.Status(400).JSON(fiber.Map{
		"status": "ERROR",
		"error":  "Token is required",
	})
}
