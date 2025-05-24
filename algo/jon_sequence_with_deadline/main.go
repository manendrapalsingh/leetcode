package main

import (
	"fmt"
	"sort"
)

type Jobs struct {
	profit   int
	deadLine int
}

func main() {

	profit := []int{5, 12, 35, 30, 15, 20, 25}
	deadLine := []int{4, 1, 3, 4, 3, 2, 4}

	jobs := make([]Jobs, len(profit))

	maxTimeLimt := 0

	for i := range profit {
		jobs[i] = Jobs{
			profit:   profit[i],
			deadLine: deadLine[i],
		}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].profit > jobs[j].profit
	})

	for i := range deadLine {
		maxTimeLimt = max(maxTimeLimt, deadLine[i])
	}

	schedule := make([]int, maxTimeLimt)
	totalProfit := 0

	for _, job := range jobs {

		for i := job.deadLine - 1; i >= 0; i-- {

			if schedule[i] == 0 {

				schedule[i] = job.profit
				totalProfit = +job.profit

				break
			}

		}
	}

	// Print results
	fmt.Println("Scheduled Jobs Profit:", schedule)
	fmt.Println("Total Profit:", totalProfit)
	fmt.Println("Optimal Job Sequence:")
	for i, profit := range schedule {
		if profit != 0 {
			fmt.Printf("Time slot %d: Profit %d\n", i+1, profit)
		}
	}

}
