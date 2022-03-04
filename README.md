## 基础篇

* 功能

|             模块                    |  类目                   |                 练习名称   |       代码传送门       |   备注   |
| :------: | :-----------: | :---------------------: | :------: | :------: |
|     基础语法  |  操作  |   简单运算   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/operation/i.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/operation/i.php) | 在golang中的++和--并不是操作而是表达式，<br>所以不可以像php一样赋值，同时也没有--x和++x     |
|     基础语法  |  操作  |   移动数组为0的元素到末尾   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/moveZeroes/main.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/moveZeroes/main.php) | 1.(将0的元素删除并插入末尾) : php 删除元素不会改变下标位置，golang会因为元素变化自动位移下标，不注意可能就会产生死循环<br>2(不为0的数据依次前移).golang可以多变量赋值    |
|     基础语法  |  操作  |   判断数组是否有重复元素   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/checkRepeat/main.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/checkRepeat/main.php) | 1.(先进行数组排序，再进行依次对比) : 无差别，<br>2.(使用哈希表记录元素是否出现过) : php的array_key_exists在go可以使用ok来判断,php的in_array在go没有，需要自己循环    |
|     基础语法  |  操作  |   删除数组里有重复内容的元素   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/showOne/main.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/showOne/main.php) | 1.(通过计数器进行统计出现的次数再进行过滤) : 实现方式基本相同<br>2.(先排序，再挨个数字找) :  golang没有unset,需要自己位移下标   |
|     基础语法  |  操作  |   最佳时机买卖股票   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/stockTrading/main.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/stockTrading/main.php) | 依次比较找出最小值 和 最大值与最小值的差值，go只对flot类型提供了内置的max和min函数，int类型需要自行判断   |
|     基础语法  |  操作  |   合并两个指定长度的数组   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/merge/main.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/merge/main.php) | 截断数组>合并数组   |
|     基础语法  |  操作  |   数组元素去重   |  [go](https://github.com/hanbinglove/golang/blob/main/basic/arrUnique/main.go)  [php](https://github.com/hanbinglove/golang/blob/main/basic/arrUnique/main.php) | 1.php array_unique搞定,go没有，<br>2.先排序，再依次对比删除<br>3.标记元素是否出现过，再进行重复元素删除   |