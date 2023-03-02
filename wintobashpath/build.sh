go build -o ./bin/ . && upx --best --lzma ./bin/wintobashpath.exe
# go build -ldflags "-s -w -extldflags" -o ./bin/ . # && upx --best --lzma ./bin/wintobashpath.exe