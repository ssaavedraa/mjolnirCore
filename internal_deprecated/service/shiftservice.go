package shiftService

import (
	"time"

	"hex/cms/internal_deprecated/db"
	"hex/cms/internal_deprecated/model"
)

func GetShiftsInRage(lowerLimit, upperLimit, userId int) ([]model.Shift, error) {
	var shifts = []model.Shift{}

	database, err := db.GetInstance()

	if err != nil {
		return shifts, err
	}

	today := time.Now()
	lowerDate := today.AddDate(0, 0, -lowerLimit)
	upperDate := today.AddDate(0, 0, upperLimit)

	if (lowerLimit > 0 && upperLimit > 0) {
		database = database.Where("to_char(start, 'YYYY-MM-DD') >= ?", lowerDate.Format("2006-01-02")).Where("to_char(start, 'YYYY-MM-DD') <= ?", upperDate.Format("2006-01-02"))
	} else if lowerLimit > 0 && upperLimit == 0 {
		database = database.Where("to_char(start, 'YYYY-MM-DD') >= ?", lowerDate.Format("2006-01-02")).Where("to_char(start, 'YYYY-MM-DD') < ?", today.Format("2006-01-02"))
	} else if lowerLimit == 0 && upperLimit > 0 {
		database = database.Where("to_char(start, 'YYYY-MM-DD') > ?", today.Format("2006-01-02")).Where("to_char(start, 'YYYY-MM-DD') <= ?", upperDate.Format("2006-01-02"))
	}

	database.Where("user_id = ?", userId).Order("start ASC").Find(&shifts)

	/*
		{
			date: string,
			start: string,
			end: string,
			location: string,
			clockIn: string,
			clockOut: string
		}[]
	*/

	return shifts, nil
}