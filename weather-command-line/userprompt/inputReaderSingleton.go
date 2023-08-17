package userprompt

import (
	"bufio"
	"os"
	"sync"
)

var lock = &sync.Mutex{}
var readerInstance *bufio.Reader

func getReaderInstance() *bufio.Reader{
	if readerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		
		if readerInstance == nil {
			readerInstance = bufio.NewReader(os.Stdin) 
		}
	}
	return readerInstance
}