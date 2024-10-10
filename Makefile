build:
	go build -o tmp/gomodoro main.go

run: build
	./tmp/gomodoro start

clean:
	rm -rf ./tmp/gomodoro
