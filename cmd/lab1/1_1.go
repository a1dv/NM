package lab1

import (
    "fmt"
)

type Massiv [][]float64

func Task1(a Massiv, b []float64) {
    //a := Massiv{{-9, 8, 8, 6},{-7, -9, 5, 4},{-3, -1, 8, 0},{3, -1, -4, -5}}
    //a := Massiv{{10, -7, 0},{-3, 6, 2}, {5, -1, 5}}
    l := Create_mass(len(a))
    u := Create_mass(len(a))
    LU(a, l, u)
    fmt.Println("A =")
    a.print_mass()
    fmt.Println("L =")
    l.print_mass()
    fmt.Println("U =")
    u.print_mass()
    res := multiply(l, u)
    fmt.Println("L * U =")
    res.print_mass()
    x := make([]float64, len(a))
    z := make([]float64, len(a))
    solve(u, a, l, z, x, b)
    fmt.Println("z is ", z)
    fmt.Println("x is ", x)
    fmt.Println("determinant is", determinant(l, u))
}

func solve(u Massiv, a Massiv, l Massiv, z []float64, x []float64, b []float64) {
    z[0] = b[0]
    for i := 1; i < len(a); i++ {
        z[i] = b[i]
        summ := float64(0)
        for j := 0; j < i; j ++ {
            summ += l[i][j] * z[j]
        }
        z[i] -= summ
    }
    x[len(a) - 1] = z[len(a) - 1] / u[len(a) - 1][len(a) - 1]
    for i := len(a) - 2; i >= 0; i -- {
        x[i] = (1 / u[i][i])
        summ2 := float64(0)
        for j := i + 1; j < len(a); j++ {
            summ2 += u[i][j] * x[j]
        }
        x[i] *= (z[i] - summ2)
    }
}

func determinant(l Massiv, u Massiv) float64 {
    res := float64(1)
    for i := 0; i < len(u); i++ {
        res *= l[i][i] * u[i][i]
    }
    return res
}

func LU(a Massiv, l Massiv, u Massiv) {
    for i := 0; i < len(a); i++ {
        l[i][i]++
    }
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a); j++ {
            if i <= j {
                u[i][j] = a[i][j] - summary(l, u, i, j, i)
            }
            if i > j {
                l[i][j] = (a[i][j] - summary(l, u, i, j, j)) / u[j][j]
            }
        }
    }
}

func summary(l Massiv, u Massiv, i int, j int, flag int) float64 {
    sum := float64(0)
    for k := 0; k < flag; k++ {
        sum += l[i][k] * u[k][j]
    }
    return sum
}

func Create_mass(size int) Massiv {
    l := make(Massiv, size)
    for i, _ := range l {
        l[i] = make([]float64, size)
    }
    return l
}

func multiply(l Massiv, u Massiv) Massiv {
    res := Create_mass(len(l))
    for i := 0; i < len(l); i++ {
		for j := 0; j < len(u[0]); j++ {
            tmp := float64(0)
			for k := 0; k < len(u); k++ {
				tmp += l[i][k] * u[k][j]
            }
            res[i][j] = tmp
        }
    }
    return res
}

func multiply_by_value(a Massiv, b float64) Massiv{
    res := undef_mass(len(a), len(a[0]))
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[0]); j++ {
            res[i][j] = a[i][j] * b
        }
    }
    return res
}

func subt(a Massiv, b Massiv) Massiv{
    res := undef_mass(len(a), len(b))
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a); j++ {
            res[i][j] = a[i][j] - b[i][j]
        }
    }
    return res
}

func (a Massiv) print_mass() {
    for i := range a {
        fmt.Println(a[i])
    }
    fmt.Println("==================================")
}
