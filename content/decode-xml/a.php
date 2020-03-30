<?php
# example 1
$s1 = <<<eof
<Sunday ten="even">
   <Monday eleven="odd"/>
</Sunday>
eof;
$o1 = simplexml_load_string($s1);
# example 2
$o2 = simplexml_load_file('a.xml');
# example 3
$o3 = $o1->Monday;
# example 4
$o4 = $o1['ten'];
# example 5
$s2 = (string) $o1['ten'];
# print
var_dump($o1, $o2, $o3, $o4, $s2);
