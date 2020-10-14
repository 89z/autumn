<?php
$s = 'June';
$s2 = strpbrk($s, 'abcde');
var_dump($s2 !== false);
