package piscine

func AdvancedSortWordArr(table []string, f func(a, b string) int) {
	for min_index := 0; min_index < len(table)-1; min_index++ {
		for index := min_index + 1; index < len(table); index++ {
			if f(table[index], table[min_index]) == -1 {
				SwapString(&table[index], &table[min_index])
			}
		}
	}
}
