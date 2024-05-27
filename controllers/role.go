package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-authentication/database"
	"github.com/muhrobby/go-authentication/models/entity"
)

func CreateRole(c *fiber.Ctx) error {

	var CreateRole entity.Role
	err := c.BodyParser(&CreateRole)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to get role from body",
			"error":   err.Error(),
		})
	}

	var Role entity.Role

	Role.Name = CreateRole.Name

	err = database.ConnectDB().Create(&Role).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to create role",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Role created successfully",
		"data":    Role,
	})

}

func ShowRole(c *fiber.Ctx) error {

	var Role []entity.Role

	db := database.ConnectDB()

	err := db.Find(&Role).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.ErrBadRequest,
			"message": "failed to find role",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get role",
		"data":    Role,
	})

}

func RoleDestroy(c *fiber.Ctx) error {

	var Role entity.Role

	id := c.Params("id")

	db := database.ConnectDB()

	err := db.Where("id=?", id).First(&Role).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed to get Role",
			"error":   err.Error(),
		})
	}

	err = db.Delete(&Role).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed to Delete Role",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Successfully deleted Role",
	})
}

func RoleUpdate(c *fiber.Ctx) error {

	var Role entity.Role
	id := c.Params("id")

	db := database.ConnectDB()

	err := db.Find(&Role, "id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Role not found",
			"error":   err.Error(),
		})
	}

	var RoleUpdate entity.RoleUpdate

	err = c.BodyParser(&RoleUpdate)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed to body parse role",
			"error":   err.Error(),
		})
	}
	if RoleUpdate.Name == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Tidak Boleh Kosong",
			"error":   err.Error(),
		})
	}

	Role.Name = RoleUpdate.Name

	db.Save(&Role)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Successfully Updating Data",
	})

}
