<?php

function timetostr($n) {
   $date_o = new DateTime;
   $zone_o = new DateTimeZone('America/Chicago');
   $date_o->setTimezone($zone_o);
   $date_o->setTimestamp($n);
   return $date_o->format(DATE_W3C);
}

$s = timetostr(1577858399);
var_dump($s == '2019-12-31T23:59:59-06:00');
