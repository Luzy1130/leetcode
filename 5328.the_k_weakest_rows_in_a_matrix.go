// package main
// import (
// 	"fmt"
// 	"sort"
// )

type IndexSum struct {
    Sum int
    Index int
}

type IndexSums []IndexSum

func (s IndexSums) Len() int{ return len(s) }
func (s IndexSums) Less(i, j int) bool { 
	if s[i].Sum < s[j].Sum {
		return true
	} else if s[i].Sum == s[j].Sum {
		return s[i].Index < s[j].Index
	} else {
		return false
	}
}
func (s IndexSums) Swap(i, j  int) { s[i], s[j] = s[j], s[i] }

func kWeakestRows(mat [][]int, k int) []int {
    sums := IndexSums(make([]IndexSum, len(mat)))
    
    for i := 0; i < len(mat); i++ {
        sums[i].Sum = 0
        for j := 0; j < len(mat[0]); j++ {
            sums[i].Sum += mat[i][j]
        }
        sums[i].Index = i
    }
    // fmt.Println(sums)
    sort.Sort(sums)
    // fmt.Println(sums)
    var res []int
    for i := 0; i < k; i++ {
        res = append(res, sums[i].Index)
    }
    return res
}

// func main() {
// 	var input [][]int = [][]int {
// 		[]int{1,1,0},
// 		[]int{1,1,0},
// 		[]int{1,1,1},
// 		[]int{1,1,1},
// 		[]int{0,0,0},
// 		[]int{1,1,1},
// 		[]int{1,0,0},
// 	}
// 	res := kWeakestRows(input, 6)
// 	fmt.Println(res)

// }