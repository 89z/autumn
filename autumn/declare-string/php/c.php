<?php

# slash
$s = <<<'eof'
south\north
eof;
var_dump($s);

# quote
$s = <<<'eof'
south"north
south'north
eof;
var_dump($s);

# newline
$s = <<<'eof'
south
north
eof;
var_dump($s);
