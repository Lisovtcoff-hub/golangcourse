package main

import (
	"context"
	"fmt"

	. "gitlab.golang-school.ru/potok-2/lessons/lesson-16/pkg/profile_client"
)

var ctx = context.Background()

func main() {
	profile := New(Config{Host: "localhost", Port: "8080"})

	id, err := profile.Create(ctx, "John", 25, "john@gmail.com", "+73003002020")
	if err != nil {
		panic(err)
	}

	p, err := profile.GetProfile(ctx, id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
	fmt.Println(p.Age)
	fmt.Println(p.Name)
	fmt.Println(p.Contacts.Email)
	fmt.Println(p.Contacts.Phone)

	var (
		name  = "John Doe"
		age   = 26
		email = "new-john@gmail.com"
		phone = "+73003004000"
	)

	err = profile.Update(ctx, id.String(), &name, &age, &email, &phone)
	if err != nil {
		panic(err)
	}

	p, err = profile.GetProfile(ctx, id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
	fmt.Println(p.Age)
	fmt.Println(p.Name)
	fmt.Println(p.Contacts.Email)
	fmt.Println(p.Contacts.Phone)

	err = profile.Delete(ctx, id.String())
	if err != nil {
		panic(err)
	}

	_, err = profile.GetProfile(ctx, id.String())

	fmt.Println("Get request:", err)
}
