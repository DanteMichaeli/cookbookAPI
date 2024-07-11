package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DanteMichaeli/CookBookAPI/graph/model"
)

// CreateRecipe is the resolver for the createRecipe field. Creates a recipe from the given input, stores it in the recipes directory as a JSON file.
func (r *mutationResolver) CreateRecipe(ctx context.Context, id string, title string, description string, ingredients []string, steps []string) (*model.Response, error) {
	// check if id is empty
	err := idCheck(id)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to create recipe."}, err
	}

	// instantiate recipe
	recipe := model.Recipe{
		ID:          id,
		Title:       title,
		Description: description,
		Ingredients: ingredients,
		Steps:       steps,
	}

	// encode to JSON
	recipeJSON, err := encodeRecipe(&recipe)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to create recipe.", Recipe: nil}, err
	}

	// Write recipe to directory
	err = writeToDir(id, recipeJSON)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to create recipe.", Recipe: nil}, err
	}

	return &model.Response{Success: true, Message: "Recipe created successfully.", Recipe: &recipe}, nil
}

// UpdateRecipe is the resolver for the updateRecipe field. Updates data of an existing recipe (ID immutable)
func (r *mutationResolver) UpdateRecipe(ctx context.Context, id string, title *string, description *string, ingredients []string, steps []string) (*model.Response, error) {
	// check if id is empty
	err := idCheck(id)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to update recipe.", Recipe: nil}, err
	}

	// find and decode recipe JSON file
	recipePtr, err := decodeRecipe(id)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to update recipe.", Recipe: nil}, err
	}

	// update recipe with provided fields ID
	if title != nil {
		recipePtr.Title = *title
	}
	if description != nil {
		recipePtr.Description = *description
	}
	if ingredients != nil {
		recipePtr.Ingredients = ingredients
	}
	if steps != nil {
		recipePtr.Steps = steps
	}

	// encode updated recipe to JSON
	recipeJSON, err := encodeRecipe(recipePtr)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to update recipe.", Recipe: nil}, err
	}

	// write updated recipe to directory
	err = writeToDir(id, recipeJSON)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to update recipe.", Recipe: nil}, err
	}

	return &model.Response{Success: true, Message: "Recipe updated successfully.", Recipe: recipePtr}, nil
}

// DeleteRecipe is the resolver for the deleteRecipe field.
func (r *mutationResolver) DeleteRecipe(ctx context.Context, id string) (*model.Response, error) {
	// check if id is empty
	err := idCheck(id)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to delete recipe.", Recipe: nil}, err
	}

	// delete JSON file with corresponding ID
	fileName := fmt.Sprintf("%s.json", id)
	filePath := filepath.Join(recipesDir, fileName)
	err = os.Remove(filePath)
	if err != nil {
		return &model.Response{Success: false, Message: "Failed to delete recipe.", Recipe: nil}, err
	}

	return &model.Response{Success: true, Message: "Recipe deleted successfully.", Recipe: nil}, nil
}

// Recipes is the resolver for the recipes field.
func (r *queryResolver) Recipes(ctx context.Context, title *string) ([]*model.Recipe, error) {
	panic(fmt.Errorf("not implemented: Recipes - recipes"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ListRecipes(ctx context.Context) ([]*model.Recipe, error) {
	// list files in recipes directory
	files, err := os.ReadDir(recipesDir)
	if err != nil {
		return nil, fmt.Errorf("error reading recipes directory: %w", err)
	}

	recipes := []*model.Recipe{}
	for _, file := range files {
		// read recipe file
		recipeJSON, err := os.ReadFile(filepath.Join(recipesDir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("error reading recipe file: %w", err)
		}

		// decode recipe
		recipe := &model.Recipe{}
		err = json.Unmarshal(recipeJSON, recipe)
		if err != nil {
			return nil, fmt.Errorf("error decoding recipe: %w", err)
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
func (r *queryResolver) OpenRecipe(ctx context.Context, title string) (*model.Recipe, error) {
	// check if title is empty
	if title == "" {
		return nil, fmt.Errorf("title is required")
	}

	// construct recipe file path, based on title
	fileName := fmt.Sprintf("%s.json", title)
	filePath := filepath.Join(recipesDir, fileName)

	// check if recipe file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("recipe named %s not found", title)
		}
		return nil, fmt.Errorf("error checking recipe file: %w", err)
	}

	// read recipe file
	recipeJSON, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading recipe file: %w", err)
	}

	// decode recipe
	recipe := &model.Recipe{}
	err = json.Unmarshal(recipeJSON, recipe)
	if err != nil {
		return nil, fmt.Errorf("error decoding recipe: %w", err)
	}

	return recipe, nil
}
