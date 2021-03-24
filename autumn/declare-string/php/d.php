<?php

# slash
$s = 'south\north';
var_dump($s);

# quote
$s = 'south"north';
var_dump($s);
$s = 'south\'north';
var_dump($s);

# newline
$s = 'south
north';
var_dump($s);
