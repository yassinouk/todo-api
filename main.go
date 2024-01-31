package main

func main() {
	server := NewAPIServer(":8081")
	server.Run()

}
