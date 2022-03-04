<?php
$arr = [2,5,2,1,2,6,457,2345,2345,5,16,67,23445,2345,234];

//数组排重-内置函数去重
function arrUnique(array $arr) {
    return array_unique($arr);
}

//数组排重-先排序，再依次对比删除
function arrUnique1(array $arr) {
    sort($arr);
    $j = 0;
    $len = count($arr);
    for($i = 1;$i < $len;$i++){
        if ($arr[$j] == $arr[$i]) {
            unset($arr[$i]);
        } else{
            $j = $i;
        }
    }
    return $arr;
}

//数组排重-放在一个新数组装起来
function arrUnique2(array $arr){
    $newArr = [];
    foreach($arr as $num)  {
        if (!in_array($num,$newArr)) {
            $newArr[] = $num;
        }
    }
    return $newArr;
}

var_dump(arrUnique1($arr));