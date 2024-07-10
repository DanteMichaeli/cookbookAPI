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

// CreateRecipe is the resolver for the createRecipe field.
func (r *mutationResolver) CreateRecipe(ctx context.Context, id string, title string, description string, image string, ingredients []string, steps []string) (*model.Recipe, error) {
	//new recipe instance
	recipe := &model.Recipe{
		ID:          id,
		Title:       title,
		Description: description,
		Image:       image,
		Ingredients: ingredients,
		Steps:       steps,
	}

	//encode to JSON
	recipeJSON, err := json.Marshal(recipe)
	if err != nil {
		return nil, fmt.Errorf("error encoding recipe: %w", err)
	}

	// Directory info for saving recipes
	err = os.MkdirAll(recipesDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("error creating recipes directory: %w", err)
	}

	// Write recipe to file
	fileName := fmt.Sprintf("%s.json", id)
	filePath := filepath.Join(recipesDir, fileName)
	err = os.WriteFile(filePath, recipeJSON, 0644)
	if err != nil {
		return nil, fmt.Errorf("error writing recipe to file: %w", err)
	}

	return recipe, nil
}

// UpdateRecipe is the resolver for the updateRecipe field.
func (r *mutationResolver) UpdateRecipe(ctx context.Context, title string, description *string, image *string, ingredients []string, steps []string) (*model.Recipe, error) {
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

	// load recipe from JSON file
	recipeJSON, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading recipe file: %w", err)
	}

	// create recipe instance
	existingRecipe := &model.Recipe{}
	err = json.Unmarshal(recipeJSON, existingRecipe)
	if err != nil {
		return nil, fmt.Errorf("error decoding recipe: %w", err)
	}

	// apply updates
	if description != nil {
		existingRecipe.Description = *description
	}
	if image != nil {
		existingRecipe.Image = *image
	}
	if ingredients != nil {
		existingRecipe.Ingredients = ingredients
	}
	if steps != nil {
		existingRecipe.Steps = steps
	}

	// encode to JSON
	recipeJSON, err = json.Marshal(existingRecipe)
	if err != nil {
		return nil, fmt.Errorf("error encoding recipe: %w", err)
	}

	// write recipe to file
	err = os.WriteFile(filePath, recipeJSON, 0644)
	if err != nil {
		return nil, fmt.Errorf("error writing recipe to file: %w", err)
	}

	return existingRecipe, nil
}

// DeleteRecipe is the resolver for the deleteRecipe field.
func (r *mutationResolver) DeleteRecipe(ctx context.Context, title string) (*model.Recipe, error) {
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

	// delete recipe file
	err = os.Remove(filePath)
	if err != nil {
		return nil, fmt.Errorf("error deleting recipe file: %w", err)
	}

	return &model.Recipe{Title: title}, nil
}

// ListRecipes is the resolver for the listRecipes field.
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

// OpenRecipe is the resolver for the openRecipe field.
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
const recipesDir = "recipes"
