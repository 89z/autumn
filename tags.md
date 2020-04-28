# Tags

~~~
content/tags/c/index.md
content/tags/c/talk/index.md
content/tags/go/index.md
content/tags/go/talk/index.md
content/tags/index.md
content/tags/talk/index.md
~~~

C is a draft language, so it will not have a listing of articles:

~~~
content/tags/c/index.md
content/tags/go/index.md
content/tags/go/talk/index.md
content/tags/index.md
content/tags/talk/index.md
~~~

Need to split:

~~~
content/tags/index.md
content/tags/go/index.md
~~~

~~~
content/tags/c/index.md
content/tags/go/talk/index.md
content/tags/talk/index.md
~~~

Need to rename the second group:

~~~
content/tags/index.md
content/tags/go/index.md
~~~

~~~
content/languages/index.md
content/languages/go/index.md
content/languages/c/index.md
~~~

We dont have files for the first group, those are automatically generated.
However we will have 2 sets of links:

~~~
/autumn/tags
/autumn/tags/go
~~~

~~~
/autumn/languages
/autumn/languages/go
/autumn/languages/c
~~~

I dont like having two links for Go, or do I? The alternative is having Go links
and Talk together, and that is worse. The last issue is, we need to be able to
easily find all these. These are good, can just put links in the site header:

~~~
/autumn/tags
/autumn/tags/go
~~~

but what about these?

~~~
/autumn/languages
/autumn/languages/go
/autumn/languages/c
~~~

## References

<https://en.wikipedia.org/wiki/Sara_Del_Valle>
