package User

import (
	// "practice/Mocking/GoMock/IUser"

	"github.com/Param-Scalent/Go_Practice/Mocking/GoMock/IUser"
)

type User struct {
	IUser IUser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "Sample Test")
}
