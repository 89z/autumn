<?php
# example 1
class Jan {
   public $Sun = 10;
   public $Mon = 11;
}
$o1 = new Jan;
# example 2
$o2 = new class{
   public $Sun = 10;
   public $Son = 11;
};
# print
var_dump($o1, $o2);
