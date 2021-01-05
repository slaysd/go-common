// 채널 유틸리티 패키지
package channel

import (
	"reflect"
)

// MergeChannel 여러개의 채널을 단일 채널로 합침
func MergeChannel(out interface{}, channels ...interface{}) {
	cases := make([]reflect.SelectCase, len(channels))
	for idx, ch := range channels {
		cases[idx] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}

	output := reflect.ValueOf(out)
	go func() {
		for {
			_, value, ok := reflect.Select(cases)
			if !ok {
				break
			}
			output.Send(value)
		}
		output.Close()
	}()
}