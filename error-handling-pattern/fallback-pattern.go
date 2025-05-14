package main

import "fmt"

func main() {
	var dataTreshold = map[string]int64{}{10}

	userLogin := []string{
		"john", "doe",
	}

	for i, usr := range userLogin {
		dataUser := userDatabase[usr]
		if dataUser.LoyaltyPoint < 1 {
			fmt.Println(`[log] Failed to get data loyalty from user...`)
			// continue process only print username
			fmt.Println(fmt.Sprintf("%d. Username: %s", i+1, usr))
			continue
		}

		fmt.Println(fmt.Sprintf("%d. Username: %s have loyalty point: %d", i+1, usr, dataUser.LoyaltyPoint))
	}

	// login is success
	fmt.Println(`[log] Successfully get all of data users`)
}
