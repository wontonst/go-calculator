package lexical

import (
"io"
"os"
"fmt"
)
type Lexical struct {
file *os.File
}

func (l *Lexical) read() {
buffer := make([]byte, 1024)
for {
n, err := l.file.Read(buffer)
if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
fmt.Printf("%s",buffer)
}
}