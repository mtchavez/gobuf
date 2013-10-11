package gobuf

import (
    "fmt"
)

func Example_NewBuf() {
    buf := NewBuf()
    fmt.Printf("Bufer Size: %v\n", size)
    fmt.Printf("Bufer Index: %v\n", buf.index)
    fmt.Printf("Bufer Next Index: %v\n", buf.nextIndex)
    fmt.Printf("Bufer Get Index: %v\n", buf.getIndex)

    // Output:
    // Bufer Size: 8192
    // Bufer Index: 0
    // Bufer Next Index: 1
    // Bufer Get Index: 1
}

func ExampleAdd() {
    buf := NewBuf()
    buf.Add(35)
    fmt.Printf("Bufer[0] is %v\n", 35)
    fmt.Printf("Bufer Index: %v\n", buf.index)
    fmt.Printf("Bufer Next Index: %v\n", buf.nextIndex)
    fmt.Printf("Bufer Get Index: %v\n", buf.getIndex)

    // Output:
    // Bufer[0] is 35
    // Bufer Index: 1
    // Bufer Next Index: 2
    // Bufer Get Index: 1
}

func ExampleGet() {
    buf := NewBuf()
    buf.Add(35)
    buf.Add(100)
    fmt.Printf("Bufer Index: %v\n", buf.index)
    fmt.Printf("Bufer Next Index: %v\n", buf.nextIndex)
    fmt.Printf("Bufer Get Index: %v\n", buf.getIndex)

    fmt.Printf("\nBufer Get yielded: %v\n", buf.Get())
    fmt.Printf("Bufer Index: %v\n", buf.index)
    fmt.Printf("Bufer Next Index: %v\n", buf.nextIndex)
    fmt.Printf("Bufer Get Index: %v\n", buf.getIndex)

    fmt.Printf("\nBufer Get yielded: %v\n", buf.Get())
    fmt.Printf("Bufer Index: %v\n", buf.index)
    fmt.Printf("Bufer Next Index: %v\n", buf.nextIndex)
    fmt.Printf("Bufer Get Index: %v\n", buf.getIndex)

    // Output:
    // Bufer Index: 2
    // Bufer Next Index: 3
    // Bufer Get Index: 1
    //
    // Bufer Get yielded: 35
    // Bufer Index: 2
    // Bufer Next Index: 3
    // Bufer Get Index: 2
    //
    // Bufer Get yielded: 100
    // Bufer Index: 2
    // Bufer Next Index: 3
    // Bufer Get Index: 3
}
