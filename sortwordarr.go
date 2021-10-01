package piscine

func SortWordArr(table []string) {
	for min_index := 0; min_index < len(table)-1; min_index++ {
		for index := min_index + 1; index < len(table); index++ {
			if table[index] < table[min_index] {
				SwapString(&table[index], &table[min_index])
			}
		}
	}
}

func SwapString(a *string, b *string) {
	*a, *b = *b, *a
}
