package main

import (
    "fmt"
    "math/rand"
    )
    
func bestCase(voidIndex int)int{ 
    smallest := 100 
    heuristicIndexes := make(map[int]int)
    randHeuristics := make([]int,9)
    for index := 0 ; index < len(currentArr) ; index++{ 
        if (index == voidIndex - 3 || index == voidIndex + 3 || ( index == voidIndex - 1 && voidIndex % 3 != 0) || (index == voidIndex + 1 && ( voidIndex != 1 || voidIndex != 5 || voidIndex != 8 )) || (voidIndex == 4 && index != 4)){ 
            tempCurrenArr := currentArr 
            tempCurrenArr[index],tempCurrenArr[voidIndex] = tempCurrenArr[voidIndex], tempCurrenArr[index]
            heuristicIndexes[index] = calcHeuristic(tempCurrenArr, correctAns) 
        } 
    } 
    fmt.Println(heuristicIndexes) 
    for _, small := range heuristicIndexes{ 
        if small <= smallest{ 
            smallest = small
         } 
    } 
    for index, small := range heuristicIndexes{ 
        if smallest == small /*|| smallest == small + 1*/{ 
            randHeuristics = append(randHeuristics, index) 
        } 
    } 
    bestCase := rand.Intn(len(randHeuristics))
    fmt.Println(randHeuristics) 
    return randHeuristics[bestCase] 
}

func calcHeuristic(curr []int,ans []int)int{ 
    value := 0 
    for i := 0 ; i < len(ans) ; i++ { 
        if curr[i] != ans[i]{ 
            value = value + 1 
        } 
    } 
    return value 
}
var correctAns = []int{1,2,3,8,0,4,7,6,5}
var currentArr = []int{0,3,1,5,2,4,6,7,8}
func main() {
	
	
	heuristic := calcHeuristic(currentArr,correctAns) 
     
    if heuristic > 0 { 
        voidIndex := 0
        for i := 0 ; i < len(currentArr) ; i++ { 
            if currentArr[i] == 0 {
                voidIndex = i
            }
        } 
        indexToSwap := bestCase(voidIndex) 
        //swap
        voidIndex, indexToSwap = indexToSwap, voidIndex
        heuristic = calcHeuristic(currentArr, correctAns) 
    } 
}

 
