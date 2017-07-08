echo go-torch
go get github.com/uber/go-torch
# git clone git@github.com:brendangregg/FlameGraph.git
go-torch --binaryname nest-client-go.test -b prof.cpu
open torch.svg 
