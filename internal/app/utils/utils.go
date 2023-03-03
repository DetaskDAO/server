package utils

import (
	"encoding/json"
	"github.com/lib/pq"
	"golang.org/x/exp/constraints"
	"strconv"
)

// SliceIsExist 判断元素是否在slice
func SliceIsExist[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func StructToMap(inter []any) map[string]interface{} {
	var m map[string]interface{}
	for _, v := range inter {
		if v == nil {
			continue
		}
		ja, _ := json.Marshal(v)
		json.Unmarshal(ja, &m)
	}
	return m
}

func MapPushStruct(m map[string]interface{}, inter []any) map[string]interface{} {
	for _, v := range inter {
		ja, _ := json.Marshal(v)
		json.Unmarshal(ja, &m)
	}
	return m
}

func SliceMax[T constraints.Ordered](slice []T) (index int, m T) {
	for i, e := range slice {
		if i == 0 || e > m {
			m = e
			index = i
		}
	}
	return
}

func SliceMin[T constraints.Ordered](slice []T) (index int, m T) {
	for i, e := range slice {
		if i == 0 || e < m {
			m = e
			index = i
		}
	}
	return
}

var CurrencyNames = map[uint8]string{1: "USD", 2: "ETH", 3: "BNB", 4: "MATIC", 5: "OP"}

func ParseSkills(skills int64) (res pq.Int64Array) {
	skillsBin := strconv.FormatInt(skills, 2)
	var skillsList pq.Int64Array
	for i := 0; i <= len(skillsBin)/12; i++ {
		if i == len(skillsBin)/12 {
			skillInt, _ := strconv.ParseInt(skillsBin[:len(skillsBin)-12*i], 2, 64)
			skillsList = append(skillsList, skillInt)
		} else {
			skillInt, _ := strconv.ParseInt(skillsBin[len(skillsBin)-12*(i+1):len(skillsBin)-12*i], 2, 64)
			skillsList = append(skillsList, skillInt)
		}
	}
	return skillsList
}
