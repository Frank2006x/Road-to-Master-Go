package controller

import (
	"github.com/Frank2006x/Fibre/src/db"
	"github.com/Frank2006x/Fibre/src/model"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(c fiber.Ctx) error {
	userId,ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	type body struct{
		Title string `json:"title"`
		Description string `json:"description"`
		Status string `json:"status"`
	}
	var data body
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.Status!=string(model.StatusComplete) && data.Status!=string(model.StatusIncomplete)	 {
		data.Status=string(model.StatusIncomplete)
	}

	todo:=bson.M{
		"id": primitive.NewObjectID(),
		"title": data.Title,
		"description": data.Description,
		"status": data.Status,
		"userId": userId,
	}
	_,err:=db.GetCollection("todos").InsertOne(c.Context(),todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
		"todo": todo,
	})
}

func GetTodos(c fiber.Ctx) error {
	userId,ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	var todos []model.Todo
	cursor,err:=db.GetCollection("todos").Find(c.Context(),bson.M{"userId": userId})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch todos",
		})
	}
	if err:= cursor.All(c.Context(),&todos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse todos",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos": todos,
	})
}

func DeleteTodo(c fiber.Ctx) error {
	userId,ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	id:=c.Params("id")

	objId,err:=primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	result,err:=db.GetCollection("todos").DeleteOne(c.Context(),bson.M{"_id": objId,"userId": userId})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete todo",
		})
	}
	if result.DeletedCount==0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todo": result,
	})
}


func UpdateTodo(c fiber.Ctx) error {
	id:=c.Params("id")
	objId,err:=primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}
	userId,ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	update:=bson.M{}
	type body struct{
		Title string `json:"title"`
		Description string `json:"description"`
		Status string `json:"status"`
	}
	var data body
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if data.Title!="" {
		update["title"]=data.Title
	}
	if data.Description!="" {
		update["description"]=data.Description
	}
	if data.Status!="" {
		update["status"]=data.Status
	}
	if len(update)==0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No fields to update",
		})
	}
	result,err:=db.GetCollection("todos").UpdateOne(c.Context(),bson.M{"_id": objId,"userId": userId},bson.M{"$set": update})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update todo",
		})
	}
	if result.MatchedCount==0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todo": result,
	})

}
