<?php
$n = 1;

# example 1
if ($n == 3) {
   echo "three\n";
} else if ($n == 2) {
   echo "two\n";
} else {
   echo "else\n";
}

# example 2
if ($n == 3) {
   echo "three\n";
} elseif ($n == 2) {
   echo "two\n";
} else {
   echo "else\n";
}

# example 3
$n == 2 ? print "two\n" : print "else\n";
