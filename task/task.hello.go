package task

import "fmt"

func SayHello(params string) error {
	fmt.Println("hello, this is task client demo.")
	return nil
}

func LocalTask(params string) error {
	fmt.Println("local task running...")
	return nil
}
