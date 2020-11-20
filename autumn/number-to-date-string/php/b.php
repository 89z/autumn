<?php

function timetostr($n) {
   $date_o = date_create();
   $zone_o = timezone_open('America/Chicago');
   date_timezone_set($date_o, $zone_o);
   date_timestamp_set($date_o, $n);
   return date_format($date_o, DATE_W3C);
}

$s = timetostr(1577858399);
var_dump($s == '2019-12-31T23:59:59-06:00');
