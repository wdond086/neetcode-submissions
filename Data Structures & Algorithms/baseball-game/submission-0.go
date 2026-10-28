func calPoints(operations []string) int {
	record := []int{}
	for _, j := range(operations) {
		switch j {
			case "+":
				record = append(record, record[len(record) - 1] + record[len(record) - 2])
			case "D":
				record = append(record, record[len(record) - 1] * 2)
			case "C":
				record = record[:len(record) - 1]
			default:
				intVal, err := strconv.Atoi(j)
				if err != nil {
					panic("unexpected string impossible to convert to int")
				}
				record = append(record, intVal)
		}
	}

	sum := 0
	for _, i := range(record) {
		sum += i
	}

	return sum
}
