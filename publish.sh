git checkout master
HASH=$(git rev-parse HEAD)
git checkout -b "rebuild_$HASH"

env GOOS=windows GOARCH=amd64 go build -o compiled/pubsub-loadgen_amd64.exe
env GOOS=windows GOARCH=386 go build -o compiled/pubsub-loadgen_x86.exe
env GOOS=linux GOARCH=amd64 go build -o compiled/pubsub-loadgen_amd64_linux
env GOOS=linux GOARCH=386 go build -o compiled/pubsub-loadgen_x86_linux
env GOOS=darwin GOARCH=amd64 go build -o compiled/pubsub-loadgen_amd64_mac
env GOOS=darwin GOARCH=386 go build -o compiled/pubsub-loadgen_x86_mac

git add .

git commit -am "Updated Compiled Source with master hash:$HASH"
git push



