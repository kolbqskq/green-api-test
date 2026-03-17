
test:
	go test -v ./tests

test-no-cache:
	go test -v -count=1 ./tests
