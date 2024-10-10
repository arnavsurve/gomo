build:
	go build -o tmp/gomo main.go

run: build
	./tmp/gomo start

clean:
	rm -rf ./tmp/gomo
