// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package utils

import (
	"math"
	"strconv"
)

//int ->>> string
func (*GalangTransform) IntToString(i int) string {
	return strconv.Itoa(i)
}

//int64 ->>> string
func (*GalangTransform) Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//string ->>> int64
func (*GalangTransform) StringToInt64(i string) (int64, error) {
	return strconv.ParseInt(i, 10, 64)
}

//string ->>> int
func (*GalangTransform) StringToInt(i string) (int, error) {
	return strconv.Atoi(i)
}

//float64 ->>> int64
func (*GalangTransform) Float64ToInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

//float64 ->>> string
func (*GalangTransform) Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

//float64 ->>> string
func (*GalangTransform) Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

//int64 ->>> float64
func (*GalangTransform) Int64ToFloat64(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

func (*GalangTransform) BoolToString(b bool) string {
	return strconv.FormatBool(b)
}
