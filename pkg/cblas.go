// simplistic wrapper around the C-interface to the BLAS
package cblas

/*
 #include "cblas.h"
*/
import "C"

/*
 * Enumerated and derived types
 */

type Index int

type Order C.enum_CBLAS_ORDER
const (
	RowMajor = Order(101)
	ColMajor = Order(102)
)

type Transpose C.enum_CBLAS_TRANSPOSE
const (
	NoTrans   = Transpose(111)
	Trans     = Transpose(112)
	ConjTrans = Transpose(113)
)

type UpLo C.enum_CBLAS_UPLO
const (
	Upper = UpLo(121)
	Lower = UpLo(122)
)

type Diag C.enum_CBLAS_DIAG
const (
	NonUnit = Diag(131)
	Unit    = Diag(132)
	)

type Side C.enum_CBLAS_SIDE
const (
	Left = Side(141)
	Right= Side(142)
)

// EOF
