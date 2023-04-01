package services

import (
	"context"

	"github.com/szpnygo/VecTextSearch/config"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func addVector(client *weaviate.Client, appConfig *config.AppConfig, id string, dataSchema map[string]interface{}, vector []float32) error {
	_, err := client.Data().Creator().
		WithClassName(appConfig.WeaviateClassName).
		WithID(id).
		WithProperties(dataSchema).
		WithVector(vector).
		Do(context.Background())

	return err
}

func searchVectors(client *weaviate.Client, appConfig *config.AppConfig, vector []float32) (*models.GraphQLResponse, error) {
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
		WithClassName(appConfig.WeaviateClassName).
		WithFields(name, content, _additional).
		WithNearVector(nearVector).
		Do(ctx)

	return result, err
}

func findTextByContent(client *weaviate.Client, appConfig *config.AppConfig, content string) (*models.GraphQLResponse, error) {
	// 创建 where 过滤器
	whereFilter := filters.Where().
		WithPath([]string{"content"}).
		WithOperator(filters.Equal).
		WithValueText(content)

	// 定义查询所需的字段
	fields := []graphql.Field{
		{Name: "name"},
		{Name: "content"},
	}

	// 构建 GraphQL 查询
	queryBuilder := client.GraphQL().Get().
		WithClassName(appConfig.WeaviateClassName).
		WithFields(fields...).
		WithWhere(whereFilter)

	// 执行查询并获取结果
	ctx := context.Background()
	result, err := queryBuilder.Do(ctx)

	return result, err
}
