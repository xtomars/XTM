# generate wallet golang code
protoc --go_out=plugins=grpc:. -I $GOPATH/src:. wallet.proto
# # generate wallet java code 
# protoc --java_out=./ exchange.proto
# # package java code to jar
# jar cvf Exchange.jar io/bhex/base/wallet/Exchange.java
# # remove java package
# rm -rf io