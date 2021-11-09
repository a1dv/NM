package main

import (
    "fmt"
    "github/NM/cmd/lab1"
    "github/NM/cmd/lab2"
    "github/NM/cmd/lab3"
    "github/NM/cmd/lab4"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigName("a")
    viper.AddConfigPath("./tests/")
    var aM map[string][][]float64
    var bM map[string][]float64
    eps := 0.01
    vipErr := viper.ReadInConfig()
    if vipErr != nil {
        panic(vipErr)
    }
    configErr := viper.Unmarshal(&aM)
    if configErr != nil {
        panic(configErr)
    }
    viper.SetConfigName("b")
    vipErr = viper.ReadInConfig()
    if vipErr != nil {
        panic(vipErr)
    }
    configErr = viper.Unmarshal(&bM)
    if configErr != nil {
        panic(configErr)
    }
    var l int
    for {
        fmt.Println("Enter the\n 1 - to run Lab1\n 2 - to run Lab2\n 3 - to run Lab3\n 4 - to run Lab4\n 5 - to run CP")
        fmt.Scan(&l)
        if l == 1 {
            for {
                var th int
                fmt.Println("Enter the\n 1 - to run theme1\n 2 - to run theme2\n 3 - to run theme3\n 4 - to run theme4\n 5 - to run theme5")
                fmt.Scan(&th)
                var a lab1.Massiv
                var b []float64
                if th == 1 {
                    var choice string
                    fmt.Println("------\n\n------\nType \"e\" to enter the data for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        a, b = reader()
                    } else if choice == "d" {
                        a = aM["a1"]
                        b = bM["b1"]
                    }
                    lab1.Task1(a, b)
                } else if th == 2 {
                    var choice string
                    fmt.Println("------\n\n------\nType \"e\" to enter the data for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        a, b = reader()
                    } else if choice == "d" {
                        a = aM["a2"]
                        b = bM["b2"]
                    }
                    lab1.Task2(a, b)
                } else if th == 3 {
                    var choice string
                    fmt.Println("------\n\n------\nType \"e\" to enter the data for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        a, b = reader()
                    } else if choice == "d" {
                        a = aM["a3"]
                        b = bM["b3"]
                    }
                    lab1.Task3(a, b, eps)
                } else if th == 4 {
                    var choice string
                    fmt.Println("------\n\n------\nType \"e\" to enter the data for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        a = short_reader()
                    } else if choice == "d" {
                        a = aM["a4"]
                    }
                    lab1.Task4(a, eps)
                }else if th == 5 {
                    var choice string
                    fmt.Println("------\n\n------\nType \"e\" to enter the data for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        a = short_reader()
                    } else if choice == "d" {
                        a = aM["a5"]
                    }
                    lab1.Task5(a, eps)
                } else {
                    break
                }
            }
        } else if l == 2 {
            for {
                fmt.Println("Enter the\n 1 - to run theme1\n 2 - to run theme2\n")
                var th int
                var choice string
                fmt.Scan(&th)
                if th == 1 {
                    fmt.Println("------\n\n------\nType \"e\" to enter the eps for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        fmt.Scan(&eps)
                    }
                    lab2.Task1(eps)
                } else if th == 2 {
                    fmt.Println("------\n\n------\nType \"e\" to enter the eps for theme 1 or type \"d\" to run with default input")
                    fmt.Scan(&choice)
                    if choice == "e" {
                        fmt.Scan(&eps)
                    }
                    lab2.Task2(eps)
                } else {
                    break
                }
            }
        } else if l == 3 {
            for {
                fmt.Println("Enter the\n 1 - to run theme1\n 2 - to run theme2\n 3 - to run theme3\n 4 - to run theme4\n 5 - to run theme5")
                var th int
                fmt.Scan(&th)
                if th == 1 {
                    lab3.Task1()
                } else if th == 2 {
                    lab3.Task2()
                } else if th == 3 {
                    lab3.Task3()
                } else if th == 4 {
                    lab3.Task4()
                } else if th == 5 {
                    lab3.Task5()
                } else {
                    break
                }
            }
        } else if l == 4 {
            for {
                fmt.Println("Enter the\n 1 - to run theme1\n 2 - to run theme2\n")
                var th int
                fmt.Scan(&th)
                if th == 1 {
                    lab4.Task1()
                }
                if th == 2 {
                    lab4.Task2()
                } else {
                    break
                }
            }
        }
    }
}


type config struct{
    eps float64 `json:"eps"`
    datastore1_1 dataStore1 `json:"datastore1_1"`
    datastore1_2 dataStore1 `json:"datastore1_2"`
    datastore1_3 dataStore1 `json:"datastore1_3"`
    datastore1_4 dataStore2 `json:"datastore1_4"`
    datastore1_5 dataStore2 `json:"datastore1_5"`
}

type dataStore1 struct{
    size int `json:"size"`
    a lab1.Massiv `json:"a"`
    b []float64 `json:"b"`
}

type dataStore2 struct{
    size int `json:"size"`
    a lab1.Massiv `json:"a"`
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
