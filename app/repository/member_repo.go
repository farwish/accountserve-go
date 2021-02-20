package repository

import "github.com/gopher-lego/ginger/config"

type MemberResultData struct{
	MemberId string `json:"member_id"`
}

func MemberMobilePasswordQuery(mobile string, password string) (MemberResultData, error) {
	var memberResultData MemberResultData

	result := config.MyDB.
		Table("members").
		Select("member_id").
		Where("mobile = ?", mobile).
		Where("password = ?", password).
		First(&memberResultData)

	if result.Error != nil {
		return memberResultData, result.Error
	}

	return memberResultData, nil
}
