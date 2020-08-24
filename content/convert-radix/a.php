<?php

class Radix {
   function __construct(string $s_in) {
      $this->s_dig = $s_in;
      $this->n_len = strlen($s_in);
   }

   function encode(int $n_in): string {
      $s_out = '';
      do {
         $s_out = $this->s_dig[$n_in % $this->n_len] . $s_out;
         $n_in = intdiv($n_in, $this->n_len);
      } while ($n_in > 0);
      return $s_out;
   }

   function decode(string $s_in): int {
      $n_out = 0;
      $a_in = str_split($s_in);
      foreach ($a_in as $s_chr) {
         $n_out = $n_out * $this->n_len + strpos($this->s_dig, $s_chr);
      }
      return $n_out;
   }
}

class Radix64 extends Radix {
   function __construct() {
      $this->s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';
      $this->n_len = 64;
   }
}

$n = time();
$o = new Radix64;
$s = $o->encode($n);
$n2 = $o->decode($s);
var_dump($n, $s, $n2 == $n);
