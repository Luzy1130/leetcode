package main

import "fmt"

func matrixBlockSum(mat [][]int, K int) [][]int {
    var res [][]int 
    m := len(mat)
    n := len(mat[0])
    
    for i := 0; i < m; i++ {
        initTmp := make([]int, n)
        res = append(res, initTmp)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            rs := i - K
            re := i + K
            cs := j - K
            ce := j + K
            
            if rs < 0 {
                rs = 0
            }
            if re > m -1 {
                re = m-1
            }
            if cs < 0 {
                cs = 0
            }
            if ce > n -1 {
                ce = n-1
            }
            
            tmp := 0
            for ii := rs; ii <=re; ii++ {
                for jj := cs; jj <= ce; jj++ {
                    tmp += mat[ii][jj]   
                }
            } 
            res[i][j] = tmp
        }
    } 
    return res
}

func main() {
	mat := [][]int {
		[]int {1, 2, 3},
		[]int {4, 5, 6},
		[]int {7, 8, 9},
	}

	res := matrixBlockSum(mat, 2)

	fmt.Println(res)
}