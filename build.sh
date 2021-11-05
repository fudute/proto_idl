# !/bin/sh -e
IDL=$1

if [ $# -eq 0 ]; then
    echo please provide idl
    exit -1
fi

REPO_NAME=$(basename ${IDL%\.proto})

protoc --go_out=$HOME/go/src --go-grpc_out=$HOME/go/src $IDL

pushd ../$REPO_NAME

go mod init && go mod tidy

git init
gh repo create $REPO_NAME --public -y
git add *
git commit "first commit"
git push --set-upstream origin master

popd 
