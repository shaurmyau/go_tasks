package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var filtered []BrainrotMeme
	for _, meme := range memes {
		if meme.Views > minViews {
			filtered = append(filtered, meme)
		}
	}
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].TrendLevel > filtered[j].TrendLevel
	})
	return filtered
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)
	for _, meme := range memes {
		impact[meme.Category] += meme.Views
	}
	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var result []string
	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			result = append(result, meme.Name)
		}
	}
	return result
}

func main() {
	memes := []BrainrotMeme{
		{"Skibidi Toilet", 10, "Skibidi", 120.5},
		{"GigaChad", 9, "Sigma", 85.2},
		{"На кондициях", 8, "Subo Bratik", 45.7},
		{"Pepe the frog", 6, "Other", 392.1},
		{"Ohio Final Boss", 7, "Sigma", 32.4},
		{"Это моя машина", 5, "Subo Bratik", 78.9},
		{"Skibidi Dop Dop", 10, "Skibidi", 210.3},
		{"Mewing Tutorial", 3, "Mewing", 15.8},
		{"TUNTUN Sahur", 8, "TUNTUNTUNSAHUR", 67.4},
		{"Sigma Rules", 9, "Sigma", 52.6},
	}

	fmt.Println("Топ трендовые мемы (просмотров > 40)")
	topTrending := FindTopTrending(memes, 40)
	for _, meme := range topTrending {
		fmt.Printf("%s: TrendLevel=%d, Views=%.1fM\n", meme.Name, meme.TrendLevel, meme.Views)
	}

	fmt.Println("\nВлияние по категориям (сумма просмотров)")
	categoryImpact := CalculateCategoryImpact(memes)
	for category, totalViews := range categoryImpact {
		fmt.Printf("%s: %.1fM просмотров\n", category, totalViews)
	}

	fmt.Println("\nМемы по сложному условию")
	filteredNames := FilterByComplexCondition(memes)
	for _, name := range filteredNames {
		fmt.Println(name)
	}
}