/*******
* @Author:qingmeng
* @Description:
* @File:uuid_test
* @Date:2022/7/17
 */

package tools

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	get := GetUUID()
	if get == "" {
		t.Errorf("got:%s", get)
	}
}
