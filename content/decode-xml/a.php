<?php
# example 1
$o1 = simplexml_load_file('a.xml');
# example 2
$s1 = <<<eof
<Sunday ten="even">
   <Monday eleven="odd"/>
</Sunday>
eof;
$o2 = simplexml_load_string($s1);
# print
var_dump($o1, $o2);
