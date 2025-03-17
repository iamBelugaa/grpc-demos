package main

import pb "github.com/iamNilotpal/grpc/proto"

func main() {
	person := pb.Person{
		FirstName: "Nilotpal",
		LastName:  "Deka",
		Dob:       "2001-09-22",
		Email:     "iamnilotpaldeka@gmail.com",
	}

	println(person.GetFirstName())
	println(person.GetLastName())
	println(person.GetDob())
	println(person.GetEmail())
}
