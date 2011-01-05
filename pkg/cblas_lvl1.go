// simplistic wrapper around the C-interface to the BLAS
package cblas

//#include "cblas.h"
//#include <stdlib.h>
//#include <string.h>
import "C"
import "unsafe"

func Sdsdot(N int, alpha float, 
	X []float, incX int,
	Y []float, incY int) float {
	c_X := (*C.float)(unsafe.Pointer(&X[0]))
	c_Y := (*C.float)(unsafe.Pointer(&Y[0]))

	return float(
		C.cblas_sdsdot(C.int(N), C.float(alpha),
		 c_X, C.int(incX),
		 c_Y, C.int(incY)))
}

// functions having standard 4 prefixes (S D C Z)

func Isamax(N int, X []float, incX int) int {
	c_X := (*C.float)(unsafe.Pointer(&X[0]))
	return int(C.cblas_isamax(C.int(N), c_X, C.int(incX)))
}
