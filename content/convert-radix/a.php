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

$n = 1577858399;
$o = new Radix64;
// example 1
$s1 = $o->encode($n);
$n1 = $o->decode($s1);
// example 2
$s2 = $o->encode($n - 1);
$n2 = $o->decode($s2);
// print
var_dump($s1 == '0T22KU', $n1 == $n, $s2 == '0T22KT', $n2 == $n - 1);
