# !/bin/sh -e

IDL=$1

if [ $# -eq 0 ]; then
    echo please provide idl
    exit -1
fi

REPO_NAME=$(basename ${IDL%\.proto})

if [ -d ../$REPO_NAME ] 
then
	is_new_repo=false
else
	is_new_repo=true
fi

protoc --go_out=$HOME/go/src --go-grpc_out=$HOME/go/src $IDL

pushd ../$REPO_NAME


new_dir(){
	go mod init && go mod tidy
	git init
	gh repo create $REPO_NAME --public -y
	git add *
	git commit -m "first commit"
	git push --set-upstream origin master

}

old_dir(){
	go mod tidy
	git add *
	git commit -m "$(date)"	
	git push
}

if [ $is_new_repo == true ] 
then
	echo create a new repo
	new_dir
else
	echo repo already exist
	old_dir
fi

popd 

