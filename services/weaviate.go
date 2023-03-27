package services

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func addVector(client *weaviate.Client, id string, dataSchema map[string]interface{}, vector []float32) error {
	_, err := client.Data().Creator().
		WithClassName("Text").
		WithID(id).
		WithProperties(dataSchema).
		WithVector(vector).
		Do(context.Background())

	return err
}

func searchVectors(client *weaviate.Client, vector []float32) (*models.GraphQLResponse, error) {
	className := "Text"
	name := graphql.Field{Name: "name"}
	content := graphql.Field{Name: "content"}
	_additional := graphql.Field{
		Name: "_additional",
		Fields: []graphql.Field{
			{Name: "certainty"},
			{Name: "distance"},
		},
	}

	nearVector := client.GraphQL().NearVectorArgBuilder().
		WithVector(vector)

	ctx := context.Background()
	result, err := client.GraphQL().Get().
		WithClassName(className).
		WithFields(name, content, _additional).
		WithNearVector(nearVector).
		Do(ctx)

	return result, err
}
