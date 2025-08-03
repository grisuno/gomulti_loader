linux:
	GOOS=linux GOARCH=amd64 garble -literals -tiny build -ldflags="-s -w"  -o loader_linux ; upx loader_linux

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc garble -literals -tiny build -ldflags="-s -w"  -o loader_windows.exe ; upx loader_windows.exe

clean:
	rm -f loader_linux loader_windows.exe
