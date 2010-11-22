include $(GOROOT)/src/Make.inc

TARG=libgeo
GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
	libgeo.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} -w libgeo.go
	${GOFMT} -w example.go
