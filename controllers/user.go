package controllers

import (
	"github.com/RianIhsan/go-rest-api/config"
	"github.com/RianIhsan/go-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

func ReadAll(c *fiber.Ctx) error {
	var user []models.User

	config.DB.Find(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func Read(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")

	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak boleh kosong",
		})
		return nil
	}

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data user tidak ditemukan",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data ditemukan!",
		"user":    user,
	})
}

func Create(c *fiber.Ctx) error {
	user := new(models.UserReq)

	if err := c.BodyParser(user); err != nil {
		return err
	}
	newUser := models.User{
		Nama:     user.Nama,
		Kelas:    user.Kelas,
		Semester: user.Semester,
		Prodi:    user.Prodi,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan data baru",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User berhasil ditambahkan",
		"user":    newUser,
	})
}

func Update(c *fiber.Ctx) error {
	user := new(models.UserReq)

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	id := c.Params("id")

	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak boleh kosong",
		})
		return nil
	}

	updateUser := models.User{
		Nama:     user.Nama,
		Kelas:    user.Kelas,
		Semester: user.Semester,
		Prodi:    user.Prodi,
	}

	if config.DB.Model(&updateUser).Where("id = ?", id).Updates(&updateUser).RowsAffected == 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data user",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil di update",
		"user":    updateUser,
	})

}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if err := config.DB.First(&user, id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahaan !",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Berhasil dihapus",
	})
}
