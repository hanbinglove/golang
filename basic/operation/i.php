<?php 
$a = 10;
$b = 20;

echo $a + $b . "\r\n";
echo $a - $b . "\r\n";
echo $a * $b . "\r\n";
echo $a / $b . "\r\n";
echo $a % $b . "\r\n";

$c = $a++;
$d = $a--;
$e = --$a;
$f = ++$a;
var_dump($c,$d,$e,$f);