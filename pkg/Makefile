# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=cblas

CGO_LDFLAGS=-lcblas

CGOFILES=\
	cblas_lvl1.go\
	cblas_lvl2.go\
	cblas_lvl3.go\
	cblas.go


include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
