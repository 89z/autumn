# Talk

## Convert to integer

One of these:

~~~
content/convert-to-integer/talk.md
content/convert-to-integer/talk/index.md
~~~

## Two way conditions

One of these:

~~~
content/two-way-conditions/talk.md
content/two-way-conditions/talk/index.md
~~~

## Web server

One of these:

~~~
content/web-server/talk.md
content/web-server/talk/index.md
~~~

## Input output

Normally we would put this here:

~~~
content/categories/input-output/talk.md
~~~

However Hugo has a strange bug, where if you put too many files in a taxonomy
folder, it converts to Kind section. This kills the normal use of that location,
which is for links to pages with that taxonomy. So we much put talk pages
somewhere else. It doesnt really matter where, so I chose this:

~~~
content/category/input-output/talk.md
~~~

Make sure to add front matter to the talk page:

~~~
title: 'Talk:Input output'
categories: [input-output]
~~~

## Type

~~~
content/category/type/talk.md
~~~

## Go

~~~
content/tag/go/talk.md
~~~

## C

~~~
content/tag/c/talk.md
~~~

## Languages

Languages is not a tag, so it will not appear on any tag page. But it is a
section, so it should appear in the page listing.

~~~
content/tag/talk.md
~~~
