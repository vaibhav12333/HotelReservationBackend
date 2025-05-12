.PHONY: build
build:
	go build HotelReservationBackend

.PHONY: run
run:
	./HotelReservationBackend

.PHONY: fresh
fresh:
	go build HotelReservationBackend
	./HotelReservationBackend



