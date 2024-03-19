โดยไป download protoc ตาม os ของเราจาก https://github.com/protocolbuffers/protobuf/releases ได้เลย จากนั้น extract zip ออกมาไว้สักที่จะได้ content ของ zip ได้แก่ bin, include และ readme.txt โดยให้เราเอา path ของ bin ไปใส่ไว้ใน $PATH variable

หลังจากนั้นลองพิมพ์ command

$ protoc --version

ถ้าแสดงเลข version ออกมาก็เป็นอันเรียบร้อย

จากนั้นเราก็ลง plugin ภาษา Go ให้กับ protoc โดย execute command นี้

go get -u github.com/golang/protobuf/protoc-gen-go

Go get package in project

go get google.golang.org/protobuf/cmd/protoc-gen-go go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

Install gRPC tool in project go install google.golang.org/protobuf/cmd/protoc-gen-go go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

//gen protoc calculator.proto --go_out=../server --go-grpc_out=../server

//same postman for gRPC https://github.com/ktr0731/evans
