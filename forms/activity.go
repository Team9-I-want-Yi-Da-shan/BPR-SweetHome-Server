package forms

type ActivityForm struct{}

type AllActivityForm struct {
	Name            string `form:"name" json:"title"`
	PersonID        int    `form:"person_id" json:"user_id"`
	FamilyID        int    `form:"family_id" json:"family_id"`
	Description     string `form:"description" json:"description"`
	Start_at        int    `form:"start_at" json:"start_at"`
	Finish_at       int    `form:"finish_at" json:"finish_at"`
	IsFinish        int    `form:"is_finish" json:"is_finish"`
	IsRepeat        int    `form:"is_repeat" json:"is_repeat"`
	Reminder        int    `form:"reminder" json:"reminder"`
	IsAlarm         int    `form:"is_alarm" json:"is_alarm"`
	Repeat_interval int    `form:"repeat_interval" json:"repeat_interval"`
	Participants    []int  `form:"participant" json:"participant"`
}

type ActivityDataFamilyForm struct {
	FamilyID int `form:"family_id" json:"family_id"`
	DateUnix int `form:"date" json:"date"`
}

type ActivityDataPersonForm struct {
	PersonId int `form:"person_id" json:"user_id"`
	DateUnix int `form:"date" json:"date"`
}

// type PersonalActivityForm struct {
// 	Name        string `form:"name" json:"name"`
// 	PersonID    int    `form:"person_id" json:"person_id" binding:"required"`
// 	Description string `form:"description" json:"description"`
// 	Start_at    int    `form:"start_at" json:"start_at"`
// 	Finish_at   int    `form:"finish_at" json:"finish_at"`
// 	isFinish    int    `form:"isFinish" json:"isFinish"`
// 	Reminder    int    `form:"reminder" json:"reminder"`
// }

// type FamilyActivityForm struct {
// 	Name        string `form:"name" json:"name"`
// 	FamilyID    int    `form:"family_id" json:"family_id" binding:"required"`
// 	Description string `form:"description" json:"description"`
// 	Start_at    int    `form:"start_at" json:"start_at"`
// 	Finish_at   int    `form:"finish_at" json:"finish_at"`
// 	isFinish    int    `form:"isFinish" json:"isFinish"`
// 	Reminder    int    `form:"reminder" json:"reminder"`
// }
