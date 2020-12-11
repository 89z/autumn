<?php

function seconds(object $new_o, object $old_o): float {
   $o = $new_o->diff($old_o);
   return $o->days * 86400 + $o->h * 3600 + $o->i * 60 + $o->s + $o->f;
}

$o = DateTime::createFromFormat('!Y-m-d', '2019-12-31');
$n = seconds(new DateTime, $o);
var_dump($n);
