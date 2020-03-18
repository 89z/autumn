<?php
$o1 = new class{
   public $Sun = 10;
};
# example 1
$b1 = property_exists($o1, 'Sun');
# example 2
$b2 = array_key_exists('Sun', $o1);
# print
var_dump($b1, $b2);
