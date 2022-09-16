package services

import (
	"basesdedatos/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := models.User{
		Id:        oid,
		Nombre:    "brahyan",
		Apellido:  "jimenez",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := Create(user)
	if err != nil {
		t.Error("La prueba de persistencia de datos del usuario a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}

func TestRead(t *testing.T) {
	users, err := Read()
	if err != nil {
		t.Error("se ha presentado un error en la consulta de usuarios")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("La consulta no retorno datos")
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}

func TestUpdate(t *testing.T) {
	user := models.User{
		Nombre:   "brahyan jose ",
		Apellido: "jimenez echeverri",
	}

	err := Update(user, userId)
	if err != nil {
		t.Error("Error al tratar de actualizar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba de actualizacion finalizo con exito!")
	}
}

func TestDelete(t *testing.T) {
	err := Delete(userId)
	if err != nil {
		t.Error("Error al tratar de eliminar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba de eliminacion finalizo con exito!")
	}

}
