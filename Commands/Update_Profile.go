package commands

import (
	dtos "github.com/datrine/alumni_business/Dtos/Command"
	entities "github.com/datrine/alumni_business/Entities"
)

func UpdateUserProfile(data *dtos.UpdateUserProfileCommandData) (*entities.User, error) {
	/**/ user, err := entities.GetUserByID("")
	if err != nil {
		return nil, err
	}
	return user.UpdateUserProfile(&dtos.UpdateUserProfileCommandData{
		FirstName:         data.FirstName,
		LastName:          data.LastName,
		Education:         data.Education,
		Profession:        data.Profession,
		JobTitle:          data.JobTitle,
		Certifications:    data.Certifications,
		Employer:          data.Employer,
		GraduationYear:    data.GraduationYear,
		Industry:          data.Industry,
		ProfilePictureUrl: data.ProfilePictureUrl,
		Skills:            data.Skills,
		DOB:               data.DOB,
		Location:          data.Location,
	})

}
