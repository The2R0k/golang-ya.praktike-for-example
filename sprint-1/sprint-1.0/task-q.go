package main

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	length := strconv.Atoi(scanner.Text())