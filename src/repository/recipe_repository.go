package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/dbiagi/gororoba/src/model"
	"log/slog"
)

const (
	RecipeTable = "Recipe"
)

type RecipeRepository struct {
	*dynamodb.DynamoDB
}

func NewRecipeRepository(db *dynamodb.DynamoDB) RecipeRepository {
	return RecipeRepository{DynamoDB: db}
}

func (r *RecipeRepository) GetRecipesByCategory(category string) []domain.Recipe {

	input := dynamodb.QueryInput{
		TableName:              aws.String(RecipeTable),
		KeyConditionExpression: aws.String("category = :category"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":category": {
				S: aws.String(category),
			},
		},
	}

	result, err := r.DynamoDB.Query(&input)

	if err != nil {
		slog.Error("Error querying for recipes by category", slog.String("error", err.Error()))
		return nil
	}

	var recipes []domain.Recipe
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recipes)
	if err != nil {
		slog.Error("Error unmarshalling query result", slog.String("error", err.Error()))
		return nil
	}

	return recipes
}

func (r *RecipeRepository) CreateRecipe(recipe model.RecipeModel) *domain.Error {
	marshalledItem, marshallError := dynamodbattribute.MarshalMap(recipe)

	if marshallError != nil {
		slog.Error("Error marshalling recipe", slog.String("error", marshallError.Error()))
		return &domain.Error{
			Message: "Error marshalling recipe. Message: " + marshallError.Error(),
			Cause:   marshallError,
		}
	}
	_, putError := r.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(RecipeTable),
		Item:      marshalledItem,
	})

	if putError != nil {
		slog.Error("Error putting recipe", slog.String("error", putError.Error()))
		return &domain.Error{
			Message: "Error putting recipe. Message: " + putError.Error(),
			Cause:   putError,
		}
	}

	return nil
}
