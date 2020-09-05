<?php
class Day {
   public $Sunday = 10;
}
$o = new Day;
$b = property_exists($o, 'Sunday');
var_dump($b);
