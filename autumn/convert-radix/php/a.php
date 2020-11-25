<?php

function encode36(int $n): string {
   $s = (string) $n;
   return base_convert($s, 10, 36);
}

function decode36(string $s): int {
   return (int) base_convert($s, 36, 10);
}

$n = 1577858399;
$s = encode36($n);
$n2 = decode36($s);
var_dump($n2 === $n);
