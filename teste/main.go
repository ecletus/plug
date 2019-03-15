package main

func main() {
	{
		defer println("sss")
	}
	println("a")
}
