+++
title = "Go"
tags = [ "go" ]
+++

size      | command
----------|--------
2,093,568 | `go build`
2,046,976 | `go build -gcflags all=-B`
1,925,120 | `go build -gcflags all=-l`
1,885,696 | `go build -gcflags 'all=-B -l'`
1,495,040 | `go build -ldflags -s`
1,461,248 | `go build -ldflags -s -gcflags all=-B`
1,420,800 | `go build -ldflags -s -gcflags all=-l`
1,390,080 | `go build -ldflags -s -gcflags 'all=-B -l'`

<https://golang.org/cmd/go#hdr-Compile_packages_and_dependencies>

{{< r "install.ps1" >}}

<https://golang.org/cmd/go#hdr-Compile_and_install_packages_and_dependencies>

{{< r "list.ps1" >}}

<https://golang.org/cmd/go#hdr-List_packages_or_modules>

{{< r "mod.ps1" >}}

<https://golang.org/cmd/go#hdr-Module_maintenance>

{{< r "run.ps1" >}}

<https://golang.org/cmd/go#hdr-Compile_and_run_Go_program>

{{< r "test.ps1" >}}

<https://golang.org/cmd/go#hdr-Test_packages>

## References

- <https://hyperpolyglot.org/c#hello-world>
- <https://rosettacode.org/wiki/Hello_world/Newbie>
