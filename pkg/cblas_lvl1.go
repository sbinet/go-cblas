// simplistic wrapper around the C-interface to the BLAS
package cblas

/*
 #include <complex.h>
 #undef I
 #include "cblas.h"
*/
import "C"
import "unsafe"

/*
 * ===========================================================================
 * Prototypes for level 1 BLAS functions (complex are recast as routines)
 * ===========================================================================
 */

/*
 float  cblas_sdsdot(const int N, const float alpha, 
 const float *X, const int incX, 
 const float *Y, const int incY);
*/
func Sdsdot(alpha float32, x, y []float32) float32 {
	if len(x) != len(y) {
		panic("slices' size differ")
	}

	c_N := C.int(len(x))
	c_alpha := C.float(alpha)

	c_X := (*C.float)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	c_Y := (*C.float)(unsafe.Pointer(&y[0]))
	c_incY := C.int(-1)

	return float32(
		C.cblas_sdsdot(c_N, c_alpha,
			c_X, c_incX,
			c_Y, c_incY))
}

/*
 double cblas_dsdot(const int N, const float *X, const int incX, const float *Y,
 const int incY);
*/
func Dsdot(x, y []float32) float64 {
	if len(x) != len(y) {
		panic("slices' size differ")
	}

	c_N := C.int(len(x))

	c_X := (*C.float)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	c_Y := (*C.float)(unsafe.Pointer(&y[0]))
	c_incY := C.int(-1)

	return float64(
		C.cblas_dsdot(c_N,
			c_X, c_incX,
			c_Y, c_incY))
}

/*
 float  cblas_sdot(const int N, const float  *X, const int incX,
 const float  *Y, const int incY);
*/
func Sdot(x, y []float32) float32 {
	if len(x) != len(y) {
		panic("slices' size differ")
	}

	c_N := C.int(len(x))

	c_X := (*C.float)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	c_Y := (*C.float)(unsafe.Pointer(&y[0]))
	c_incY := C.int(-1)

	return float32(
		C.cblas_sdot(c_N,
			c_X, c_incX,
			c_Y, c_incY))

}

/*
 double cblas_ddot(const int N, const double *X, const int incX,
 const double *Y, const int incY);
*/
func Ddot(x, y []float64) float64 {
	if len(x) != len(y) {
		panic("slices' size differ")
	}

	c_N := C.int(len(x))

	c_X := (*C.double)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	c_Y := (*C.double)(unsafe.Pointer(&y[0]))
	c_incY := C.int(-1)

	return float64(
		C.cblas_ddot(c_N,
			c_X, c_incX,
			c_Y, c_incY))
}

/*
 void   cblas_cdotu_sub(const int N, const void *X, const int incX,
 const void *Y, const int incY, void *dotu);
*/
func Cdotu(x, y []complex64) complex64 {

	if len(x) != len(y) {
		panic("slices' size differ")
	}

	var ret C.complexfloat
	c_ret := unsafe.Pointer(&ret)

	c_N := C.int(len(x))

	c_X := unsafe.Pointer(&x[0])
	c_incX := C.int(1)

	c_Y := unsafe.Pointer(&y[0])
	c_incY := C.int(-1)

	C.cblas_cdotu_sub(c_N,
		c_X, c_incX,
		c_Y, c_incY,
		c_ret)

	return complex(
		float32(C.crealf(ret)),
		float32(C.cimagf(ret)))
}

/*
 void   cblas_cdotc_sub(const int N, const void *X, const int incX,
 const void *Y, const int incY, void *dotc);
*/
func Cdotc(x, y []complex64) complex64 {

	if len(x) != len(y) {
		panic("slices' size differ")
	}

	var ret C.complexfloat
	c_ret := unsafe.Pointer(&ret)

	c_N := C.int(len(x))

	c_X := unsafe.Pointer(&x[0])
	c_incX := C.int(1)

	c_Y := unsafe.Pointer(&y[0])
	c_incY := C.int(-1)

	C.cblas_cdotc_sub(c_N,
		c_X, c_incX,
		c_Y, c_incY,
		c_ret)

	return complex(
		float32(C.crealf(ret)),
		float32(C.cimagf(ret)))
}

/*
 void   cblas_zdotu_sub(const int N, const void *X, const int incX,
                        const void *Y, const int incY, void *dotu);
*/
func Zdotu(x, y []complex128) complex128 {

	if len(x) != len(y) {
		panic("slices' size differ")
	}

	var ret C.complexdouble

	c_ret := unsafe.Pointer(&ret)

	c_N := C.int(len(x))

	c_X := unsafe.Pointer(&x[0])
	c_incX := C.int(1)

	c_Y := unsafe.Pointer(&y[0])
	c_incY := C.int(-1)

	C.cblas_zdotu_sub(c_N,
		c_X, c_incX,
		c_Y, c_incY,
		c_ret)

	return complex(
		float64(C.creal(ret)),
		float64(C.cimag(ret)))
}

/*
 void   cblas_zdotc_sub(const int N, const void *X, const int incX,
 const void *Y, const int incY, void *dotc);
 func Zdotc(x,y []complex64) complex64 {

 if len(x) != len(y) {
 panic("slices' size differ")
 }

 var ret C.complex64_t
 c_ret := (*C.char)(unsafe.Pointer(ret))

 c_N := C.int(len(x))

 c_X    := (*C.complex64_t)(unsafe.Pointer(&x[0]))
 c_incX := C.int(1)

 c_Y    := (*C.complex64_t)(unsafe.Pointer(&y[0]))
 c_incY := C.int(-1)

 C.cblas_cdotu_sub(c_N,
 c_X, c_incX,
 c_Y, c_incY,
 c_ret)

 return cmplx(
 float64(ret.real), 
 float64(ret.imag))
 }
*/

/*
 float  cblas_snrm2(const int N, const float *X, const int incX);
*/
func Snrm2(x []float32) float32 {

	c_N := C.int(len(x))

	c_X := (*C.float)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	return float32(C.cblas_snrm2(c_N, c_X, c_incX))
}

/*
 float  cblas_sasum(const int N, const float *X, const int incX);
*/
func Sasum(x []float32) float32 {

	c_N := C.int(len(x))

	c_X := (*C.float)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	return float32(C.cblas_sasum(c_N, c_X, c_incX))
}

/*
 double cblas_dnrm2(const int N, const double *X, const int incX);
*/
func Dnrm2(x []float64) float64 {
	c_N := C.int(len(x))

	c_X := (*C.double)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	return float64(C.cblas_dnrm2(c_N, c_X, c_incX))
}

/*
 double cblas_dasum(const int N, const double *X, const int incX);
*/
func Dasum(x []float64) float64 {
	c_N := C.int(len(x))

	c_X := (*C.double)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)

	return float64(C.cblas_dasum(c_N, c_X, c_incX))
}

/*
 float  cblas_scnrm2(const int N, const void *X, const int incX);
 func Snrm2(x []complex) float {

 c_N := C.int(len(x))

 c_X    := (*C.float)(unsafe.Pointer(&x[0]))
 c_incX := C.int(1)

 return float(C.cblas_scnrm2(c_N, c_X, c_incX))
 }
*/

/*
 float  cblas_scasum(const int N, const void *X, const int incX);
*/
func Scasum(x []complex64) float32 {

	c_N := C.int(len(x))

	c_X := unsafe.Pointer(&x[0])
	c_incX := C.int(1)

	return float32(C.cblas_scasum(c_N, c_X, c_incX))
}

/*
 double  cblas_dznrm2(const int N, const void *X, const int incX);
*/
func Dznrm2(x []complex64) float64 {

	c_N := C.int(len(x))

	c_X := unsafe.Pointer(&x[0])
	c_incX := C.int(1)

	return float64(C.cblas_dznrm2(c_N, c_X, c_incX))
}

/*
 double  cblas_dzasum(const int N, const void *X, const int incX);
*/
func Dzasum(x []complex128) float64 {

	c_N := C.int(len(x))

	c_X := unsafe.Pointer(&x[0])
	c_incX := C.int(1)

	return float64(C.cblas_dzasum(c_N, c_X, c_incX))
}


// functions having standard 4 prefixes (S D C Z)

/*
 CBLAS_INDEX cblas_isamax(const int N, const float  *X, const int incX);
*/
func Isamax(x []float32) Index {
	c_N := C.int(len(x))
	c_X := (*C.float)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)
	return Index(C.cblas_isamax(c_N, c_X, c_incX))
}

/*
 CBLAS_INDEX cblas_idamax(const int N, const double *X, const int incX);
*/
func Idamax(x []float64) Index {
	c_N := C.int(len(x))
	c_X := (*C.double)(unsafe.Pointer(&x[0]))
	c_incX := C.int(1)
	return Index(C.cblas_idamax(c_N, c_X, c_incX))
}
