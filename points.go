package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func countAlphanumericCharacters(str string) int64 {
	var count int64 = 0
	for _, c := range str {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			count++
		}
	}
	return count
}

// CalculatePoints calculates the points for a given receipt
// The points are calculated using the following rules:

// a) One point for every alphanumeric character in the retailer name.
// b) 50 points if the total is a round dollar amount with no cents.
// c) 25 points if the total is a multiple of 0.25.
// d) 5 points for every two items on the receipt.
// e) If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// f) 6 points if the day in the purchase date is odd.
// g) 10 points if the time of purchase is after 2:00pm and before 4:00pm.

func CalculatePoints(receipt Receipt) int64 {
	var points int64 = 0

	// One point for every alphanumeric character in the retailer name.
	points += countAlphanumericCharacters(receipt.Retailer)
	fmt.Println("Retailer points: ", points)

	// 50 points if the total is a round dollar amount with no cents.
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
		fmt.Println("Ends in .00 points added: 50")
	}
	// 25 points if the total is a multiple of 0.25.
	var quarterMultiples [4]string = [4]string{".00", ".25", ".50", ".75"}
	for _, multiple := range quarterMultiples {
		if strings.HasSuffix(receipt.Total, multiple) {
			points += 25
			fmt.Println("Ends in ", multiple, " points added: 25")
			break
		}
	}
	// 5 points for every two items on the receipt.
	points += int64(len(receipt.Items) / 2 * 5)
	fmt.Println(len(receipt.Items)/2, " pairs found, points added: ", int64(len(receipt.Items)/2*5))

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			points += int64(math.Ceil(0.2 * price))
			fmt.Println("Item description is multiple of 3, points added: ", int64(math.Ceil(0.2*price)))
		}
	}

	// 6 points if the day in the purchase date is odd.
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	intDay, err := strconv.Atoi(dateParts[2])
	if err == nil && intDay%2 != 0 {
		points += 6
		fmt.Println("Day is odd, points added: 6")
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	timeParts := strings.Split(receipt.PurchaseTime, ":")
	if len(timeParts) != 2 {
		fmt.Println("Invalid time format")
	}
	intHour, _ := strconv.Atoi(timeParts[0])
	if intHour >= 14 && intHour <= 16 {
		points += 10
		fmt.Println("Time is between 2:00pm and 4:00pm, points added: 10")
	}
	return points
}
