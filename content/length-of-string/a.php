<?php
$s1 = '♠';
# bytes
$n1 = strlen($s1);
# characters
$n2 = mb_strlen($s1);
# characters
$n3 = iconv_strlen($s1);
# characters
$n4 = grapheme_strlen($s1);
# print
var_dump($n1, $n2, $n3, $n4);
