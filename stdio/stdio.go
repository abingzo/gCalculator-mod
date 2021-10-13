package stdio

import (
	"gCalculator-mod/base"
	"io/ioutil"
	"time"
)

// ReadStdIn 间隔0.1秒读一次数据
// 阻塞读取，读取到数据则解除阻塞
func ReadStdIn() ([]byte,error) {
	for {
		time.Sleep(time.Millisecond * 100)
		bytes, _ := ioutil.ReadAll(base.OsStdIn)
		if len(bytes) > 0 {
			return bytes,nil
		}
	}
}

func WriteStdOut(tmp []byte) error {
	_, err := base.OsStdOut.Write(tmp)
	if err != nil {
		return err
	}
	return nil
}
