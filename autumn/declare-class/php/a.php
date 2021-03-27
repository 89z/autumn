<?php

class date {
   public $month = 1;
   public $day = 1;
   function add() {
      $this->day++;
   }
}

$d = new date;
$d->add();
var_dump($d->day == 2);
