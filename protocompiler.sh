#!/bin/sh
PKG_OK=$(command -v protoc)
if [ "" = "$PKG_OK" ]; then
    echo "No protobuf. Setting up..."
    # Make sure you grab the latest version
    curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip
    unzip protoc-3.2.0-linux-x86_64.zip -d protoc3
      
    sudo mv protoc3/bin/* /usr/bin/
    sudo mv proto3/include/* /usr/include

    # clean up
    sudo rm -r proto3 protoc-3.2.0-linux-x86_64.zip
fi

for i in `find "$(cd services; pwd)" -name "*.proto"`; do
    echo "compiling protobuf..."
    protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src $i    
done

echo "done"