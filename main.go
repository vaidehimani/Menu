package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Recipe struct represents a recipe with basic attributes
type Recipe struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    MealType    string   `json:"mealType"`
    Ingredients []string `json:"ingredients"`
    Tags        []string `json:"tags"`
    PrepTime    string   `json:"prepTime,omitempty"`
}

// Dummy data - a list of recipes
var recipes = []Recipe{
    {
        ID: 1,
        Name: "Paneer Butter Masala",
        MealType: "Dinner",
        Ingredients: []string{"Paneer", "Butter", "Spices"},
        Tags: []string{"veg", "gravy"},
        PrepTime: "30 minutes",
    },
    {
        ID: 2,
        Name: "French toast",
        MealType: "Breakfast",
        Ingredients: []string{"Bread", "milk", "egg"},
        Tags: []string{"non veg", "quick", "egg"},
        PrepTime: "10 minutes",
    },
    {
        ID: 3,
        Name: "Yellow Dal",
        MealType: "Lunch",
        Ingredients: []string{"toor dal", "onions", "Ghee"},
        Tags: []string{"veg", "dal"},
        PrepTime: "20 minutes",
    },
}

// GetRecipesHandler handles the GET request to return the list of recipes
func GetRecipesHandler(response http.ResponseWriter, request *http.Request) {
    // Set the Content-Type header to application/json
    response.Header().Set("Content-Type", "application/json")
    
    // Convert the recipes slice to JSON and send it as the response
    json.NewEncoder(response).Encode(recipes)
}

func main() {
    // Set up the GET route to handle the request for recipes
    http.HandleFunc("/recipes", GetRecipesHandler)

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
