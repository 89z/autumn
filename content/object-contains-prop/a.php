<?php
$o = new class {
   public $Sunday = 10;
};
# example 1
$b1 = property_exists($o, 'Sunday');
# example 2
$b2 = array_key_exists('Sunday', $o);
# print
var_dump($b1, $b2);
