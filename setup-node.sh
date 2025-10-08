#! /bin/bash

# setup a maelstrom node
echo "enter the node name (e.g. unique-ids):"
read name
proj="maelstrom-$name"
echo "new dir will be \"$proj\". correct? [y/n]"
read yn
if [[ $yn == "y" ]]; then
    echo "setting up new project \"$proj\""
    mkdir $proj
    cd $proj
    go mod init $proj
    go mod tidy
    cp ../setup/main-template.go ./main.go
    go get github.com/jepsen-io/maelstrom/demo/go
    go install .
else
    echo "exiting"
fi