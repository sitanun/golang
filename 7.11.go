func main() {
	go serverside()
	go clientSide()

	var input string
	fmt.Scan(&input)
}