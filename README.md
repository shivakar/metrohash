# metrohash
A pure Go port of metrohash algorithm

For more information about `metrohash`, see:

* https://github.com/jandrewrogers/MetroHash

## Installation

```bash
go get github.com/shivakar/metrohash
```

## Usage

```
package main

import (
	"fmt"
	"github.com/shivakar/metrohash"
)

func main() {
    // Create a new instance of the hash engine with default seed
    h := metrohash.NewMetroHash64()

    // Create a new instance of the hash engine with custom seed
    _ = metrohash.NewSeedMetroHash64(uint64(10))

    // Write some data to the hash
    h.Write([]byte("Hello, World!!"))

    // Write some more data to the hash
    h.Write([]byte("How are you doing?"))

    // Get the current hash as a byte array
    b := h.Sum(nil)
    fmt.Println(b)

    // Get the current hash as an integer (uint64) (little-endian)
    fmt.Println(h.Uint64())

    // Get the current hash as a hexadecimal string (big-endian)
    fmt.Println(h.String())

    // Reset the hash
    h.Reset()

    // Output:
    // [205 190 61 93 89 212 164 71]
    // 14825354494498612295
    // cdbe3d5d59d4a447
}
```

## License

`metrohash` is licensed under a MIT license.
