package main

func main() {
	service := RestController{&Telega{}}
	service.Start()
}
