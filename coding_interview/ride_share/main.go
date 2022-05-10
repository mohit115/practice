package main

func createCabs() {
	cm := getCabMangerIns()
	for i := 1; i < 10; i++ {
		cm.registerCab(&cab{capacity: 4, cabId: i})
	}
}

func createUser() {

}

func main() {
	createCabs()

}
