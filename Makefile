build:
	go build -o tmp/gomodoro main.go

run: build
	./tmp/gomodoro

clean:
	rm -rf ./tmp/gomodoro
