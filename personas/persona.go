package personas

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Agente struct {
	Id       string `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type PersonasExecutor interface {
	CreateAgente(agente Agente) error
	VerAgente(id string) (Agente, error)
	VerAgentes() ([]Agente, error)
	EliminarAgente(agente Agente) error
}

type Personas struct {
	database *mongo.Database
}

func NewPersona(database *mongo.Database) PersonasExecutor {
	return &Personas{
		database: database,
	}
}

func (p *Personas) CreateAgente(agente Agente) error {
	_, err := p.database.Collection("agentes").InsertOne(context.Background(), agente)
	if err != nil {
		return err
	}
	return nil
}

func (p *Personas) VerAgente(identificacion string) (Agente, error) {
	ctx := context.Background()
	cursor, err := p.database.Collection("agentes").Find(ctx, bson.M{"id": identificacion})
	if err != nil {
		return Agente{}, err
	}

	var resultado []Agente

	err = cursor.All(ctx, &resultado)
	if err != nil {
		return Agente{}, err
	}

	return resultado[0], nil
}

func (p *Personas) VerAgentes() ([]Agente, error) {

	ctx := context.Background()
	cursor, err := p.database.Collection("agentes").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	resultadoAgentes := []Agente{}

	err = cursor.All(ctx, &resultadoAgentes)
	if err != nil {
		return nil, err
	}

	return resultadoAgentes, nil
}

func (p *Personas) EliminarAgente(agente Agente) error {
	return nil
}
