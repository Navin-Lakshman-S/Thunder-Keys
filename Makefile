.PHONY: all linux windows mac

OUT = Thunder-Keys
FILE = TypingSpeedTest.go

all: linux windows mac

linux:
<tab>GOOS=linux GOARCH=amd64 go build -o $(OUT)-linux $(FILE)

windows:
<tab>GOOS=windows GOARCH=amd64 go build -o $(OUT)-win.exe $(FILE)

mac:
<tab>GOOS=darwin GOARCH=arm64 go build -o $(OUT)-mac $(FILE)

clean:
<tab>rm -f $(OUT)-linux $(OUT)-win.exe $(OUT)-mac
