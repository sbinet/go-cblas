package main

import "fmt"
import "github.com/sbinet/go-cblas/pkg/cblas"

func main() {
	{
		x := []float32{1., 2., 3.}
		fmt.Printf("x=%v\n", x)
		sumx := cblas.Sasum(x)
		fmt.Printf("sasum(x)=%v\n", sumx)
	}

	{
		x := []float64{1., 2., 3.}
		fmt.Printf("x=%v\n", x)
		sumx := cblas.Dasum(x)
		fmt.Printf("dasum(x)=%v\n", sumx)
	}

	{
		x := []complex64{complex(1., 1.),
			complex(2., 2.),
			complex(3., 3.)}
		fmt.Printf("x=%v\n", x)
		sumx := cblas.Scasum(x)
		fmt.Printf("scasum(x)=%v\n", sumx)
	}

	{
		x := []complex128{complex(1., 1.),
			complex(2., 2.),
			complex(3., 3.)}
		fmt.Printf("x=%v\n", x)
		sumx := cblas.Dzasum(x)
		fmt.Printf("dzasum(x)=%v\n", sumx)
	}

	{
		x := []float32{0.733}
		y := []float32{0.825}
		alpha := float32(0.)

		exp := float32(0.604725)
		val := cblas.Sdsdot(alpha, x, y)
		fmt.Printf("sdsdot: exp: %v\n", exp)
		fmt.Printf("sdsdot: val: %v\n", val)
	}
	{
		x := []complex128{complex(-0.87, -0.631)}
		y := []complex128{complex(-0.7, -0.224)}

		exp := complex128(complex(0.467656, 0.63658))
		val := cblas.Zdotu(x, y)

		fmt.Printf("zdotu: exp: %v\n", exp)
		fmt.Printf("zdotu: val: %v\n", val)
	}
	return
}

/*
     #include <stdio.h>

     int
     main (void)
     {
       int lda = 3;

       float A[] = { 0.11, 0.12, 0.13,
                     0.21, 0.22, 0.23 };

       int ldb = 2;

       float B[] = { 1011, 1012,
                     1021, 1022,
                     1031, 1032 };

       int ldc = 2;

       float C[] = { 0.00, 0.00,
                     0.00, 0.00 };

       // Compute C = A B 

       cblas_sgemm (CblasRowMajor, 
                    CblasNoTrans, CblasNoTrans, 2, 2, 3,
                    1.0, A, lda, B, ldb, 0.0, C, ldc);

       printf ("[ %g, %g\n", C[0], C[1]);
       printf ("  %g, %g ]\n", C[2], C[3]);

       return 0;  
     }

 // expected output:
 // [ 367.76, 368.12
 //   674.06, 674.72 ]

*/
