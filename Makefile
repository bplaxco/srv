
all: srv

srv:
	go build srv.go

clean:
	rm -rf ./srv
