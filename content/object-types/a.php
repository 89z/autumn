<?php
# example 1
class jan {
   public $sun = 10;
   public $mon = 11;
}
$o1 = new jan;
# example 2
$o2 = new class{
   public $sun = 10;
   public $mon = 11;
};
# print
var_dump($o1, $o2);
