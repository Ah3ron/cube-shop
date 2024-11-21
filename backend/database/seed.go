package database

import "cube-shop/models"

func SeedData() {
	products := []models.Product{
		{
			Name:        "Classic 3x3",
			Description: "Standard speed cube",
			Price:       9.99,
			Stock:       50,
			Category:    "3x3",
			Difficulty:  "Medium",
			Brand:       "SpeedCube",
			ImageURL:    "https://placehold.co/400/141414/FFFFFF/png",
		},
		{
			Name:        "Pro 4x4",
			Description: "Professional cube",
			Price:       24.99,
			Stock:       30,
			Category:    "4x4",
			Difficulty:  "Hard",
			Brand:       "SpeedCube",
			ImageURL:    "https://placehold.co/400/141414/FFFFFF/png",
		},
		{
			Name:        "Beginner 2x2",
			Description: "Perfect for starters",
			Price:       7.99,
			Stock:       100,
			Category:    "2x2",
			Difficulty:  "Easy",
			Brand:       "CubeMaster",
			ImageURL:    "https://placehold.co/400/141414/FFFFFF/png",
		},
		{
			Name:        "Pyramid Cube",
			Description: "Triangle puzzle",
			Price:       12.99,
			Stock:       45,
			Category:    "Special",
			Difficulty:  "Medium",
			Brand:       "PuzzlePro",
			ImageURL:    "https://placehold.co/400/141414/FFFFFF/png",
		},
		{
			Name:        "Mirror Cube",
			Description: "Shape-shifting puzzle",
			Price:       15.99,
			Stock:       35,
			Category:    "Special",
			Difficulty:  "Hard",
			Brand:       "CubeMaster",
			ImageURL:    "https://placehold.co/400/141414/FFFFFF/png",
		},
	}

	for _, product := range products {
		DB.Create(&product)
	}
}
