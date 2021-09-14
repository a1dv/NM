package lab1

import (
    "fmt"
    "math"
)

func Task4(a Massiv, eps float64) {
    Jacobi_rotative(a, eps)
}

func Jacobi_rotative(m Massiv,eps float64) {
    n := len(m)
    a := m
    e := Create_mass(n)
    for i := 0; i != n; i++ {
        e[i][i] = float64(1)
    }
    criteria := float64(0)
    u_comp := e
    count := 0
    u := e
    for {
        max := float64(-1e-10)
        i_max := 0
        j_max := 0
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if j != i {
                    if Abs(a[i][j]) > Abs(max) {
                        max = a[i][j]
                        i_max = i
                        j_max = j
                    }
                }
            }
        }

        phi := float64(0)
        if (Abs(a[i_max][i_max] - a[j_max][j_max]) < eps) {
            phi = math.Pi / 4
        } else {
            phi = math.Atan((2 * a[i_max][j_max]) / (a[i_max][i_max] - a[j_max][j_max])) / 2
        }

        u = Create_mass(n)
        for i := 0; i != n; i++ {
            u[i][i] = float64(1)
        }
        u[i_max][j_max], u[i_max][i_max] = math.Sincos(phi)
        u[j_max][j_max] = u[i_max][i_max]
        u[j_max][i_max] = - u[i_max][j_max]

        u_comp = multiply(u_comp, u)
        a = multiply(multiply(Transpose(u), a), u)
        criteria = t(a)
        fmt.Println("max |a| = ", max)
        fmt.Println("phi = ", phi)
        fmt.Printf("U_%d is \n",count)
        u.print_mass()
        fmt.Printf("A_%d is \n",count + 1)
        a.print_mass()
        fmt.Println("t = ", criteria)

        count++;

        if criteria < eps {
            break
        }
    }
    val := make([]float64, n)
    vect := Create_mass(n)
    for i := 0; i < n; i++ {
        val[i] = a[i][i]
    }
    for j := 0; j < n; j++ {
        for i := 0; i < n; i++ {
            vect[j][i] = u_comp[i][j]
        }
    }
    fmt.Println("values:\n", val)
    fmt.Println("vectors:")
    u_comp.print_mass()
}


func Transpose (a Massiv) Massiv {
    b := Create_mass(len(a))
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a); j++ {
            b[j][i] = a[i][j]
        }
    }
    return b
}

func t (a Massiv) float64{
    n := len(a)
    sum := float64(0)
    for i := 0; i != n; i++ {
        for j := 0; j < n; j++ {
            if j != i {
                sum += a[i][j] * a[i][j]
            }
        }
    }
    return math.Sqrt(sum)
}

func mult(l Massiv, u Massiv) Massiv {
    res := Create_mass(len(u))
    for i := 0; i < len(u); i++ {
		for j := 0; j < len(u); j++ {
            tmp := float64(0)
			for k := 0; k < len(u); k++ {
                fmt.Println(tmp, l[i][k] * u[k][j], i, j, k)
				tmp += l[i][k] * u[k][j]
            }
            res[i][j] = tmp
            fmt.Println("res =",res[i][j], i, j)
        }
    }
    return res
}
