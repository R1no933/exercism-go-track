package matrix

// Define the Matrix type here.
import (
	"errors"
	"strconv"
	"strings"
)

// import "fmt"
type Matrix [][]int

func New(s string) (Matrix, error) {
	m := Matrix{}
	//     fmt.Println(">>", strings.Split(s, "\n"), s)
	for _, row := range strings.Split(s, "\n") {
		mRow := []int{}
		for _, c := range strings.Split(row, " ") {
			if c == "" {
				continue
			}
			i, err := strconv.Atoi(c)
			if err != nil {
				return m, err
			}
			mRow = append(mRow, i)
		}
		if len(m) > 0 && len(mRow) != len(m[len(m)-1]) {
			return m, errors.New("Uneven rows of matrix")
		}
		m = append(m, mRow)
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for j, _ := range m[0] {
		cols[j] = make([]int, len(m))
		i := 0
		for {
			if i >= len(m) {
				break
			}
			cols[j][i] = m[i][j]
			i++
		}
	}
	return cols
}
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, row := range m {
		rows[i] = make([]int, len(row))
		for j, elm := range row {
			rows[i][j] = elm
		}
	}
	return rows
}
func (m Matrix) Set(row, col, val int) bool {
	if row >= len(m) || col >= len(m[0]) || col < 0 || row < 0 {
		return false
	}
	m[row][col] = val
	return true
}
