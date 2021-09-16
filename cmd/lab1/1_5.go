package lab1

import(
    "fmt"
    "math"
)

func Task5(a Massiv, eps float64) {
    q, r := QR_decompose(a, eps)
    eigen_qr(q, r, a, eps)
}

func QR_decompose(m Massiv, eps float64) (Massiv, Massiv){
    h_comp := ed_mat(len(m))
    a := m
    e := ed_mat(len(a))
    for i := 0; i < len(a) - 1; i ++ {
        v := Undef_mass(1, len(a))
        for j := i; j < len(a); j++ {
            if j == i {
                v[0][j] = a[j][j] + sign(a[j][j]) * euclid(a, j, i)
            } else {
                v[0][j] = a[j][i]
            }
        }
        for j := 0; j < i; j ++ {
            v[0][j] = 0
        }
        h := Subt(e, multiply_by_value(multiply(transp(v), v), (float64(2) / multiply(v, transp(v))[0][0])))
        fmt.Printf("H_%d is\n",i)
        h.print_mass()
        a = multiply(h, a)
        fmt.Printf("A_%d is\n",i)
        a.print_mass()
        h_comp = multiply(h_comp, h)
    }
    fmt.Printf("Q * R is\n")
    multiply(h_comp, a).print_mass()
    fmt.Printf("A is\n")
    m.print_mass()
    return h_comp, a
}

func eigen_qr(q Massiv, r Massiv, m Massiv, eps float64) {
    eigen_i := 0
    eigen_j := 0
    eigen_el := m[eigen_i][eigen_j]
    count := 0
    a := multiply(r, q)
    for {
        a = multiply(r, q)
        fmt.Printf("A_%d is\n",count + 1)
        a.print_mass()
        eigen_el = a[eigen_i][eigen_j]
        count++
        if check_values(a, eigen_j, eps) {
            break;
        }
        q,r = QR_decompose(a, eps)
    }
    x1, x2 := Quadratic_roots(a, eigen_i)
    fmt.Println("Eigen values:")
    fmt.Println(eigen_el, x1, x2)
}

func Quadratic_roots(a Massiv, eigen_i int) (complex128, complex128) {
    tmp := eigen_i + 1
    var x1 complex128
    var x2 complex128
    c := (a[tmp][tmp] * a[tmp + 1][tmp + 1]) - ((a[tmp + 1][tmp] * a[tmp][tmp + 1]))
    b := - a[tmp][tmp] - a[tmp + 1][tmp + 1]
    d := (b * b) - (4 * c)
    if d > 0 {
        x1 = complex((-b - math.Sqrt(d)) / 2, 0)
        x2 = complex((-b + math.Sqrt(d)) / 2, 0)
    } else if d == 0 {
        x1 = complex(-b / 2, 0)
        x2 = complex(-b / 2, 0)
    } else if d < 0{
        x1 = complex(-b / 2, -math.Sqrt(-d) / 2)
        x2 = complex(-b / 2, math.Sqrt(-d) / 2)
    }
    return x1, x2
}

func check_values(a Massiv, b int, eps float64) bool{
    for i := b + 1; i < len(a); i++ {
        if a[i][b] > eps {
            return false
        }
    }
    return true
}

func euclid (a Massiv, i int, j int) float64{
    sum := float64(0)
    for k := j; k < len(a); k++ {
        sum += a[k][i] * a[k][i]
    }
    return math.Sqrt(sum)
}

func sign(a float64) float64 {
    if a < 0 {
        return float64(-1)
    }
    return float64(1)
}

func ed(a Massiv) {
    for i := 0; i < len(a);i++ {
        for j := 0; j < len(a[0]); j++ {
            a[i][j] = float64(1)
        }
    }
}

func ed_mat(size int) Massiv{
    e := Create_mass(size)
    for i := 0; i < size; i++ {
        e[i][i] = 1
    }
    return e
}

func transp(a Massiv) Massiv {
    res := Undef_mass(len(a[0]), len(a))
    for i := 0; i < len(a[0]); i++ {
        for j := 0; j < len(a); j++ {
            res[i][j] = a[j][i]
        }
    }
    return res
}

func Undef_mass(l int, w int) Massiv{
    res := make(Massiv, l)
    for i := 0; i < l; i++ {
        res[i] = make([]float64, w)
    }
    return res
}
