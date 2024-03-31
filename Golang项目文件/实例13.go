package trans

import "math"

var Pi float64

func init() {
   Pi = 4 * math.Atan(1) // init() function computes Pi
}

package main

import (
   "fmt"
   "./trans"
)

var twoPi = 2 * trans.Pi

func main() {
   fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}