package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

//PlanForm
type PlanForm struct{}

//add personal plan form
type PersonalPlanForm struct {
	Name        string `form:"name" json:"title"`
	PersonId    int    `form:"person_id" json:"user_id" binding:"required"`
	Description string `form:"description" json:"description"`
	Comment     string `form:"comment" json:"comment"`
}
type FamilyPlanForm struct {
	FamilyId    int    `form:"family_id" json:"family_id" binding:"required"`
	Name        string `form:"name" json:"title"`
	Description string `form:"description" json:"description"`
	Comment     string `form:"comment" json:"comment"`
}

//PersonId ...
func (f PlanForm) PersonId(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your PersonId"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}
func (f PlanForm) FamilyId(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your FamilyId"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

//addplan ...
func (f PlanForm) addPersonalPlan(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "PersonId" {
				return f.PersonId(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
