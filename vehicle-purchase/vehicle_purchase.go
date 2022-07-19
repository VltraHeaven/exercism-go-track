package purchase

import (
    "strings";
    "fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true
    }
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    var carLexOrder = strings.Compare(option1, option2)
	if carLexOrder == 1 {
        return fmt.Sprint(option2, " is clearly the better choice.")
    }
	return fmt.Sprint(option1, " is clearly the better choice.")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3.0 {
    	return originalPrice * 0.8   
    } else if age >= 10.0 {
    	return originalPrice * 0.5
    }
	return originalPrice * 0.7
}
