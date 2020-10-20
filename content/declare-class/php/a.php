<?php

class Time {
   var $hours = 23;
   function duration($minutes) {
      return $this->hours * 60 + $minutes;
   }
}

$o = new Time;
$n = $o->duration(59);
var_dump($n == 1439);
