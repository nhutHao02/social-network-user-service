package user

type UserQueryRepository interface {
	QueryMethod()
}

type UserCommandRepository interface {
	CommandMethod()
}
