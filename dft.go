package main

import (
  "fmt"
  "math"
)

// Ported from https://www.nayuki.io/page/how-to-implement-the-discrete-fourier-transform
func computeDft(inreal, inimag, outreal, outimag []float64) {
  n := float64(len(inreal))
  for k := 0.0; k < n; k++ {  // For each output element
    sumreal := 0.0
    sumimag := 0.0
    for t := 0.0; t < n; t++ {  // For each input element
      angle := 2.0 * math.Pi * t * k / n
      ti := int(t)
      sumreal +=  inreal[ti] * math.Cos(angle) + inimag[ti] * math.Sin(angle)
      sumimag += -inreal[ti] * math.Sin(angle) + inimag[ti] * math.Cos(angle)
    }
    ki := int(k)
    outreal[ki] = sumreal
    outimag[ki] = sumimag
  }
}
