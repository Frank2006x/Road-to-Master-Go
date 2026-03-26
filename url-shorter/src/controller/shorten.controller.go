package controller

import (
	"fmt"
	"urlshorter/src/db"

	"github.com/gofiber/fiber/v3"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func ShortenURL(c fiber.Ctx) error {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	type body struct {
		URL string `json:"url"`
	}
	var data body
	if err:=c.Bind().Body(&data) ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	id,err:=gonanoid.Generate(alphabet,10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate short URL",
		})
	}
	
	
	_, err = db.DB.Exec(
    "INSERT INTO urls (id, url, short_url) VALUES (?, ?, ?)",
    id, data.URL, id, 
)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save URL",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func GetURL(c fiber.Ctx) error {
	result,err:=db.DB.Query("SELECT * FROM urls ");
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to connect to database",
		})
	}
	defer result.Close()
	type row struct{
		ID string `json:"id"`
		URL string `json:"url"`
		ShortURL string `json:"short_url"`
		Count int `json:"count"`
	}
	var url []row

	for result.Next(){
		var r row
		err := result.Scan(&r.ID, &r.URL, &r.ShortURL, &r.Count)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read data",
			})
		}
		url = append(url, r)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"urls": url,
	})
}

func RedirectURL(c fiber.Ctx) error {
	id := c.Params("id")
	var url string
	err := db.DB.QueryRow("SELECT url FROM urls WHERE short_url = ?", id).Scan(&url)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}
	fmt.Println(url)
	_,err=db.DB.Exec("UPDATE urls set count=count +1 where short_url=?",id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update count",
		})
	}
	return c.Redirect().To(url)

}

