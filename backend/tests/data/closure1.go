package main

func main() {
    i := 0
    a := 44
    b := func() int {
        return a + 10
    }
    c := b() + 10
}
//64