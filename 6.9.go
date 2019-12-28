func serverSide() {
	rpc.Register (new(ExamleStruct))
    ln, _ := nrt.Listen("tcp","8191")
}