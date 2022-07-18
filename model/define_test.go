/*******
* @Author:qingmeng
* @Description:
* @File:define_test
* @Date:2022/7/17
 */

package model

import (
	"strconv"
	"testing"
)

var MoveTests = []struct {
	isRed          bool
	x1, y1, x2, y2 int
	out            string
}{
	{true, 1, 9, 1, 1, "未选中红方棋子"},
	{true, 2, 0, 1, 1, "移动错误"},
	{true, 0, 0, 0, 1, "成功移动"},
	{true, 4, 0, 1, 1, "移动错误"},
	{true, 3, 0, 2, 0, "移动错误"},
	{true, 6, 0, 1, 1, "移动错误"},
	{true, 7, 0, 1, 1, "成功移动"},
	{true, 8, 0, 1, 1, "成功移动"},
	{true, 8, 1, 1, 1, "未选中棋子"},
	{true, 5, 1, 1, 1, "未选中棋子"},
}

func TestMove(t *testing.T) {
	for i, test := range MoveTests {
		t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
			get := Move(test.isRed, test.x1, test.y1, test.x2, test.y2)
			if get != test.out {
				t.Errorf("got %s, want %s", get, test.out)
			}
		})
	}
}
