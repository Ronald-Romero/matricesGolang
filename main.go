package main

import (
	"fmt"
	"strings"
	"strconv"
)

var vect [][]float32
var vector [][]float32
var posiciones int
var tipo_calculo int = 7
var mtrx [][]float32
var c []int
var fil int
var column int

func longitud_1(mat [][]float32) {
	for _, col := range mat {
		vector[0] = make([]float32, len(col))
		for i := 0; i < len(col); i++ {
			vector[0][i] = col[i]
		}
	}
}

func longitud_2(mat [][]float32) {
	vector[0] = make([]float32, len(mat))
	for x, rows := range mat {
		vector[0][x] = rows[0]
	}
}

func reshape_1(mat [][]float32, f int, c []int) [][]float32{ 	
	vector = make([][]float32, 1)
		if len(mat) < 2 {
			longitud_1(mat)
			column = 1
			posiciones = len(c)
		}else{
			column = len(c)
			posiciones = 1
			longitud_2(mat)			
		}
		vect = make([][]float32, column)
		vect[0] = make([]float32, 3)
		for i := 0; i < len(c); i++ {
			if len(mat) > 1 {
				vect[i] = make([]float32, 1)
			}
			if len(mat) < 2 {
				vect[0][i] = vector[0][c[i]]
			}else{
				vect[i][0] = vector[0][c[i]]
			}
		}
		
	//fmt.Println("vect 1: ", vect)
	return vect
}

func reshape_2(mat [][]float32, f int, c []int) [][]float32{ 
		vect = make([][]float32, f + 1)
		vector = make([][]float32, len(mat))
		for i,  raws := range mat {
			vector[i] = make([]float32, len(mat))
			split := strings.Split(fmt.Sprintf("%g", raws[0]), ".")
			for j, colum := range split {
				b2, _ := strconv.ParseFloat(colum, 32)
				var float float32
				float = float32(b2)
				vector[i][j] = float
			}
		}
		//fmt.Println("vector", vector[1][1])
		for fila := 0; fila <= f; fila++ {
			vect[fila] = make([]float32, len(c))
			for j, co := range c {
				vect[fila][j] = vector[fila][co]
			}
		}		
	//fmt.Println("vect 2: ", vect)
	return vect
}

func main() {
	
    switch tipo_calculo {
    case 1:
        mtrx = [][]float32{{1, 2, 3, 4}}
        c = []int{0, 0, 2}
		fil = 0
    case 2:
        mtrx = [][]float32{{1.2}, {3.4}}
        c = []int{0}
		fil = 0
    case 3:
        mtrx = [][]float32{{1.2}, {3.4}}
        c = []int{0, 0}
		fil = 0
    case 4:
        mtrx = [][]float32{{1.2}, {3.4}}
        c = []int{0, 0, 1, 1}
		fil = 0
	case 5:
        mtrx = [][]float32{{1.2}, {3.4}}
        c = []int{0}
		fil = 1
	case 6:
        mtrx = [][]float32{{1.2}, {3.4}}
        c = []int{0, 0}
		fil = 1
	case 7:
        mtrx = [][]float32{{1.2}, {3.4}}
        c = []int{0, 0, 1, 1}
		fil = 1
	default:
		fmt.Println("Tipo de calculo no existe:", tipo_calculo)
    }
	
	if tipo_calculo <= 4 {
		fmt.Println("reshape1")
		//reshape1(mtrx, fil, c)
		fmt.Println("vect 1: ", reshape_1(mtrx, fil, c))
	}else{
		fmt.Println("reshape2")
		//reshape2(mtrx, fil, c)
		fmt.Println("vect 2: ", reshape_2(mtrx, fil, c))
	}
}