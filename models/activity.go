package models

import (
	"errors"
	"gin-Home-server/db"
	"gin-Home-server/forms"
)

type Activity struct {
	ID              int64  `db:"id, primarykey, autoincrement" json:"activity_id"`
	Name            string `db:"name" json:"title"`
	PersonID        int    `db:"person_id" json:"user_id"`
	FamilyID        int    `db:"family_id" json:"family_id"`
	Description     string `db:"description" json:"description"`
	Start_at        int    `db:"start_at" json:"start_at"`
	Finish_at       int    `db:"finish_at" json:"finish_at"`
	IsFinish        int    `db:"isFinish" json:"is_finish"`
	Reminder        int    `db:"reminder" json:"reminder"`
	IsAlarm         int    `db:"isAlarm" json:"is_alarm"`
	IsRepeat        int    `db:"isRepeat" json:"is_repeat"`
	Repeat_interval int    `db:"repeat_interval" json:"repeat_interval"`
	Participant     []int  `db:"participant" json:"participant"`
}

type ActivityModel struct{}

//Create ...
func (m ActivityModel) Create(activity forms.AllActivityForm) (activityID int64, err error) {
	//personal activity
	if activity.FamilyID == 0 && activity.PersonID != 0 {
		tx, err := db.GetDB().Begin()
		if err != nil {
			return 0, err
		}
		defer tx.Commit()

		err = tx.QueryRow("INSERT INTO public.activity(person_id, name, description, start_at, finish_at,isFinish,reminder,isRepeat,isFamily,isalarm) "+
			"VALUES($1, $2, $3,$4,$5,$6,$7,$8,0,$9) RETURNING id",
			activity.PersonID, activity.Name, activity.Description, activity.Start_at, activity.Finish_at, activity.IsFinish,
			activity.Reminder, activity.IsRepeat, activity.IsAlarm).Scan(&activityID)
		if activity.IsRepeat == 1 && activity.Repeat_interval != 0 {
			repeatError := tx.QueryRow("INSERT INTO public.repeat(activity_id, repeat_interval, repeat_start) VALUES($1, $2, $3)", activityID, activity.Repeat_interval, activity.Start_at)
			if repeatError != nil {
				tx.Rollback()
				return activityID, repeatError.Err()
			}
		}
		return activityID, err

		//family activity
	} else if activity.PersonID == 0 && activity.FamilyID != 0 {
		tx, err := db.GetDB().Begin()
		if err != nil {
			return 0, err
		}
		defer tx.Commit()
		err = tx.QueryRow("INSERT INTO public.activity(family_id, name, description, start_at, finish_at,isFinish,reminder,isRepeat,isFamily,isAlarm) VALUES($1, $2, $3,$4,$5,$6,$7,$8,1,$9) RETURNING id;", activity.FamilyID, activity.Name, activity.Description, activity.Start_at, activity.Finish_at, activity.IsFinish, activity.Reminder, activity.IsRepeat, activity.IsAlarm).Scan(&activityID)
		if err != nil {
			tx.Rollback()
			return activityID, err
		}
		if activity.IsRepeat == 1 && activity.Repeat_interval != 0 {
			repeatError := db.GetDB().QueryRow("INSERT INTO public.repeat(activity_id, repeat_interval, repeat_start) VALUES($1, $2, $3)", activityID, activity.Repeat_interval, activity.Start_at)
			if repeatError != nil {
				tx.Rollback()
				return activityID, repeatError.Err()
			}
		}
		err = m.DBAddParticipant(activityID, activity)
		if err != nil {
			tx.Rollback()
			return activityID, err
		}

		return activityID, err
	}

	return activityID, errors.New("please enter family_id or person_id")
}

//Update
func (m ActivityModel) Update(activity_id int64, form forms.AllActivityForm) (err error) {
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	operation, err := tx.Exec("UPDATE public.activity SET name=$2,description=$3,start_at=$4,finish_at=$5,reminder=$6,isfinish=$7,isrepeat=$8, isalarm=$9 WHERE id=$1", activity_id, form.Name, form.Description, form.Start_at, form.Finish_at, form.Reminder, form.IsFinish, form.IsRepeat, form.IsAlarm)
	if err != nil {
		return err
	}
	if form.IsRepeat == 1 && form.Repeat_interval != 0 {
		operation, err := tx.Exec("UPDATE public.repeat SET repeat_interval=$2, repeat_start=$3 WHERE activity_id=$1", activity_id, form.Repeat_interval, form.Start_at)
		if err != nil {
			tx.Rollback()
			return err
		}
		operation.RowsAffected()
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		tx.Rollback()
		return errors.New("updated 0 records")
	}

	return err
}

func (m ActivityModel) DBAddParticipant(activityID int64, activity forms.AllActivityForm) (err error) {

	if activity.FamilyID != 0 {
		if activity.Participants != nil || len(activity.Participants) != 0 {

			for i := 0; i < len(activity.Participants); i++ {
				err = db.GetDB().QueryRow("INSERT INTO public.activity_participant(user_id, activity_id) VALUES($1, $2)", activity.Participants[i], activityID).Err()
				if err != nil {
					db.GetDB().Query("ROLLBACK;")
					return err
				}
			}
			return err
		}
	}
	return nil
}

//FamilyIDDate
func (m ActivityModel) FamilyIDDate(form forms.ActivityDataFamilyForm) (activities []FamilyList, err error) {
	_, err = db.GetDB().Select(&activities, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data FROM(SELECT a.id AS activity_id,a.family_id,a.name as title,a.description,a.start_at,a.finish_at,a.isfinish as is_finish,a.reminder,a.isalarm as is_alarm,a.isrepeat as is_repeat,(array(select user_id FROM activity_participant p WHERE p.activity_id=a.id)) AS participant ,r.repeat_interval FROM public.activity a LEFT JOIN  repeat r on a.id = r.activity_id where (family_id=1 AND start_at BETWEEN $2 AND $2+86400) OR (family_id=$1 AND a.isrepeat =1 AND ($2-r.repeat_start)%r.repeat_interval<86400)) d", form.FamilyID, form.DateUnix)
	return activities, err
}

//PersonIDDate
func (m ActivityModel) PersonIDDate(form forms.ActivityDataPersonForm) (activities []FamilyList, err error) {
	_, err = db.GetDB().Select(&activities, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data FROM(SELECT a.id AS activity_id,a.person_id AS user_id,a.name as title,a.description,a.start_at,a.finish_at,a.isfinish as is_finish,a.reminder,a.isalarm as is_alarm,a.isrepeat as is_repeat, r.repeat_interval  FROM public.activity a LEFT JOIN  repeat r on a.id = r.activity_id  where (person_id=$1 AND start_at BETWEEN $2 AND $2+86400) OR (person_id=$1 AND a.isrepeat =1 AND ($2-r.repeat_start)%r.repeat_interval<86400)) d", form.PersonId, form.DateUnix)
	return activities, err
}

//GetActivityByID
func (m ActivityModel) GetActivityByID(id int64) (activity Activity, err error) {
	err = db.GetDB().SelectOne(&activity, "SELECT a.id,a.family_id,a.name as title,a.description,a.start_at,a.finish_at,a.isfinish as is_finish,a.reminder,a.isalarm as is_alarm,a.isrepeat as is_repeat,(array(select user_id FROM activity_participant p WHERE p.activity_id=a.id)) AS partivipant,r.repeat_interval FROM public.activity a LEFT JOIN  repeat r on a.id = r.activity_id where a.id=$1  LIMIT 1", id)
	return activity, err
}

//delete activity by id

func (m ActivityModel) DeleteActivityByID(id int64) (err error) {

	operation, err := db.GetDB().Exec("DELETE FROM public.activity WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}

//setFinish
func (m ActivityModel) SetFinish(id int64) (err error) {

	operation, err := db.GetDB().Exec("UPDATE public.activity SET isfinish = 1 WHERE id = $1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}
