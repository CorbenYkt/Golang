package main

import "fmt"

func main() {
	const rewardPoints = 10 //untyped constant with integer value

	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints) //int

	var totalRewardPoints float64 = 150.55 //if const will not work

	fmt.Printf("Week type of rewardPoints is %T\n", rewardPoints)

	//adding untyped const to float64 - valid
	totalRewardPoints = totalRewardPoints + rewardPoints

	fmt.Printf("Total rewardPoints is %.2f\n", totalRewardPoints)

}
