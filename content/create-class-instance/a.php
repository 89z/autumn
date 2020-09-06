<?php
# example 1
$o1 = new stdClass;
$o1->Sunday = 10;
# example 2
class Day {
   public $Sunday = 10;
}
$o2 = new Day;
# example 3
$o3 = new class {
   public $Sunday = 10;
};
# print
var_dump($o1, $o2, $o3);
