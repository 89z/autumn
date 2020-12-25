<?php
extension_loaded('mbstring') or die('mbstring');

function fold(string $s, string $t): bool {
   $n = MB_CASE_FOLD;
   return mb_convert_case($s, $n) == mb_convert_case($t, $n);
}

$b1 = fold("\u{006B}", "\u{212A}"); # LATIN SMALL LETTER K, KELVIN SIGN
$b2 = fold("\u{00B5}", "\u{03BC}"); # MICRO SIGN, GREEK SMALL LETTER MU
var_dump($b1, $b2);
