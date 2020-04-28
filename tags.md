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

Also what about this:

~~~
content/categories/input-output/index.md
~~~

Heres an idea:

What              | Link
------------------|---------------
List of languages | `/autumn/tags`
List of Go pages  | `/autumn/tags/go`
List of topics    | `/autumn/categories`
List of IO pages  | `/autumn/categories/input-output`
New language talk | `/autumn/tag`
Go language talk  | `/autumn/tag/go`
New topic talk    | `/autumn/category`
IO topic talk     | `/autumn/cateogry/input-output`
