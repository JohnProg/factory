all: factory

rebuild: clean factory

factory: deps
	go build -o factory ./src

deps:
	go get github.com/gorilla/mux
	touch deps

clean:
	rm -f factory deps