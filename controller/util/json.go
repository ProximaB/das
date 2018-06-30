// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func DasParseHtmlInputDateTime(input string) time.Time {
	layout := "2006-01-02T15:04"
	t, _ := time.Parse(layout, input)
	return t
}

func DasParseHtmlInputDate(input string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	return t
}

func decodeIntArray(value string) reflect.Value {
	value = strings.Replace(value, "[", "", -1)
	value = strings.Replace(value, "]", "", -1)
	split := strings.Split(value, ",")
	data := make([]int, 0)
	for _, each := range split {
		val, err := strconv.Atoi(each)
		if err != nil {
			return reflect.Value{}
		}
		data = append(data, val)
	}
	return reflect.ValueOf(data)
}

// register this to gorilla schema decoder to decode time
func decodeDate(value string) reflect.Value {
	s, err := time.Parse("2006-01-02", value)
	if err != nil {
		return reflect.ValueOf(DasParseHtmlInputDateTime(value))
	}
	return reflect.ValueOf(s)
}
