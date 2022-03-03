<?php
$array = [1,2,0,4,0,1,2,3,0];

//移动所有0到末尾，为0的数据插入末尾
function moveZeroes(array $array){
    for ($i = 0 ; $i < count($array) ; $i++) {
        if ($array[$i] == 0) {
            $array[] = $array[$i];
            unset($array[$i]);
        }
    }
    //数组重新排列
    return array_merge($array);
}

//移动所有0到末尾,不为0的数据依次前移
function moveZeroes1(array $array){
    for ($i = $j = 0; $i < count($array); $i++) {
        if ($array[$i] != 0) {
            $c = $array[$j];
            $array[$j] = $array[$i];
            $array[$i] = $c;
            $j++;
        }
    }
    return $array;
}


var_dump(moveZeroes($array));