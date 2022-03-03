<?php
$arr = [2,5,2,1,2,6,457,2345,2345,5,16,67,23445,2345,234];

//通过计数器进行统计出现的次数再进行过滤
function showOne (array $arr){
    $newArr = [];
    for ($i = 0 ;$i < count($arr) ; $i++) {
        $newArr[$arr[$i]] = isset($newArr[$arr[$i]]) ? 2 : 1;
    }
    return array_keys(array_filter($newArr,function ($num) {
        if ($num == 1) {
            return $num;
        }
    }));
}

//先排序，再挨个数字找
function showOne1(array $arr){
    sort($arr);
    for ($i = 0 ;$i < count($arr) - 1 ; $i++) {
        $c = $arr[$i];
        for ($j = $i+1 ;$j < count($arr); $j++) {
            if ($c == $arr[$j]) {
                unset($arr[$i]);
                unset($arr[$j]);
                $i = $j;
            } else {
                continue;
            }
        }     
    }
    return $arr;
}

var_dump(showOne1($arr));