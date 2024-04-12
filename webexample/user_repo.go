package main

type userRepo struct {
	users []user
}

func newUserRepo() userRepo {
	users := []user{
		newUser("ajones", "ajones@aol.com"),
		newUser("sanderson", "sanderson@gmail.com"),
		newUser("mwilliams", "mwilliams@gmail.com"),
		newUser("jsmith", "jsmith@outlook.com"),
	}
	return userRepo{users}
}
