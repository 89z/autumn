<?php
declare(strict_types = 1);

echo bin2hex("\xa\xa");
# 0a0a

print_r(unpack('H*', "\xa\xa"));
/*
Array
(
    [1] => 0a0a
)
*/
