package main

import (
    "fmt"
)


// func main() {
//     naturals := make(chan int)
//     squares := make(chan int)

//     // Counter
//     go func() {
//         for x := 0; x <= 10; x++ {
//             naturals <- x
//         }
//     }()

//     // // Squarer
//     // go func() {
//     //     for {
//     //         x := <-naturals
//     //         squares <- x * x
//     //     }
//     // }()


// // Squarer
//     go func() {
//         for {
//             x, ok := <-naturals
//             if !ok {
//                 break // channel was closed and drained
//             }
//             squares <- x * x
//         }
//         close(squares)
//     }()


//     // Printer (in main goroutine)
//     for {
//         fmt.Println(<-squares)
//     }
// }






func main() {
    naturals := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; x < 10; x++ {
            naturals <- x
        }
        close(naturals)  // for 完了 要close
    }()

    // Squarer       用迭代 channel   用完naturals close
    go func() {
        for x := range naturals {
            squares <- x * x
        }
        close(squares)        
    }()

    // Printer (in main goroutine)
    for x := range squares {
        fmt.Println(x)
    }
}









