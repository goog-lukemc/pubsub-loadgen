env GOOS=windows GOARCH=amd64 go build -o pubsub-loadgen_amd64.exe
env GOOS=windows GOARCH=386 go build -o pubsub-loadgen_x86.exe
env GOOS=linux GOARCH=amd64 go build -o pubsub-loadgen_amd64_linux
env GOOS=linux GOARCH=386 go build -o pubsub-loadgen_x86_linux
env GOOS=darwin GOARCH=amd64 go build -o pubsub-loadgen_amd64_mac
env GOOS=darwin GOARCH=386 go build -o pubsub-loadgen_x86_mac
git commit add .
git commit -am "publish build"
git push master




