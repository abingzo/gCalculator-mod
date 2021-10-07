package stdio

import (
	"gCalculator-mod/base"
	"io/ioutil"
	"time"
)

// ReadStdIn 间隔0.1秒读一次数据
// 阻塞读取，读取到数据则解除阻塞
func ReadStdIn() error {
	for {
		time.Sleep(time.Millisecond * 100)
		bytes, err := ioutil.ReadAll(base.OsStdIn)
		if err != nil {
			return err
		}
		if len(bytes) > 0 {
			_, err = base.StdIn.Write(bytes)
			if err != nil {
				return err
			}
			return nil
		}
	}
}


func WriteStdOut() error {
	tmp := make([]byte,0)
	_,err := base.StdOut.Read(tmp)
	if err != nil {
		return err
	}
	_, err = base.OsStdOut.Write(tmp)
	if err != nil {
		return err
	}
	return nil
}
