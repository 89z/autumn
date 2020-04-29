# Hugo lookup order

The page in question is kind `taxonomy`, here are the options:

~~~
layouts/_default/category.html
layouts/_default/list.html
layouts/_default/taxonomy.html
~~~

Now `list.html` is interesting, because it can also be used for `home`,
`section` and `term`. So I could do this:

~~~
baseof.html
index.html => remove
section.html => remove
single.html
taxonomy.html => remove
terms.html => list.html
~~~

<https://gohugo.io/templates/lookup-order>
