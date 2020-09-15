<?php
$a_src = ['May', 'June'];
$f_red = fn ($s_acc, $s_cur) => $s_acc . $s_cur;
$s_acc = array_reduce($a_src, $f_red);
var_dump($s_acc);
