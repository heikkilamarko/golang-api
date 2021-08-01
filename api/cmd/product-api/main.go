package main

import "product-api/internal/service"

func main() {
	(&service.Service{}).Run()
}
