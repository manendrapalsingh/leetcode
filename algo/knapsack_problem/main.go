package main

import (
	"fmt"
	"sort"
)

// fime the max profit if the max weight is alllowed is 15

type Items struct {
	Object         int
	Profit         int
	Weight         int
	ProfitByWeight float64
}

func main() {

	Object := []int{1, 2, 3, 4, 5, 6, 7}
	Profit := []int{10, 5, 15, 7, 6, 18, 3}
	Weight := []int{2, 3, 5, 7, 1, 4, 1}

	cap := 15

	/***  Apprich -> to identify the selection creiatare  to select the object
		              -> in this case the object are divied
					  -> in this case we will be considering max profit by weight
					  -> elemet with higest selecation crieatare will be selected first but need to take care that weight can not be
					  more than allowed weitgh


	***/

	items := make([]Items, len(Object))

	for i := range Object {
		items[i] = Items{
			Object:         Object[i],
			Profit:         Profit[i],
			Weight:         Weight[i],
			ProfitByWeight: float64(Profit[i] / Weight[i]),
		}
	}

	//-> need to sort items in decinding order
	sort.Slice(items, func(i, j int) bool {
		return items[i].ProfitByWeight > items[j].ProfitByWeight
	})

	currentProfit := 0.0
	currentWeight := 0.0
	selectedItems := make(map[int]float32)

	for _, item := range items {

		if currentWeight > float64(cap) {
			break
		}

		remainingCap := float64(cap) - currentWeight
		takeWeight := min(float64(item.Weight), remainingCap)

		selectedItems[item.Object] = float32(takeWeight)
		currentWeight += takeWeight
		currentProfit += item.ProfitByWeight * takeWeight

	}

	fmt.Printf("Maximum Profit: %.2f\n", currentProfit)
	fmt.Println("Selected Items (object: weight taken):")
	for obj, wt := range selectedItems {
		fmt.Printf("Item %d: %.2f kg\n", obj, wt)
	}

}
