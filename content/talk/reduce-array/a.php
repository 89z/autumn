<?php
$a = ['May', 'June'];
$f = fn ($s_acc, $s_cur) => $s_acc . $s_cur;
$s = array_reduce($a, $f);
var_dump($s);
