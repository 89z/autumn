<?php
$a = ['May', 'June'];
$m = [];
$f = fn ($m_acc, $s_cur, $n_idx) => $m_acc + [$s_cur => $n_idx];

foreach ($a as $n => $s) {
   $m = $f($m, $s, $n);
}

print_r($m);
