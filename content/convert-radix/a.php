<?php

class Radix64 {
   public $s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';

   function encode(int $n_in): string {
      $s_out = '';
      do {
         $s_out = $this->s_dig[$n_in % 64] . $s_out;
         $n_in = intdiv($n_in, 64);
      } while ($n_in > 0);
      return $s_out;
   }

   function decode(string $s_in): int {
      $n_out = 0;
      $a_in = str_split($s_in);
      foreach ($a_in as $s_chr) {
         $n_out = $n_out * 64 + strpos($this->s_dig, $s_chr);
      }
      return $n_out;
   }
}

$n = time();
$o = new Radix64;
$s = $o->encode($n);
$n2 = $o->decode($s);
var_dump($n, $s, $n2 == $n);
