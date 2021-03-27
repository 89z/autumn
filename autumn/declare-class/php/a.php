<?php

class date {
   public $month = 12;
   public $day = 31;
   function year() {
      $this->month = 1;
      $this->day = 1;
   }
}

$d = new date;
$d->year();
var_dump($d->day == 1);
