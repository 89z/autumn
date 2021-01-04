<?php

function seconds(object $t, object $u): int {
   $d = $u->diff($t);
   $t->setTimestamp(0);
   $t->add($d);
   return $t->getTimestamp();
}

$o = DateTime::createFromFormat('!Y-m-d', '2020-05-04');
$n = seconds(new DateTime, $o);
var_dump($n);
