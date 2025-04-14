package main

import (
	"fmt"

	flow "github.com/amitsuthar69/flow/internal"
)

const ascii = ` ___  _    ___  _ _ _ 
| __]| |  | . || | | |
| _] | |_ | | || | | |
|_|  |___|'___'|__/_/
© Amit Suthar | github.com/amitsuthar69
`

func main() {
	fmt.Printf("%v\n", ascii)

	flow.Watch()
}
