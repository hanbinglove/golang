<?php
$array = [1,2,0,4,0,1,2,3,0];

//先进行数组排序，再进行依次对比
function checkRepeat(array $array){
    sort($array);
    for ($i = 0;$i<count($array) - 1;$i++) {
        if ($array[$i] == $array[$i+1]){
            return true;
        }
    }
    return false;
}

//使用哈希表记录元素是否出现过
function checkRepeat1(array $array){
    $newArray = [];
    for ($i = 0;$i<count($array);$i++) {
        if (in_array($array[$i],$newArray)){
            return true;
        }
        $newArray[] = $array[$i];
    }
    return false;
}

var_dump(checkRepeat($array));
var_dump(checkRepeat1($array));