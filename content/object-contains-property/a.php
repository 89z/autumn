<?php
$o1 = new class{
   public $sun = 10;
};
# example 1
$b1 = property_exists($o1, 'sun');
# example 2
$b2 = array_key_exists('sun', $o1);
# print
var_dump($b1, $b2);
