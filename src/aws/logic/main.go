package main

func main(stateJson string, ordersJson string) {
	var orders = parseOrders(ordersJson)
	var state = parseState(stateJson)
}


func parseOrders(orders string) []Order {

}

func parseState(state string) State {

}


func performNonCollidingAttacks(orders []Order) []Movement {

}


func GetUniqueTargets(orders []Order) []Order {
	seen := make(map[int]struct{}, len(orders))
	j := 0
	for _, v := range orders {
		if _, ok := seen[v.movementTo]; ok {
			continue
		}
		seen[v] = struct{}{}
		orders[j] = v
		j++
	}
	return orders[:j]
}