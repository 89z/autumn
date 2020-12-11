# example 1
go build a.go
# example 2
go build
# example 3
go build -v
# example 4
go build -ldflags -s
# example 5
go build -gcflags=all='-B -l' -ldflags=-s
