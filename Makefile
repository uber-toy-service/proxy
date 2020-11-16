clean:
	rm main
all:
	go build main.go && ./main
