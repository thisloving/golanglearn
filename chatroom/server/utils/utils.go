package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("wait for client...")

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		err = errors.New("read pkg header error")
		return
	}

	fmt.Printf("read length=%d\n", this.Buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmarshal err")
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write(len) err=", err)
		return
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write(data) err=", err)
		return
	}

	return
}
