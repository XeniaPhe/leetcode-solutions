package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
    for fuel, deficit, start, curr, end := 0, 0, 0, 0, len(gas) - 1; curr <= end; curr += 1 {
        if fuel += gas[curr] - cost[curr]; fuel < 0 {
            start, deficit, fuel = curr + 1, deficit - fuel, 0
        } else if curr == end {
            if deficit <= fuel {
                return start
            } else {
                return -1
            }
        }
    }

    return -1
}