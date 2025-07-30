linux:
	GOOS=linux GOARCH=amd64 go build -o loader_linux

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o loader_windows.exe

clean:
	rm -f loader_linux loader_windows.exe
