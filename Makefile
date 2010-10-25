all:
	6g libgeo.go
	6g example.go
	6l -o example example.6
	rm libgeo.6 example.6

clean:
	rm example
