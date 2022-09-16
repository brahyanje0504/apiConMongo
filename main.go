package main

import (
	"basesdedatos/mongo"
	personas2 "basesdedatos/personas"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	url := "mongodb+srv://admin:tueliges123@tueligescedula.i6zqf.mongodb.net/tueligescedula?retryWrites=true&w=majority"
	mongoDatabase := mongo.NewDataBase(url, "personas")
	dataBase := mongoDatabase.ConnectionDataBase()
	personas := personas2.NewPersona(dataBase)

	app.Post("/agente", func(c *fiber.Ctx) error {
		agente := personas2.Agente{}
		err := c.BodyParser(&agente)
		if err != nil {
			return c.SendString("error enviando la informacion por favor revise")
		}

		err = personas.CreateAgente(agente)
		if err != nil {
			return c.SendString("error guardando agente")
		}

		return c.SendString("datos guardados correctamente")
	})

	app.Get("/agentes", func(c *fiber.Ctx) error {
		agentes, err := personas.VerAgentes()
		if err != nil {
			return c.SendString("error consultando agentes " + err.Error())
		}

		return c.JSON(agentes)
	})

	app.Get("/agente/:identificacion", func(c *fiber.Ctx) error {
		id := c.Params("identificacion")
		agente, err := personas.VerAgente(id)
		if err != nil {
			return c.SendString("error consultando el agente" + err.Error())
		}

		return c.JSON(agente)

	})

	app.Put("/agente/:identificacion", func(c *fiber.Ctx) error {
		agente := personas2.Agente{}
		err := c.BodyParser(&agente)
		if err != nil {
			return c.SendString("error enviando la informacion por favor revise")
		}

		err = personas.CreateAgente(agente)
		if err != nil {
			return c.SendString("error guardando agente")
		}

		return c.SendString("datos guardados correctamente")
	})

	app.Listen(":3000")

}
