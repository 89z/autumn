<?php
$s = "\x43";

$n = match ($s) {
   'A' => 0x41,
   'B', 'b' => 0x42,
   default => ord($s)
};

var_dump($n == 0x43);
