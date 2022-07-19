package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {
        "quarter_of_a_dozen":	3,
        "half_of_a_dozen":		6,
        "dozen":				12,
        "small_gross":			120,
        "gross":				144,
        "great_gross":			1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// check confirms if a key exists in a map.  
func check(userMap map[string]int, key string) (ok bool) {
	_, ok = userMap[key]
    return
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) (bool) {
	switch {
        case check(bill, item) || check(units, unit):
    	bill[item] += units[unit]
        return true
    }
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) (bool) {
	switch {
		case !check(bill, item) || !check(units, unit) || (bill[item] - units[unit]) < 0:
    	return false
        case (bill[item] - units[unit]) == 0:
    	delete(bill, item)
        default:
    	bill[item] -= units[unit]
    }
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if !check(bill, item) {
    	return 0, false
    }
	return bill[item], true
}
