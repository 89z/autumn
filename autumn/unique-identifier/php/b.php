<?php

function id_decode(string $s, int $year): object {
   $t = new DateTime;
   $t->setDate($year, 1, 1)->setTime(0, 0);
   $x = (int) base_convert($s, 36, 10) + $t->getTimestamp();
   return $t->setTimestamp($x);
}

$o = id_decode('itrzz', 2020);
var_dump($o);
