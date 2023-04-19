package model

type UserBorrow struct {
	User  User
	Books []Book
}

type CurrentBooks struct {
	Book Book
	Sum  float32
}
