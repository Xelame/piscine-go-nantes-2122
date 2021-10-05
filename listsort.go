package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

// A commenter
func ListSort(l *NodeI) *NodeI {
	head := l
	// Condition d'arret
	if head == nil {
		return nil
	}
	// Récursivité pour allez jusqu'au dernier element
	head.Next = ListSort(head.Next)
	// + empile la suite de notre programme pour retourner au debut
	if head.Next != nil && head.Data > head.Next.Data {
		// Si l'élément d'avant est plus grand que celui d'apres on les interverties (leur pointage)
		head = Move(head)
	}
	// Retourne l'element le plus petit des 2
	return head
}

// Fonction pour swap 2 element et réarranger leur pointage
func Move(node *NodeI) *NodeI {
	// init nos deux element
	previous := node
	current := node.Next
	result := current
	// rearrange nos valeur jusqu'a ce que la valeur de base soit plus grande que sa précédente
	for current != nil && node.Data > current.Data {
		previous = current
		current = current.Next
	}
	previous.Next = node
	node.Next = current
	return result
}
