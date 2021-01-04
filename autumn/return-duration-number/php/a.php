<?php

function since_hours(string $s): int {
   $t = new DateTime($s);
   $d = $t->diff(new DateTime);
   return $d->days * 24 + $d->h;
}

echo since_hours('2020-12-31'), "\n";
