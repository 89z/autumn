<?php

function pattern(string $pattern, string $subject): bool {
   return preg_match($pattern, $subject) == 1;
}

$b = pattern('/n./', 'north');
var_dump($b);
