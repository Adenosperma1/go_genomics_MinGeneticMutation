package main

import "fmt"

//NOT SURE THIS WORKS AS INTENDED???
//It finds if you change the last character but if you change the last two it errors?
//But I thought it would give you 2 for 2 mutations?

//keep a list of gene sequences and a bool for if they are valid
type GeneBank map[string]bool


//minMutation will find the minimum number of mutations to transform the starting gene to sequence to the end gene sequence
func minMutation(startGene string, endGene string, bank []string) int {
	bankSet := make(GeneBank)
	for _, gene := range bank {
		bankSet[gene] = true
	}
	if !bankSet[endGene] {
		return -1
	}

	visited := make(GeneBank)
	visited[startGene] = true

	//breadth First Search of graph
	return bfs(startGene, endGene, bankSet, visited)
}


func bfs(currentGene string, endGene string, bankSet GeneBank, visited GeneBank) int {
queue := []string{currentGene}
mutationsCount := 0

for len(queue) > 0{
	for i := 0; i < len(queue); i++ {
		currentGene := queue[0]
		queue = queue[1:] //remove the currentGene from the queue

		//have we found the endGene? How many mutations did it take?
		if currentGene == endGene {
			return mutationsCount
		}
		
		//Get other possible genes from the graph
		nextGenes := generateNextGenes(currentGene)
		
		for _, nextGene := range nextGenes {
			//is this gene not in the bankset or have we not visited it before
			if !bankSet[nextGene] || visited[nextGene] {
				continue
			}

			visited[nextGene] = true
			queue = append(queue, nextGene)
		}
	}
	mutationsCount++
}
return -1
}


//generateNextGenes generates all possible next genes from the current gene by changing one nucleotide at a time
func generateNextGenes (currentGene string) []string {
 mutations := []byte{'A', 'C', 'G', 'T'}
 nextGenes := make([]string, 0, len(currentGene)*4)

 for j := 0; j < len(currentGene); j++ {
	for _, mutation := range mutations {
		if mutation == currentGene[j] {
			continue
		}
		nextGene := currentGene[:j] + string(mutation) + currentGene[j+1:]
		nextGenes = append(nextGenes, nextGene)
	}
 }

 return nextGenes
}

func main(){
	startGene := "AACCGGTT"
	endGene := "AACCGGAA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA", "AAACGGAA"}

	result :=  minMutation(startGene, endGene, bank)

	//Returns -1 for not found, shoud return an int for the number of muations? Always seems to return 1
	fmt.Printf("Result %d\n", result)
}