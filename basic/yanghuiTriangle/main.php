<?php

$numRows = 5;

function yanghuiTriangle($numRows){

    $arr = [];
    for ($i = 1;$i<=$numRows;$i++) {
        $arr[$i][1] = 1; 
        for ($j=2;$j<$i;$j++) {
            $arr[$i][$j] = $arr[$i - 1][$j - 1] + $arr[$i - 1][$j];
        }
        $arr[$i][$i] = 1;
    }
    return $arr;
}

var_dump(yanghuiTriangle($numRows));