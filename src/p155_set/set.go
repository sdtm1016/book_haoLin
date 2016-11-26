package p155_set

import (
	"bytes"
	"fmt"
)

// 类型定义
type HashSet struct {
	m map[interface{}]bool
}

// 添加
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}

	return false
}

// 删除
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

// 清空
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

// 是否包含某个元素值
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

// 长度
func (set *HashSet) Len() int {
	return len(set.m)
}

// 判断两个 HashSet 是否相同
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}

	if set.Len() != other.Len() {
		return false
	}

	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}

	return true
}

// 返回所有元素
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0

	for key := range set.m {
		if actualLen< initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}

		actualLen++
	}

	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}

	return snapshot
}

// 打印
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("set{")
	first := true

	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}

		buf.WriteString(fmt.Sprintf("%v", key))
	}

	buf.WriteString("}")
	return buf.String()
}

// 高级功能: 超集判断
func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}

	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}

	if oneLen > 0 && otherLen == 0 {
		return true
	}

	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}

	return true
}