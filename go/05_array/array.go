package _5_array

import (
	"fmt"

	"github.com/pkg/errors"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Author: seerjk
 */

type Array struct {
	data   []int // 存储数据的list
	length uint  // Array当前元素的长度，不同于capacity总容量
}

//为数组初始化内存
func NewArray(capacity uint) (*Array, error) {
	if capacity < 0 {
		panic("capacity should be larger than 0.")
	}
	return &Array{
		make([]int, capacity),
		0,
	}, nil
}

// 获取数组当前的length 元素个数
func (array *Array) Len() uint {
	return array.length
}

// 判断索引是否越界
func (array *Array) isIndexOutOfRange(index uint) bool {
	return index >= array.length
}

// 通过索引查找数组，索引范围[0, n-1]
func (array *Array) Find(index uint) (int, error) {
	if array.isIndexOutOfRange(index) {
		errMsg := fmt.Sprintf("%d is out of index.", index)
		return 0, errors.New(errMsg)
	}
	return array.data[index], nil
}

// TODO: 容量不足时，扩容capacity为1.5倍

// 插入数值到索引index上
func (array *Array) Insert(index uint, v int) error {
	// index 超过范围 index >= capacity-1
	if int(index) >= len(array.data) {
		errMsg := fmt.Sprintf("the index %d is out of array's capacity.", index)
		return errors.New(errMsg)
	}
	// array 满了，无法插入新值
	if int(array.length) == len(array.data)-1 {
		errMsg := fmt.Sprintf("the array don't have enough space.")
		return errors.New(errMsg)
	}

	// v插入到index
	for i := array.length

	array.length += 1
	return nil
}

// 插入到末尾
func (array *Array) InsertToTail(v int) error {
	return array.Insert(array.length, v)
}

// 删除索引index上的值
// func (array *Array) Delete(index uint) (int, error) {}

// 打印数列
func (array *Array) Print() {
	var i uint
	for i = 0; i < array.length; i++ {
		fmt.Println(array.data[i])
	}
}
