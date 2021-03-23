<?php

$a = [
   ['month' => 4, 'day' => 5],
   ['month' => 5, 'day' => 4]
];

$f = fn ($m, $n) => $m['day'] <=> $n['day'];
usort($a, $f);
print_r($a);
