package user

import "time"

type User struct {
	FirstName   string
	LastName    string
	Email       string
	Gender      string
	BirthDayStr string
	Phone       string
	Address     string
	JobInfo     string
	Disability  string
	Allergies   string
	RegDayTime  time.Time
	RegDayStr   string
}

func NewUser(csvArr []string) *User {
	return &User{
		FirstName:   csvArr[0],
		LastName:    csvArr[1],
		Email:       csvArr[2],
		Gender:      csvArr[3],
		BirthDayStr: csvArr[4],
		Phone:       csvArr[5],
		Address:     csvArr[6],
		JobInfo:     csvArr[7],
		Disability:  csvArr[8],
		Allergies:   csvArr[9],
		RegDayTime:  time.Now(),
	}
}

func (u *User) GetRegDayStr() string {
	if u.RegDayStr == "" {
		u.RegDayStr = u.RegDayTime.Format("2006-01-02")
	}
	return u.RegDayStr
}
