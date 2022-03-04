<?php
$nums1 = [1,2,3,0,0,0];
$m = 3;

$nums2 = [2,5,6];
$n = 3;

function merge ($nums1,$nums2,$m,$n){
    array_splice($nums1,$m);
    array_splice($nums2,$n);
    $nums1 = array_merge($nums1,$nums2);
    sort($nums1);
    return $nums1;
}

var_dump(merge($nums1,$nums2,$m,$n));