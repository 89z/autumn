# Talk

## Convert to integer

~~~
content/convert-to-integer/index.md
content/convert-to-integer/talk/index.md
~~~

## Two way conditions

~~~
content/two-way-conditions/index.md
content/two-way-conditions/talk/index.md
~~~

## Web server

~~~
content/web-server/index.md
content/web-server/talk/index.md
~~~

## Input output

The base file would be here:

~~~
content/categories/input-output/index.md
~~~

but this is automatically created, so we can disregard. We could use that file
for the talk page, but then we have the links and the talk page all together.
What I want is the links, along with a link to the Talk page at the top. Is that
possible? If so how? The talk page will be here:

~~~
content/categories/input-output/talk/index.md
~~~

I cant seem to get Hugo to recognize the page with my current configuration. It
seems this is the problem:

<https://discourse.gohugo.io/t/adding-content-for-taxonomies/6343/4>

and this possibly the solution:

<https://gohugo.io/templates/lookup-order>

The page in question is kind `taxonomy`, here are the options:

~~~
layouts/_default/category.html
layouts/_default/list.html
layouts/_default/taxonomy.html
~~~

## Rest

~~~
languages/go
languages/talk
type/talk
~~~
