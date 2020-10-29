<?php

class Time {
   public $hours = 23;
   var $minutes = 59;
   function duration($seconds) {
      return $this->hours * 60 * 60 + $this->minutes * 60 + $seconds;
   }
}

$o = new Time;
$n = $o->duration(59);
var_dump($n == 86399);
