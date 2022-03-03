## 基础篇

* 功能1：实现加减乘除,取余,自增自减


|             模块                    |  类目                   |                 练习名称   |    代码传送门   |   备注   |
| :------: | :-----------: | :---------------------: | :------: | :------: |
|     基础语法  |  操作  |   简单运算   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/operation/i.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/operation/i.php) | 在golang中的++和--并不是操作而是表达式，<br>所以不可以像php一样赋值，同时也没有--x和++x     |
|     基础语法  |  操作  |   移动数组为0的元素到末尾   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/moveZeroes/i.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/moveZeroes/i.php) | 1.php for里面取count ，就算实时改变了变量的内容，count也不会变，但是go会因为内容改变而改变len的值，导致死循环<br>2.golang的切片不能删除。。。php可以unset，go只能过滤当前元素，把其他元素重新追加一下<br>3.golang可以多变量赋值    |
