func serverSide() {
	rpc.Register (new(ExamleStruct))
	ln, _ := nrt.Listen("tcp","8191")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			go rpc.ServeConn(conn)
		}
	}
}