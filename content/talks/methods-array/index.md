---
title: 'Methods: array'
---

## Add to array

- <https://golang.org/pkg/builtin#append>
- {{< a `https://developer.mozilla.org/Web/JavaScript/Reference/
   Global_Objects/Array/push` >}}

## Array contains index

<https://stackoverflow.com/questions/27252152>

## Array contains value

See [get index from value](#get-index-from-value).

## Get array index from value

Array search is usually a footgun. Using this input:

{{< r "a.php" >}}

Here is a naive example:

{{< r "b.php" >}}

Better to create map from the keys:

{{< r "c.php" >}}

Best to create map from the keys and values:

{{< r "d.php" >}}

Note the difference will be more pronounced on language like Go, as it does
not have an search function. You would need to iterate the header each time
you read another row.

- <http://hyperpolyglot.org/scripting#array-element-index>
- <https://rosettacode.org/wiki/Search_a_list>

## Iterate array

- <https://golang.org/ref/spec#For_range>
- {{< a `https://developer.mozilla.org/Web/JavaScript/Reference/
   Global_Objects/Array/forEach` >}}

## Length of array

- <https://golang.org/pkg/builtin#len>
- {{< a `https://developer.mozilla.org/Web/JavaScript/Reference/
   Global_Objects/Array/length` >}}

## Slice array

- <https://php.net/function.array-slice>
- {{< a `https://developer.mozilla.org/Web/JavaScript/Reference/
   Global_Objects/Array/slice` >}}

## References

- <https://developer.mozilla.org/Web/JavaScript/Reference/Global_Objects/Array>
- <https://golang.org/ref/spec#Slice_types>
