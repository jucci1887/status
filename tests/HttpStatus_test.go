/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"fmt"
	"github.com/jucci1887/status"
	"github.com/jucci1887/test"
	"testing"
)

type statusCode struct {
	code int
	msg  string
	data interface{}
}

func TestHttpStatusStart(t *testing.T) {
	test.Test.Start("HttpStatus")
}

func TestHttpStatus(t *testing.T) {
	msg := "Test http status "
	assert := 200
	result := status.OK.Code
	if result == assert {
		test.Test.T(t).Logs(msg + "code: " + fmt.Sprint(assert)).Ok(result)
	} else {
		test.Test.T(t).Logs(msg + "code: " + fmt.Sprint(assert)).No(result)
	}

	assertMsg := "OK"
	resultMsg := status.OK.Msg
	if assertMsg == resultMsg {
		test.Test.T(t).Logs(msg + "Msg: " + assertMsg).Ok(resultMsg)
	} else {
		test.Test.T(t).Logs(msg + "Msg: " + assertMsg).No(resultMsg)
	}
}

func TestHttpStatusEnd(t *testing.T) {
	test.Test.End("HttpStatus")
}
