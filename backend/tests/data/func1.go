package main

func funca(i int) int {
    j := i + 2 + 2 + i - 1
    return j
}

func funcb(i int, iii int) (a int, b int) {
    fff := 4
    fff = 5
    fff = fff + 100 + i + funca(fff)
    return fff, fff
}


func main() {
    re, re2 := funcb(3, 4)
} 