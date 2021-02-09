package main
//go:generate mkwinsyscall -output zexec.go exec.go
//sys GetBinaryType(name string, binaryType *int) (err error) = kernel32.GetBinaryTypeW
