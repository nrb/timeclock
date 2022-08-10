package main

func main() {
	r := CreateRoutes()

	s := NewServer("", "", r)

	s.Run()

}
