package main

import (
    "fmt"
    "github/NM/cmd/lab1"
    "github/NM/cmd/lab2"
)

func main() {
    var l int
    for {
        fmt.Println("Enter the\n 1 - to run Lab1\n 2 - to run Lab2\n 3 - to run Lab3\n 4 - to run Lab4\n 5 - to run Lab5")
        fmt.Scan(&l)
        if l == 1 {
            for {
                var th int
                fmt.Println("Enter the\n 1 - to run theme1\n 2 - to run theme2\n 3 - to run theme3\n 4 - to run theme4\n 5 - to run theme5")
                fmt.Scan(&th)
                var a lab1.Massiv
                var b []float64
                var eps float64
                if th == 1 {
                    fmt.Println("------\n\n------\nEnter the data for theme 1")
                    a, b := reader()
                    lab1.Task1(a, b)
                } else if th == 2 {
                    fmt.Println("------\n\n------\nEnter the data for theme 2")
                    a, b = reader()
                    lab1.Task2(a, b)
                } else if th == 3 {
                    fmt.Println("------\n\n------\nEnter the data for theme 3")
                    fmt.Scan(&eps)
                    a, b = reader()
                    lab1.Task3(a, b, eps)
                } else if th == 4 {
                    fmt.Println("------\n\n------\nEnter the data for theme 4")
                    fmt.Scan(&eps)
                    a = short_reader()
                    lab1.Task4(a, eps)
                }else if th == 5 {
                    fmt.Println("------\n\n------\nEnter the data for theme 5")
                    fmt.Scan(&eps)
                    a = short_reader()
                    lab1.Task5(a, eps)
                }
            }
        } else if l == 2 {
            var eps float64
            fmt.Println("Enter the\n 1 - to run theme1\n 2 - to run theme2\n")
            var th int
            fmt.Scan(&th)
            if th == 1 {
                fmt.Scan(&eps)
                lab2.Task1(eps)
            }
        }
    }
}

func reader() (lab1.Massiv, []float64){
    var n int
    fmt.Scan(&n)
    a := lab1.Create_mass(n)
    b := make([]float64, n)
    for i := 0; i < n; i++ {
        for j := 0; j < n + 1; j++ {
            if j == n {
                fmt.Scan(&b[i])
            } else {
                fmt.Scan(&a[i][j])
            }
        }
    }
    return a, b
}

func short_reader() lab1.Massiv{
    var n int
    fmt.Scan(&n)
    a := lab1.Create_mass(n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Scan(&a[i][j])
        }
    }
    return a
}
