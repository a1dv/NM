package lab3

import (
    "fmt"
)

func Task4() {
    x := []float64{0., 0.1, 0.2, 0.3, 0.4}
    y := []float64{1., 1.1052, 1.2214, 1.3499, 1.4918}
    exactX := 0.2
    i := 0
    for i = 0; i < len(x); i++ {
        if x[i] <= exactX && exactX <= x[i + 1] {
            break
        }
    }
    fmt.Println(firstDerivative(x, y, exactX, i))
    fmt.Printf("%f", secondDerivative(x,y,exactX,i))
}

func firstDerivative(x []float64, y []float64, exactX float64, i int) float64{
    part1 := (y[i + 1] - y[i]) / (x[i + 1] - x[i])
    part2 := ((y[i + 2] - y[i + 1]) / (x[i + 2] - x[i + 1]) - part1) / (x[i + 2] - x[i]) * (2 * exactX - x[i] - x[i + 1])
    return part1 + part2
}

func secondDerivative(x []float64, y []float64, exactX float64, i int) float64{
    part1 := (y[i + 2] - y[i + 1]) / (x[i + 2] - x[i + 1])
    part2 := (y[i + 1] - y[i]) / (x[i + 1] - x[i])
    fmt.Println(part1, part2, x[i + 2] - x[i], part1 - part2, 2 *(part1 - part2), 0.246 / 0.19999999999999998)
    return 2 * (part1 - part2) / (x[i + 2] - x[i])
}
