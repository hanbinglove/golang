<?php
$arr = [2,5,2,2,6,457,2,2345,2345,5,16,67,23445,1,2345,234];

//依次比较找出最小值 和 最大值与最小值的差值
function stockTrading(array $arr){
    if (empty($arr)) return 0;
    $min = $arr[0];
    $max = 0;
    for ($i = 1;$i<count($arr);$i++) {
        if ($arr[$i] < $min) {
            $max = 0;
            $min = $arr[$i];
        } else {
            $max = max($max,$arr[$i] - $min);
        }
    }

    return $max;
}

var_dump(stockTrading($arr));