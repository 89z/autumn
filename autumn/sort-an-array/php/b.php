<?php

$a = [
   ['month' => 10, 'day' => 31], ['month' => 11, 'day' => 30]
];

$f = fn ($m, $n) => $m['day'] <=> $n['day'];
usort($a, $f);
print_r($a);
