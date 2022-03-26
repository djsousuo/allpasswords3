SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -v -a -o release/allpasswords_windows_amd64.exe


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build  -v -a -o release/allpasswords_linux_amd64

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
go build  -v -a -o release/allpasswords_linux_arm64

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -v -a -o release/allpasswords_darwin_amd64

cd ./release/

C:/zip/zip.exe -r ./allpasswords_windows_amd64.zip ./allpasswords_windows_amd64.exe
tar -zcvf allpasswords_linux_amd64.tar.gz allpasswords_linux_amd64
tar -zcvf allpasswords_darwin_amd64.tar.gz allpasswords_darwin_amd64
tar -zcvf allpasswords_linux_arm64.tar.gz allpasswords_linux_arm64