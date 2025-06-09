package Handler

import (
	"auth-api/Database"
	"strings"
	"auth-api/Form"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var reqbody Form.UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	if isValidEmail(reqbody.Email) && isPasswordValid(reqbody.Password) {
		token, err := Database.Login(reqbody.Email, reqbody.Password)
		if err != nil {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "User registered successfully",
				"token":   token,
			})
		}
	}
 return nil
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func isPasswordValid(password string) bool {
	return len(password) >= 8
}

func Register(c *fiber.Ctx) error {
	var reqbody Form.UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	if isValidEmail(reqbody.Email) && isPasswordValid(reqbody.Password) {
		err, token := Database.Register(reqbody.Email, reqbody.Password, reqbody.Username)
		if err != true {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "User registered successfully",
				"token":   token,
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": "ERROR",
			"error":  "Geçersiz parola yada mail",
		})
	}

	return nil
}

func User(c *fiber.Ctx) error {
	var reqbody Form.UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	user, err := Database.Users(reqbody.Token)
	if err != nil {
		c.JSON(fiber.Map{
			"status": "error",
		})
		return err
	}

	c.JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})

	return nil
}