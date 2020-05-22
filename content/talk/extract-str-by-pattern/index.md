---
title: 'Talk:Extract string by pattern'
---

If I have some code like this:

~~~jl
s1 = "package"
o3 = match(r"p(..)", s1)
a1 = o3.captures
println(a1)
~~~

I get this result:

~~~jl
Union{Nothing, SubString{String}}["ac"]
~~~

I was a little confused by all the decoration on the result. So I went to look
at the documentation, but then I realized there wasnt any. The only mention of
this `struct` is here:

> The matching substring can be retrieved by accessing `m.match` and the
> captured sequences can be retrieved by accessing `m.captures`

<https://docs.julialang.org/en/v1/base/strings#Base.match>

So I went searching and found the `struct` definition in the code:

<https://github.com/JuliaLang/julia/blob/40a8b42d45a274a0f91565693830360faa12081c/base/regex.jl#L134-L140>

- c++
- https://golang.org/pkg/regexp/#Regexp.FindStringSubmatch
- python
- rust
