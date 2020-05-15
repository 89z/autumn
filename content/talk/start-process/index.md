---
title: 'Talk:Start process'
---

The current PHP options are garbage for Windows. Example 1:

{{< r "a.php" >}}

Example 2:

{{< r "b.php" >}}

The second example seems promising, but if you try this:

{{< r "c.php" >}}

The quote is stripped and you get here:

<https://example.com/sunmon>

The workaround is to split on double quotes, then join again with double quotes.
Interestingly, interleaving quotes do not need to be escaped:

<https://stackoverflow.com/questions/7760545>

Implementation:

{{< r "d.php" >}}
