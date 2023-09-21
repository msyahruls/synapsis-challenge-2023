package helper

import (
	"synapsis-challange/model/domain"
	"synapsis-challange/model/web"
)

// User
// func ToDetailUserResponse(user domain.User) web.UserDetailResponse {
// 	// occupation := ToOccupationUserResponse(user.Occupation)
// 	// recruitment := ToRecruitResponse(user.Recruitment)
// 	var linkImage string
// 	if user.ProfileImage != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), user.ProfileImage)
// 	}
// 	return web.UserDetailResponse{
// 		Id:           user.Id,
// 		OccupationId: user.OccupationId,
// 		NIK:          user.NIK,
// 		Name:         user.Name,
// 		Email:        user.Email,
// 		Phone:        user.Phone,
// 		Gender:       user.Gender,
// 		DistrictId:   user.DistrictId,
// 		Longitude:    user.Longitude,
// 		Latitude:     user.Latitude,
// 		ReferralCode: user.ReferralCode,
// 		ProfileImage: linkImage,
// 		Status:       user.Status,
// 		Occupation:   occupation,
// 		Recruitment:  recruitment,
// 		CreatedAt:    user.CreatedAt,
// 		UpdatedAt:    user.UpdatedAt,
// 	}
// }

func ToUserResponse(user domain.User) web.UserResponse {
	// occupation := ToOccupationUserResponse(user.Occupation)
	// var linkImage string
	// if user.ProfileImage != "" {
	// 	linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), user.ProfileImage)
	// }

	return web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		// Gender:     user.Gender,
		// Phone:      user.Phone,
		// NIK:        user.NIK,
		// DistrictId: user.DistrictId,
		// Longitude:  user.Longitude,
		// Latitude:   user.Latitude,
		// ProfileImage: linkImage,
		// Occupation:   occupation,
		// Status:    user.Status,
		// Accesses:  user.Accesses,
	}
}

// func ToUserAyrshareResponse(user domain.UserAyrshare) web.UserResponseAyrshare {
// 	return web.UserResponseAyrshare{
// 		Id:            user.Id,
// 		Name:          user.Name,
// 		Email:         user.Email,
// 		AyrshareToken: user.AyrshareToken,
// 		CreatedAt:     user.CreatedAt,
// 	}
// }

// func ToOccupationUserResponse(occupation domain.Occupation) web.OccupationResponseUser {
// 	return web.OccupationResponseUser{
// 		Level: occupation.Level,
// 		Name:  occupation.Name,
// 	}
// }

// func ToAllUserResponses(users []domain.User) []web.UserResponse {
// 	var userResponses []web.UserResponse
// 	for _, user := range users {
// 		userResponses = append(userResponses, ToUserResponse(user))
// 	}
// 	return userResponses
// }

// func ToAllUserAyrshareResponses(users []domain.UserAyrshare) []web.UserResponseAyrshare {
// 	var userResponses []web.UserResponseAyrshare
// 	for _, user := range users {
// 		userResponses = append(userResponses, ToUserAyrshareResponse(user))
// 	}
// 	return userResponses
// }

// func ToAllUserDetailResponses(users []domain.User) []web.UserDetailResponse {
// 	var userResponses []web.UserDetailResponse
// 	for _, user := range users {
// 		userResponses = append(userResponses, ToDetailUserResponse(user))
// 	}
// 	return userResponses
// }

// // Occupation
// func ToOccupationResponse(occupation domain.Occupation) web.OccupationResponse {
// 	return web.OccupationResponse{
// 		ID:    occupation.ID,
// 		Level: occupation.Level,
// 		Name:  occupation.Name,
// 	}
// }

// func ToOccupationResponses(occupations []domain.Occupation) []web.OccupationResponse {
// 	var occupationResponses []web.OccupationResponse
// 	for _, occupation := range occupations {
// 		occupationResponses = append(occupationResponses, ToOccupationResponse(occupation))
// 	}
// 	return occupationResponses
// }

// // Rrecuitment
// func ToRecruitResponse(recruit domain.UserRecruit) web.UserRecruitResponse {
// 	return web.UserRecruitResponse{
// 		RecruitProgress: recruit.RecruitProgress,
// 		RecruitGoal:     recruit.RecruitGoal,
// 	}
// }

// func ToAllRecruitResponse(recruit domain.UserRecruitDetail) web.UserRecruitDetail {
// 	user := ToUserForRecruitResponses(recruit.RecruitedUser)
// 	return web.UserRecruitDetail{
// 		ID:              recruit.ID,
// 		RecruitProgress: recruit.RecruitProgress,
// 		RecruitGoal:     recruit.RecruitGoal,
// 		Name:            recruit.User.Name,
// 		RecruitedUser:   user,
// 	}
// }

// func ToAllRecruitResponses(recruits []domain.UserRecruitDetail) []web.UserRecruitDetail {
// 	var recruitResponses []web.UserRecruitDetail
// 	for _, recruit := range recruits {
// 		recruitResponses = append(recruitResponses, ToAllRecruitResponse(recruit))
// 	}
// 	return recruitResponses
// }

// func ToRecruitDetailResponse(recruit domain.RecruitedDetail) web.RecruitDetailUserResponse {
// 	user := ToUserForRecruitResponses(recruit.RecruitedUser)
// 	return web.RecruitDetailUserResponse{
// 		Name:            recruit.User.Name,
// 		RecruitProgress: recruit.RecruitProgress,
// 		RecruitGoal:     recruit.RecruitGoal,
// 		RecrutedUser:    user,
// 	}

// }

// func ToUserForRecruitResponse(user domain.User) web.UserForRecruitResponse {
// 	return web.UserForRecruitResponse{
// 		ID:          user.Id,
// 		Name:        user.Name,
// 		Email:       user.Email,
// 		Nik:         user.NIK,
// 		Phone:       user.Phone,
// 		Recruitment: user.Recruitment,
// 	}
// }

// func ToUserForRecruitResponses(users []domain.User) []web.UserForRecruitResponse {
// 	var userResponses []web.UserForRecruitResponse
// 	for _, user := range users {
// 		userResponses = append(userResponses, ToUserForRecruitResponse(user))
// 	}
// 	return userResponses
// }

// // Progres Survey
// func ToProgressSurveyResponses(surveys []domain.ProgressSurvey) []web.ProgressSurvey {
// 	var progresResponses []web.ProgressSurvey
// 	for _, progres := range surveys {
// 		progresResponses = append(progresResponses, ToProgressSurveyResponse(progres))
// 	}
// 	return progresResponses
// }

// func ToProgressSurveyResponse(progress domain.ProgressSurvey) web.ProgressSurvey {
// 	return web.ProgressSurvey{
// 		SurveyProgress: progress.SurveyProgress,
// 		SurveyGoal:     progress.SurveyGoal,
// 	}
// }

// func ToAllProgressResponse(recruit domain.UserProgressDetail) web.UserProgressRespondentDetail {
// 	user := ToUserForProgressResponses(recruit.RecruitedUser)
// 	return web.UserProgressRespondentDetail{
// 		ID:             recruit.ID,
// 		SurveyProgress: recruit.SurveyProgress,
// 		SurveyGoal:     recruit.SurveyGoal,
// 		Name:           recruit.User.Name,
// 		RecruitedUser:  user,
// 	}
// }

// func ToAllProgressResponses(recruits []domain.UserProgressDetail) []web.UserProgressRespondentDetail {
// 	var recruitResponses []web.UserProgressRespondentDetail
// 	for _, recruit := range recruits {
// 		recruitResponses = append(recruitResponses, ToAllProgressResponse(recruit))
// 	}
// 	return recruitResponses
// }

// func ToProgressDetailResponse(recruit domain.ProgressSurveyDetail) web.UserProgressRespondentDetail {
// 	user := ToUserForProgressResponses(recruit.RespondentUser)
// 	return web.UserProgressRespondentDetail{
// 		Name:           recruit.User.Name,
// 		SurveyProgress: recruit.SurveyProgress,
// 		SurveyGoal:     recruit.SurveyGoal,
// 		RecruitedUser:  user,
// 	}

// }

// func ToUserForProgressResponse(user domain.Respondent) web.UserForProgressResponse {
// 	return web.UserForProgressResponse{
// 		ID:    user.Id,
// 		Name:  user.Name,
// 		Email: user.Email,
// 		Phone: user.Phone,
// 		//Respondent: user.,
// 	}
// }

// func ToUserForProgressResponses(users []domain.Respondent) []web.UserForProgressResponse {
// 	var userResponses []web.UserForProgressResponse
// 	for _, user := range users {
// 		userResponses = append(userResponses, ToUserForProgressResponse(user))
// 	}
// 	return userResponses
// }

// // Profile
// func ToDetailProfileResponse(user domain.Profile) web.UserDetailResponse {
// 	occupation := ToOccupationUserResponse(user.Occupation)
// 	recruitment := ToRecruitResponse(user.Recruitment)
// 	survey := ToProgressSurveyResponse(user.Survey)
// 	logistic := ToLogisticSumResponse(user.Logistic)
// 	var linkImage string
// 	if user.ProfileImage != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), user.ProfileImage)
// 	}

// 	return web.UserDetailResponse{
// 		Id:            user.Id,
// 		OccupationId:  user.OccupationId,
// 		NIK:           user.NIK,
// 		Name:          user.Name,
// 		Email:         user.Email,
// 		Phone:         user.Phone,
// 		Gender:        user.Gender,
// 		ReferralCode:  user.ReferralCode,
// 		GoogleAdsID:   user.GoogleAdsID,
// 		AyrshareToken: user.AyrshareToken,
// 		DetermID:      user.DetermID,
// 		ProfileImage:  linkImage,
// 		Occupation:    occupation,
// 		Recruitment:   recruitment,
// 		Logistic:      logistic,
// 		Survey:        survey,
// 		CreatedAt:     user.CreatedAt,
// 		UpdatedAt:     user.UpdatedAt,
// 		Accesses:      user.Accesses,
// 	}
// }

// func ToProfileResponse(user domain.Profile) web.ProfileResponse {
// 	return web.ProfileResponse{
// 		Id:        user.Id,
// 		Name:      user.Name,
// 		Email:     user.Email,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}
// }

// // Survey
// func ToSurveyResponse(survey domain.Survey) web.SurveyResponse {
// 	question := ToQuestionForSurveyResponses(survey.Questions)
// 	return web.SurveyResponse{
// 		ID:             survey.ID,
// 		SurveyName:     survey.SurveyName,
// 		Status:         survey.Status,
// 		TotalResponden: survey.TotalRespondent,
// 		CreatedAt:      survey.CreatedAt,
// 		UpdatedAt:      survey.UpdatedAt,
// 		Question:       question,
// 	}
// }

// func ToQuestionForSurveyResponses(questions []domain.Questions) []web.QuestionForSurvey {
// 	var questionResponses []web.QuestionForSurvey
// 	for _, question := range questions {
// 		questionResponses = append(questionResponses, ToQuestionForSurveyResponse(question))
// 	}
// 	return questionResponses
// }

// func ToQuestionForSurveyResponse(question domain.Questions) web.QuestionForSurvey {
// 	return web.QuestionForSurvey{
// 		QuestionId:   question.ID,
// 		InputType:    question.InputType,
// 		QuestionName: question.QuestionName,
// 	}
// }

// func ToSurveyDetailResponse(survey domain.Survey) web.SurveyDetailResponse {
// 	question := ToQuestionSurveyResponses(survey.Questions)
// 	return web.SurveyDetailResponse{
// 		ID:             survey.ID,
// 		SurveyName:     survey.SurveyName,
// 		Status:         survey.Status,
// 		TotalResponden: survey.TotalRespondent,
// 		CreatedAt:      survey.CreatedAt,
// 		UpdatedAt:      survey.UpdatedAt,
// 		Questions:      question,
// 	}
// }

// func ToAllSurveyResponses(surveys []domain.Survey) []web.SurveyResponse {
// 	var surveyResponses []web.SurveyResponse
// 	for _, survey := range surveys {
// 		surveyResponses = append(surveyResponses, ToSurveyResponse(survey))
// 	}
// 	return surveyResponses
// }

// func ToQuestionSurveyResponses(questions []domain.Questions) []web.SurveyQuestionResponse {
// 	var questionResponses []web.SurveyQuestionResponse
// 	for _, question := range questions {
// 		questionResponses = append(questionResponses, ToQuestionSurveyResponse(question))
// 	}
// 	return questionResponses
// }

// func ToQuestionSurveyResponse(question domain.Questions) web.SurveyQuestionResponse {
// 	options := ToOptionSurveyResponses(question.Options)
// 	return web.SurveyQuestionResponse{
// 		ID:              question.ID,
// 		SurveyId:        question.SurveyId,
// 		Section:         question.Section,
// 		InputType:       question.InputType,
// 		QuestionNumber:  question.QuestionNumber,
// 		QuestionName:    question.QuestionName,
// 		QuestionSubject: question.QuestionSubject,
// 		Options:         options,
// 	}
// }

// func ToOptionSurveyResponses(options []domain.Options) []web.SurveyOptionResponse {
// 	var optionResponses []web.SurveyOptionResponse
// 	for _, option := range options {
// 		optionResponses = append(optionResponses, ToOptionSurveyResponse(option))
// 	}
// 	return optionResponses
// }

// func ToOptionSurveyResponse(option domain.Options) web.SurveyOptionResponse {
// 	return web.SurveyOptionResponse{
// 		ID:         option.ID,
// 		QuestionId: option.QuestionId,
// 		OptionName: option.OptionName,
// 		Value:      *option.Value,
// 		Color:      option.Color,
// 	}
// }

// func ToSurveyResponseByUserId(survey domain.SurveyWithResponseId) web.SurveyResponseByUserId {
// 	return web.SurveyResponseByUserId{
// 		ID:               survey.ID,
// 		SurveyName:       survey.SurveyName,
// 		Status:           survey.Status,
// 		TotalResponden:   survey.TotalRespondent,
// 		CreatedAt:        survey.CreatedAt,
// 		UpdatedAt:        survey.UpdatedAt,
// 		ResponseIdByUser: survey.ResponseId,
// 	}
// }

// func ToAllSurveyResponsesByUserId(surveys []domain.SurveyWithResponseId) []web.SurveyResponseByUserId {
// 	var surveyResponses []web.SurveyResponseByUserId
// 	for _, survey := range surveys {
// 		surveyResponses = append(surveyResponses, ToSurveyResponseByUserId(survey))
// 	}
// 	return surveyResponses
// }

// // Options
// func ToOptionResponse(option domain.Options) web.OptionResponse {
// 	return web.OptionResponse{
// 		ID:         option.ID,
// 		QuestionId: option.QuestionId,
// 		OptionName: option.OptionName,
// 		Value:      *option.Value,
// 		Color:      option.Color,
// 	}
// }

// func ToOptionResponses(options []domain.Options) []web.OptionResponse {
// 	var optionResponses []web.OptionResponse
// 	for _, option := range options {
// 		optionResponses = append(optionResponses, ToOptionResponse(option))
// 	}
// 	return optionResponses
// }

// // Question
// func ToQuestionResponses(questions []domain.Questions) []web.QuestionResponse {
// 	var questionResponses []web.QuestionResponse
// 	for _, question := range questions {
// 		questionResponses = append(questionResponses, ToQuestionResponse(question))
// 	}
// 	return questionResponses
// }

// func ToQuestionResponse(question domain.Questions) web.QuestionResponse {
// 	options := ToOptionSurveyResponses(question.Options)
// 	return web.QuestionResponse{
// 		ID:              question.ID,
// 		SurveyId:        question.SurveyId,
// 		Section:         question.Section,
// 		InputType:       question.InputType,
// 		QuestionNumber:  question.QuestionNumber,
// 		QuestionName:    question.QuestionName,
// 		QuestionSubject: question.QuestionSubject,
// 		Options:         options,
// 	}
// }

// // Answer
// func ToAnswerResponses(answers []domain.Answer) []web.AnswerResponse {
// 	var answerResponses []web.AnswerResponse
// 	for _, answer := range answers {
// 		answerResponses = append(answerResponses, ToAnswerResponse(answer))
// 	}
// 	return answerResponses
// }

// func ToAnswerResponse(answer domain.Answer) web.AnswerResponse {
// 	question := ToAnswerQuestionResponse(answer.Question)
// 	var option domain.Options
// 	if answer.Option.ID != "" {
// 		option = answer.Option
// 	} else if answer.AnswerText != "" {
// 		option.OptionName = answer.AnswerText
// 	} else {
// 		option.OptionName = fmt.Sprintf("%d", answer.AnswerNumeric)
// 	}

// 	return web.AnswerResponse{
// 		ID:            answer.ID,
// 		ResponseId:    answer.ResponseId,
// 		QuestionId:    answer.QuestionId,
// 		OptionsId:     answer.OptionsId,
// 		AnswerText:    answer.AnswerText,
// 		AnswerNumeric: answer.AnswerNumeric,
// 		Question:      question,
// 		Options:       option,
// 	}
// }

// func ToAnswerQuestionResponse(question domain.Questions) web.AnswerQuestionResponse {
// 	return web.AnswerQuestionResponse{
// 		ID:              question.ID,
// 		SurveyId:        question.SurveyId,
// 		Section:         question.Section,
// 		InputType:       question.InputType,
// 		QuestionNumber:  question.QuestionNumber,
// 		QuestionName:    question.QuestionName,
// 		QuestionSubject: question.QuestionSubject,
// 	}
// }

// func ToAnswerOptionResponse(option domain.Options) web.AnswerOptionResponse {
// 	return web.AnswerOptionResponse{
// 		ID:         option.ID,
// 		QuestionId: option.QuestionId,
// 		OptionName: option.OptionName,
// 		Value:      option.Value,
// 		Color:      option.Color,
// 	}
// }

// // Response
// func ToResponseResponses(response []domain.Response) []web.ResponseResponse {
// 	var responseResponses []web.ResponseResponse
// 	for _, response := range response {
// 		responseResponses = append(responseResponses, ToResponseResponse(response))
// 	}
// 	return responseResponses
// }

// func ToDetailResponseResponse(response domain.Response) web.DetailResponseResponse {
// 	answer := ToResponseAnswerResponses(response.Answers)
// 	var loc = web.Location{}
// 	if response.Location != nil {
// 		primif := response.Location.(primitive.D)
// 		here := primif.Map()
// 		if here["village_id"] != nil {
// 			loc.Villageid = here["village_id"].(string)
// 		}
// 		if here["address"] != nil {
// 			loc.Address = here["address"].(string)
// 		}
// 		loc.Longitude = here["longitude"].(string)
// 		loc.Latitude = here["latitude"].(string)
// 	}
// 	return web.DetailResponseResponse{
// 		ID:           response.ID,
// 		SurveyId:     response.SurveyId,
// 		SurveyName:   response.Survey.SurveyName,
// 		Respondent:   response.User.Name,
// 		RespondentId: response.RespondentId,
// 		RecruiterId:  response.Surveyor.Id,
// 		Recruiter:    response.Surveyor.Name,
// 		CreatedAt:    response.CreatedAt,
// 		Answers:      answer,
// 		Location:     loc,
// 	}
// }

// func ToCreateResponseResponse(response domain.Response) web.DetailResponseResponse {
// 	answer := ToResponseAnswerResponses(response.Answers)
// 	var loc = web.Location{}
// 	if response.Location != nil {
// 		here := response.Location.(map[string]interface{})
// 		if here["village_id"] != nil {
// 			loc.Villageid = here["village_id"].(string)
// 		}
// 		if here["address"] != nil {
// 			loc.Address = here["address"].(string)
// 		}
// 		loc.Longitude = here["longitude"].(string)
// 		loc.Latitude = here["latitude"].(string)
// 	}
// 	return web.DetailResponseResponse{
// 		ID:           response.ID,
// 		SurveyId:     response.SurveyId,
// 		RespondentId: response.RespondentId,
// 		Answers:      answer,
// 		CreatedAt:    response.CreatedAt,
// 		Location:     loc,
// 	}
// }

// func ToResponseResponse(response domain.Response) web.ResponseResponse {
// 	var loc = web.Location{}
// 	if response.Location != nil {
// 		primid := response.Location.(primitive.D)
// 		here := primid.Map()
// 		if here["village_id"] != nil {
// 			loc.Villageid = here["village_id"].(string)
// 		}
// 		if here["address"] != nil {
// 			loc.Address = here["address"].(string)
// 		}
// 		loc.Longitude = here["longitude"].(string)
// 		loc.Latitude = here["latitude"].(string)
// 	}

// 	return web.ResponseResponse{
// 		ID:           response.ID,
// 		SurveyId:     response.SurveyId,
// 		SurveyName:   response.Survey.SurveyName,
// 		Respondent:   response.User.Name,
// 		RespondentId: response.RespondentId,
// 		RecruiterId:  response.Surveyor.Id,
// 		Recruiter:    response.Surveyor.Name,
// 		CreatedAt:    response.CreatedAt,
// 		Location:     loc,
// 	}
// }

// func ToResponseAnswerResponses(answers []domain.Answer) []web.ResponseAnswerResponse {
// 	var answerResponses []web.ResponseAnswerResponse
// 	for _, answer := range answers {
// 		answerResponses = append(answerResponses, ToResponseAnswerResponse(answer))
// 	}
// 	return answerResponses
// }

// func ToResponseAnswerResponse(answer domain.Answer) web.ResponseAnswerResponse {
// 	question := ToResponseQuestionResponse(answer.Question)
// 	var ans string
// 	var arrAns []string
// 	if answer.OptionsId != "" {
// 		if answer.Option.OptionName == "" {
// 			arrAns = append(arrAns, answer.Regency.Name)
// 			arrAns = append(arrAns, answer.District.Name)
// 			arrAns = append(arrAns, answer.Village.Name)
// 		} else {
// 			ans = answer.Option.OptionName
// 		}
// 	} else if answer.AnswerText != "" {
// 		ans = answer.AnswerText
// 	} else {
// 		ans = fmt.Sprintf("%d", answer.AnswerNumeric)
// 	}

// 	var finalAns interface{}
// 	if ans == "" {
// 		finalAns = arrAns
// 	} else {
// 		finalAns = ans
// 	}

// 	return web.ResponseAnswerResponse{
// 		ID:            answer.ID,
// 		ResponseId:    answer.ResponseId,
// 		QuestionId:    answer.QuestionId,
// 		OptionsId:     answer.OptionsId,
// 		AnswerText:    answer.AnswerText,
// 		AnswerNumeric: answer.AnswerNumeric,
// 		Question:      question,
// 		Answer:        finalAns,
// 		//Options:       option,
// 	}
// }

// func ToResponseQuestionResponse(question domain.Questions) web.ResponseQuestionResponse {
// 	return web.ResponseQuestionResponse{
// 		ID:              question.ID,
// 		SurveyId:        question.SurveyId,
// 		Section:         question.Section,
// 		InputType:       question.InputType,
// 		QuestionNumber:  question.QuestionNumber,
// 		QuestionName:    question.QuestionName,
// 		QuestionSubject: question.QuestionSubject,
// 	}
// }

// func ToResponseOptionResponse(option domain.Options) web.ResponseOptionResponse {
// 	return web.ResponseOptionResponse{
// 		ID:         option.ID,
// 		QuestionId: option.QuestionId,
// 		OptionName: option.OptionName,
// 		Value:      *option.Value,
// 		Color:      option.Color,
// 	}
// }

// // Event
// func ToEventResponse(event domain.Event) web.EventResponse {
// 	var linkImage string
// 	if event.ImageUrl != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), event.ImageUrl)
// 	}
// 	return web.EventResponse{
// 		ID:            event.ID,
// 		EventName:     event.EventName,
// 		ImageUrl:      linkImage,
// 		Link:          event.Link,
// 		DateStart:     event.DateStart,
// 		DateEnd:       event.DateEnd,
// 		Location:      event.Location,
// 		ContactPerson: event.ContactPerson,
// 		Status:        event.Status,
// 		Description:   event.Description,
// 		CreatedAt:     event.CreatedAt,
// 		UpdatedAt:     event.UpdatedAt,
// 	}
// }

// func ToEventResponses(events []domain.Event) []web.EventResponse {
// 	var eventResponses []web.EventResponse
// 	for _, event := range events {
// 		eventResponses = append(eventResponses, ToEventResponse(event))
// 	}
// 	return eventResponses
// }

// func ToEventDetailUserResponse(event domain.Event) web.EventResponseDetailUser {
// 	var linkImage string
// 	if event.ImageUrl != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), event.ImageUrl)
// 	}
// 	return web.EventResponseDetailUser{
// 		ID:            event.ID,
// 		EventName:     event.EventName,
// 		ImageUrl:      linkImage,
// 		Link:          event.Link,
// 		DateStart:     event.DateStart,
// 		DateEnd:       event.DateEnd,
// 		Location:      event.Location,
// 		ContactPerson: event.ContactPerson,
// 		Status:        event.Status,
// 		Description:   event.Description,
// 		CreatedAt:     event.CreatedAt,
// 		UpdatedAt:     event.UpdatedAt,
// 	}
// }

// func ToEventDetailAdminResponse(event domain.Event) web.EventResponseDetailAdmin {
// 	var linkImage string
// 	if event.ImageUrl != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), event.ImageUrl)
// 	}
// 	user := ToEventParticipants(event.User)
// 	return web.EventResponseDetailAdmin{
// 		ID:            event.ID,
// 		EventName:     event.EventName,
// 		ImageUrl:      linkImage,
// 		Link:          event.Link,
// 		DateStart:     event.DateStart,
// 		DateEnd:       event.DateEnd,
// 		Location:      event.Location,
// 		ContactPerson: event.ContactPerson,
// 		Status:        event.Status,
// 		Description:   event.Description,
// 		CreatedAt:     event.CreatedAt,
// 		UpdatedAt:     event.UpdatedAt,
// 		Participant:   user,
// 	}
// }

// func ToEventParticipants(paticipants []domain.User) []web.EventParticipate {
// 	var participantResponses []web.EventParticipate
// 	for _, participant := range paticipants {
// 		participantResponses = append(participantResponses, ToEventParticipant(participant))
// 	}
// 	return participantResponses
// }

// func ToEventParticipant(participant domain.User) web.EventParticipate {
// 	return web.EventParticipate{
// 		ID:   participant.Id,
// 		Name: participant.Name,
// 	}
// }

// // // Social
// // func ToSocialSummaryResponse(summary domain.Summary) web.SocialSummaryResponse {
// // 	return web.SocialSummaryResponse{
// // 		Impressions:    summary.Impressions,
// // 		Engagements:    summary.Engagements,
// // 		EngagementRate: summary.EngagementRate,
// // 		PostLinkClicks: summary.PostLinkClicks,
// // 		UpdatedAt:      summary.UpdatedAt,
// // 	}
// // }

// // func ToSocialDatapointResponse(data domain.DataPoint) web.SocialDatapointResponse {
// // 	return web.SocialDatapointResponse{
// // 		Timestamp:      data.Timestamp,
// // 		AudienceGrowth: data.AudienceGrowth,
// // 		Impression:     data.Impression,
// // 		Engagement:     data.Engagement,
// // 		EngagementRate: data.EngagementRate,
// // 		VideoViews:     data.VideoViews,
// // 	}
// // }

// // func ToSocialResponse(summary domain.Summary, data []domain.DataPoint) web.SocialResponse {
// // 	var datas web.Data

// // 	for i := 0; i < len(data); i++ {
// // 		datas.DataId = append(datas.DataId, data[i].Id)

// // 		var temp web.SocialData
// // 		var tempFloat web.SocialDataFloat
// // 		temp.Timestamp = data[i].Timestamp
// // 		tempFloat.Timestamp = data[i].Timestamp

// // 		temp.Data = data[i].AudienceGrowth
// // 		datas.AudienceGrowth = append(datas.AudienceGrowth, temp)

// // 		temp.Data = data[i].Engagement
// // 		datas.Engagement = append(datas.Engagement, temp)

// // 		tempFloat.Data = data[i].EngagementRate
// // 		datas.EngagementRate = append(datas.EngagementRate, tempFloat)

// // 		temp.Data = data[i].Impression
// // 		datas.Impression = append(datas.Impression, temp)

// // 		temp.Data = data[i].AudienceGrowth
// // 		datas.VideoViews = append(datas.VideoViews, temp)
// // 	}

// // 	return web.SocialResponse{
// // 		Summary: web.Summary{
// // 			Impressions:    summary.Impressions,
// // 			Engagements:    summary.Engagements,
// // 			EngagementRate: summary.EngagementRate,
// // 			PostLinkClicks: summary.PostLinkClicks,
// // 		},
// // 		Data: datas,
// // 	}
// // }

// // Province
// func ToProvinceResponse(province domain.Province) web.ProvinceResponse {
// 	return web.ProvinceResponse{
// 		ID:   province.ID,
// 		Name: province.Name,
// 	}
// }

// func ToProvinceResponses(provinces []domain.Province) []web.ProvinceResponse {
// 	var provincesResponse []web.ProvinceResponse
// 	for _, province := range provinces {
// 		provincesResponse = append(provincesResponse, ToProvinceResponse(province))
// 	}
// 	return provincesResponse
// }

// // Regency
// func ToRegencyResponse(regency domain.Regency) web.RegencyResponse {
// 	return web.RegencyResponse{
// 		ID:         regency.ID,
// 		ProvinceId: regency.ProvinceId,
// 		Name:       regency.Name,
// 		AltName:    regency.AltName,
// 		Longitude:  regency.Longitude,
// 		Latitude:   regency.Latitude,
// 	}
// }

// func ToRegencyResponses(regencies []domain.Regency) []web.RegencyResponse {
// 	var regenciesResponse []web.RegencyResponse
// 	for _, regency := range regencies {
// 		regenciesResponse = append(regenciesResponse, ToRegencyResponse(regency))
// 	}
// 	return regenciesResponse
// }

// // District
// func ToDistrictResponse(district domain.Districts) web.DistrictsResponse {
// 	return web.DistrictsResponse{
// 		ID:        district.ID,
// 		RegencyId: district.RegencyId,
// 		Name:      district.Name,
// 		AltName:   district.AltName,
// 		Longitude: district.Longitude,
// 		Latitude:  district.Latitude,
// 	}
// }

// func ToDistrictResponses(districts []domain.Districts) []web.DistrictsResponse {
// 	var districtsResponse []web.DistrictsResponse
// 	for _, district := range districts {
// 		districtsResponse = append(districtsResponse, ToDistrictResponse(district))
// 	}
// 	return districtsResponse
// }

// // Village
// func ToVillageResponse(village domain.Village) web.VillageResponse {
// 	return web.VillageResponse{
// 		ID:         village.ID,
// 		DistrictId: village.DistrictId,
// 		Name:       village.Name,
// 		Longitude:  village.Longitude,
// 		Latitude:   village.Latitude,
// 	}
// }

// func ToVillageResponses(villages []domain.Village) []web.VillageResponse {
// 	var villageResponse []web.VillageResponse
// 	for _, village := range villages {
// 		villageResponse = append(villageResponse, ToVillageResponse(village))
// 	}
// 	return villageResponse
// }

// // User DPT
// func ToUserDptResponse(userDpt domain.UserDpt) web.UserDptResponse {
// 	return web.UserDptResponse{
// 		ID:                 userDpt.ID,
// 		NKK:                userDpt.NKK,
// 		NIK:                userDpt.NIK,
// 		Name:               userDpt.Name,
// 		TempatLahir:        userDpt.TempatLahir,
// 		TglLahir:           userDpt.TglLahir,
// 		StatusPerkawinan:   userDpt.StatusPerkawinan,
// 		JenisKelamin:       userDpt.JenisKelamin,
// 		JlnDukuh:           userDpt.JlnDukuh,
// 		Rt:                 userDpt.Rt,
// 		Rw:                 userDpt.Rw,
// 		Disabilitas:        userDpt.Disabilitas,
// 		StatusPerekamanKtp: userDpt.StatusPerekamanKtp,
// 		Keterangan:         userDpt.Keterangan,
// 		PollingStationId:   userDpt.PollingStationId,
// 		VillageId:          userDpt.VillageId,
// 		CreatedAt:          userDpt.CreatedAt,
// 		UpdatedAt:          userDpt.UpdatedAt,
// 		PollingStation:     userDpt.PollingStation,
// 		Village:            userDpt.Village,
// 	}
// }

// func ToUserDptResponses(userDpts []domain.UserDpt) []web.UserDptResponse {
// 	var userDptResponse []web.UserDptResponse
// 	for _, userDpt := range userDpts {
// 		userDptResponse = append(userDptResponse, ToUserDptResponse(userDpt))
// 	}
// 	return userDptResponse
// }

// // Survey Result
// func ToSurveyResultResponse(survey domain.SurveyResult) web.SurveyResult {
// 	question := ToQuestionSurveyResultResponses(survey.Questions)
// 	return web.SurveyResult{
// 		ID:             survey.ID,
// 		SurveyName:     survey.SurveyName,
// 		Status:         survey.Status,
// 		TotalResponden: survey.TotalRespondent,
// 		CreatedAt:      survey.CreatedAt,
// 		UpdatedAt:      survey.UpdatedAt,
// 		Questions:      question,
// 	}
// }

// func ToQuestionSurveyResultResponses(questions []domain.SurveyQuestionResult) []web.SurveyQuestionResult {
// 	var questionResponses []web.SurveyQuestionResult
// 	for _, question := range questions {
// 		questionResponses = append(questionResponses, ToQuestionSurveyResultResponse(question))
// 	}
// 	return questionResponses
// }

// func ToQuestionSurveyResultResponse(question domain.SurveyQuestionResult) web.SurveyQuestionResult {
// 	options := ToOptionSurveyResultResponses(question.Options)
// 	var answers []string
// 	for _, answer := range question.AnswerText {
// 		answers = append(answers, answer.AnswerText)
// 	}
// 	return web.SurveyQuestionResult{
// 		ID:              question.ID,
// 		SurveyId:        question.SurveyId,
// 		Section:         question.Section,
// 		InputType:       question.InputType,
// 		QuestionNumber:  question.QuestionNumber,
// 		QuestionName:    question.QuestionName,
// 		QuestionSubject: question.QuestionSubject,
// 		AnswerText:      answers,
// 		Options:         options,
// 		Location:        question.Location,
// 	}
// }

// func ToOptionSurveyResultResponses(options []domain.SurveyOptionAnswer) []web.SurveyOptionAnswer {
// 	var optionResponses []web.SurveyOptionAnswer
// 	for _, option := range options {
// 		optionResponses = append(optionResponses, ToOptionSurveyResultResponse(option))
// 	}
// 	return optionResponses
// }

// func ToOptionSurveyResultResponse(option domain.SurveyOptionAnswer) web.SurveyOptionAnswer {
// 	return web.SurveyOptionAnswer{
// 		ID:          option.ID,
// 		QuestionId:  option.QuestionId,
// 		OptionName:  option.OptionName,
// 		Value:       option.Value,
// 		TotalAnswer: option.TotalAnswer,
// 	}
// }

// func ToResultCountResponse(result domain.ResponseCountDate) web.ResponseCountDate {
// 	//date := fmt.Sprintf("%d-%d-%d", result.Date.Year, result.Date.Month, result.Date.Day)
// 	date := time.Date(int(result.Date.Year.(int32)), time.Month(int(result.Date.Month.(int32))), int(result.Date.Day.(int32)), 0, 0, 0, 0, time.UTC)
// 	return web.ResponseCountDate{
// 		Date:             date,
// 		RespondentId:     result.RespondentId,
// 		LastResponseDate: result.LastResponseDate,
// 		Count:            result.Count,
// 	}
// }

// func ToResultCountResponses(results []domain.ResponseCountDate) []web.ResponseCountDate {
// 	var responseResults []web.ResponseCountDate
// 	for _, result := range results {
// 		responseResults = append(responseResults, ToResultCountResponse(result))
// 	}
// 	return responseResults
// }

// // Complaint
// func ToComplaintResponse(complaint domain.Complaint) web.ComplaintResponse {
// 	var linkImage string
// 	if complaint.LinkImage != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), complaint.LinkImage)
// 	}
// 	return web.ComplaintResponse{
// 		ID:                  complaint.ID,
// 		SenderId:            complaint.SenderId,
// 		Title:               complaint.Title,
// 		Content:             complaint.Content,
// 		SenderName:          complaint.User.Name,
// 		LinkImage:           linkImage,
// 		Address:             complaint.Address,
// 		Longitude:           complaint.Longitude,
// 		Latitude:            complaint.Latitude,
// 		ComplaintStatus:     web.ComplaintStatus(complaint.ComplaintStatus),
// 		ComplaintCategory:   web.ComplaintCategory(complaint.ComplaintCategory),
// 		ComplaintStatusDesc: complaint.ComplaintStatusDesc,
// 		ReadStatus:          complaint.ReadStatus,
// 		CreatedAt:           complaint.CreatedAt,
// 		UpdatedAt:           complaint.UpdatedAt,
// 	}
// }

// func ToComplaintResultResponses(complaints []domain.Complaint) []web.ComplaintResponse {
// 	var complaintResponses []web.ComplaintResponse
// 	for _, complaint := range complaints {
// 		complaintResponses = append(complaintResponses, ToComplaintResponse(complaint))
// 	}
// 	return complaintResponses
// }

// func ToComplaintWithLastChatResponse(complaint domain.ComplaintResponseWithLastChat) web.ComplaintResponseWithLastChat {
// 	var linkImage string
// 	if complaint.LinkImage != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), complaint.LinkImage)
// 	}
// 	return web.ComplaintResponseWithLastChat{
// 		ID:                  complaint.ID,
// 		SenderId:            complaint.SenderId,
// 		Title:               complaint.Title,
// 		Content:             complaint.Content,
// 		SenderName:          complaint.User.Name,
// 		LinkImage:           linkImage,
// 		Address:             complaint.Address,
// 		Longitude:           complaint.Longitude,
// 		Latitude:            complaint.Latitude,
// 		ReadStatus:          complaint.ReadStatus,
// 		ComplaintStatus:     web.ComplaintStatus(complaint.ComplaintStatus),
// 		ComplaintStatusDesc: complaint.ComplaintStatusDesc,
// 		LastChat:            ToComplaintChatResponse(complaint.LastChat),
// 		ComplaintCategory:   web.ComplaintCategory(complaint.ComplaintCategory),
// 		CreatedAt:           complaint.CreatedAt,
// 		UpdatedAt:           complaint.UpdatedAt,
// 	}
// }

// func ToComplaintWithLastChatResultResponses(complaints []domain.ComplaintResponseWithLastChat) []web.ComplaintResponseWithLastChat {
// 	var complaintResponses []web.ComplaintResponseWithLastChat
// 	for _, complaint := range complaints {
// 		complaintResponses = append(complaintResponses, ToComplaintWithLastChatResponse(complaint))
// 	}
// 	return complaintResponses
// }

// // Complaint Chat
// func ToComplaintChatResponse(complaint domain.ComplaintChat) web.ComplaintChatResponse {
// 	var linkImage string
// 	if complaint.LinkImage != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), complaint.LinkImage)
// 	}
// 	return web.ComplaintChatResponse{
// 		ID:          complaint.ID,
// 		ComplaintId: complaint.ComplaintId,
// 		SenderId:    complaint.SenderId,
// 		LinkImage:   linkImage,
// 		Message:     complaint.Message,
// 		ReadStatus:  complaint.ReadStatus,
// 		Longitude:   complaint.Longitude,
// 		Latitude:    complaint.Latitude,
// 		CreatedAt:   complaint.CreatedAt,
// 		UpdatedAt:   complaint.UpdatedAt,
// 	}
// }

// func ToComplaintChatResultResponses(complaintChats []domain.ComplaintChat) []web.ComplaintChatResponse {
// 	var complaintChatResponses []web.ComplaintChatResponse
// 	for _, complaintChat := range complaintChats {
// 		complaintChatResponses = append(complaintChatResponses, ToComplaintChatResponse(complaintChat))
// 	}
// 	return complaintChatResponses
// }

// // Respondent
// func ToRespondentResponse(responden domain.Respondent) web.RespondentResponse {
// 	return web.RespondentResponse{
// 		Id:        responden.Id,
// 		Name:      responden.Name,
// 		Phone:     responden.Phone,
// 		Gender:    responden.Gender,
// 		Email:     responden.Email,
// 		CreatedAt: responden.CreatedAt,
// 	}
// }

// func ToRespondentResponses(respondent []domain.Respondent) []web.RespondentResponse {
// 	var respondents []web.RespondentResponse
// 	for _, respondent := range respondent {
// 		respondents = append(respondents, ToRespondentResponse(respondent))
// 	}
// 	return respondents
// }

// // Response by surveyor
// func ToResponseBySurveyor(response domain.ResponseBySurveyor) web.ResponseBySurveyor {
// 	var res web.ResponseBySurveyor
// 	res.Id = response.Id
// 	res.SurveyName = response.SurveyName
// 	res.RecruiterName = response.SurveyorName
// 	res.Response = response.Response
// 	return res
// }

// func ToResponsesBySurveyor(responses []domain.ResponseBySurveyor) []web.ResponseBySurveyor {
// 	var res []web.ResponseBySurveyor
// 	for _, response := range responses {
// 		res = append(res, ToResponseBySurveyor(response))
// 	}
// 	return res
// }

// func AnalyticWebToHistoricalDomain(socialmediaId string, lastTotal domain.SocialmediaLastTotal, recent web.AyrshareAnalyticsResponse, socmedType string, isNew bool) domain.SocialmediaHistorical {
// 	// total := domain.Socialmedia{
// 	// 	Facebook: domain.SocialmediaFacebook{
// 	// 		Likes:           recent.Facebook.Analytics.FanCount,
// 	// 		Followers:       recent.Facebook.Analytics.FollowersCount,
// 	// 		Engagement:      recent.Facebook.Analytics.PageEngagedUsers,
// 	// 		Impressions:     recent.Facebook.Analytics.PageImpressions,
// 	// 		ImpressionsPaid: recent.Facebook.Analytics.PageImpressionsPaid,
// 	// 		VideoViewTime:   recent.Facebook.Analytics.PageVideoViewTime,
// 	// 		VideoViews:      recent.Facebook.Analytics.PageVideoViews,
// 	// 		VideoViewsPaid:  recent.Facebook.Analytics.PageVideoViewsPaid,
// 	// 		Reactions:       recent.Facebook.Analytics.Reactions.Total,
// 	// 	},
// 	// 	Twitter: domain.SocialmediaTwitter{
// 	// 		Followers:      recent.Twitter.Analytics.FollowersCount,
// 	// 		Following:      recent.Twitter.Analytics.FriendsCount,
// 	// 		UserLikes:      recent.Twitter.Analytics.LikedCount,
// 	// 		UserTweetCount: recent.Twitter.Analytics.TweetCount,
// 	// 	},
// 	// 	Instagram: domain.SocialmediaInstagram{
// 	// 		Comments:     recent.Instagram.Analytics.CommentsCount,
// 	// 		Followers:    recent.Instagram.Analytics.FollowersCount,
// 	// 		Following:    recent.Instagram.Analytics.FollowingCount,
// 	// 		Impressions:  recent.Instagram.Analytics.ImpressionsCount,
// 	// 		Likes:        recent.Instagram.Analytics.LikeCount,
// 	// 		PostCount:    recent.Instagram.Analytics.MediaCount,
// 	// 		ProfileViews: recent.Instagram.Analytics.ProfileViewsCount,
// 	// 		Reach:        recent.Instagram.Analytics.ReachCount,
// 	// 	},
// 	// }

// 	var historical domain.SocialMediaDatapoint
// 	if !isNew {
// 		if socmedType == "facebook" {
// 			historical = domain.SocialMediaDatapoint{
// 				Likes:           recent.Facebook.Analytics.FanCount - lastTotal.Total.Likes,
// 				Followers:       recent.Facebook.Analytics.FollowersCount - lastTotal.Total.Followers,
// 				Engagement:      recent.Facebook.Analytics.PageEngagedUsers - lastTotal.Total.Engagement,
// 				Impressions:     recent.Facebook.Analytics.PageImpressions - lastTotal.Total.Impressions,
// 				ImpressionsPaid: recent.Facebook.Analytics.PageImpressionsPaid - lastTotal.Total.ImpressionsPaid,
// 				VideoViewTime:   recent.Facebook.Analytics.PageVideoViewTime - lastTotal.Total.VideoViewTime,
// 				VideoViews:      recent.Facebook.Analytics.PageVideoViews - lastTotal.Total.VideoViews,
// 				VideoViewsPaid:  recent.Facebook.Analytics.PageVideoViewsPaid - lastTotal.Total.VideoViewsPaid,
// 				Reactions:       recent.Facebook.Analytics.Reactions.Total - lastTotal.Total.Reactions,
// 			}
// 		} else if socmedType == "twitter" {
// 			historical = domain.SocialMediaDatapoint{
// 				Followers:      recent.Twitter.Analytics.FollowersCount - lastTotal.Total.Followers,
// 				Following:      recent.Twitter.Analytics.FriendsCount - lastTotal.Total.Following,
// 				UserLikes:      recent.Twitter.Analytics.LikedCount - lastTotal.Total.UserLikes,
// 				UserTweetCount: recent.Twitter.Analytics.TweetCount - lastTotal.Total.UserTweetCount,
// 			}
// 		} else if socmedType == "instagram" {
// 			historical = domain.SocialMediaDatapoint{
// 				Comments:     recent.Instagram.Analytics.CommentsCount - lastTotal.Total.Comments,
// 				Followers:    recent.Instagram.Analytics.FollowersCount - lastTotal.Total.Followers,
// 				Following:    recent.Instagram.Analytics.FollowingCount - lastTotal.Total.Following,
// 				Impressions:  recent.Instagram.Analytics.ImpressionsCount - lastTotal.Total.Impressions,
// 				Likes:        recent.Instagram.Analytics.LikeCount - lastTotal.Total.Likes,
// 				PostCount:    recent.Instagram.Analytics.MediaCount - lastTotal.Total.PostCount,
// 				ProfileViews: recent.Instagram.Analytics.ProfileViewsCount - lastTotal.Total.ProfileViews,
// 				Reach:        recent.Instagram.Analytics.ReachCount - lastTotal.Total.Reach,
// 			}
// 		}
// 	}

// 	return domain.SocialmediaHistorical{
// 		ID:            strings.ToLower(randstr.String(10)),
// 		SocialMediaId: socialmediaId,
// 		Platform:      socmedType,
// 		Historical:    historical,
// 		CreatedAt:     time.Now(),
// 	}
// }

// func AnalyticWebToTotalDomain(socialmediaId string, recent web.AyrshareAnalyticsResponse, socmedType string, createdAt, updatedAt time.Time) domain.SocialmediaLastTotal {
// 	var historical domain.SocialMediaDatapoint
// 	if socmedType == "facebook" {
// 		historical = domain.SocialMediaDatapoint{
// 			Likes:           recent.Facebook.Analytics.FanCount,
// 			Followers:       recent.Facebook.Analytics.FollowersCount,
// 			Engagement:      recent.Facebook.Analytics.PageEngagedUsers,
// 			Impressions:     recent.Facebook.Analytics.PageImpressions,
// 			ImpressionsPaid: recent.Facebook.Analytics.PageImpressionsPaid,
// 			VideoViewTime:   recent.Facebook.Analytics.PageVideoViewTime,
// 			VideoViews:      recent.Facebook.Analytics.PageVideoViews,
// 			VideoViewsPaid:  recent.Facebook.Analytics.PageVideoViewsPaid,
// 			Reactions:       recent.Facebook.Analytics.Reactions.Total,
// 		}
// 	} else if socmedType == "twitter" {
// 		historical = domain.SocialMediaDatapoint{
// 			Followers:      recent.Twitter.Analytics.FollowersCount,
// 			Following:      recent.Twitter.Analytics.FriendsCount,
// 			UserLikes:      recent.Twitter.Analytics.LikedCount,
// 			UserTweetCount: recent.Twitter.Analytics.TweetCount,
// 		}
// 	} else if socmedType == "instagram" {
// 		historical = domain.SocialMediaDatapoint{
// 			Comments:     recent.Instagram.Analytics.CommentsCount,
// 			Followers:    recent.Instagram.Analytics.FollowersCount,
// 			Following:    recent.Instagram.Analytics.FollowingCount,
// 			Impressions:  recent.Instagram.Analytics.ImpressionsCount,
// 			Likes:        recent.Instagram.Analytics.LikeCount,
// 			PostCount:    recent.Instagram.Analytics.MediaCount,
// 			ProfileViews: recent.Instagram.Analytics.ProfileViewsCount,
// 			Reach:        recent.Instagram.Analytics.ReachCount,
// 		}
// 	}

// 	return domain.SocialmediaLastTotal{
// 		ID:        socialmediaId,
// 		Platform:  socmedType,
// 		Total:     historical,
// 		CreatedAt: createdAt,
// 		UpdatedAt: updatedAt,
// 	}
// }

// // Logistic
// func ToLogisticResponse(logistic domain.Logistic) web.LogisticResponse {
// 	var linkImage string
// 	if logistic.ImageLogistic != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), logistic.ImageLogistic)
// 	}
// 	user := ToUserLogistic(logistic.User)
// 	return web.LogisticResponse{
// 		Id:            logistic.Id,
// 		Nama:          logistic.Nama,
// 		CategoryId:    logistic.CategoryId,
// 		Location:      logistic.Location,
// 		ImageLogistic: linkImage,
// 		Longitude:     logistic.Longitude,
// 		Latitude:      logistic.Latitude,
// 		UserId:        logistic.UserId,
// 		CreatedAt:     logistic.CreatedAt,
// 		UpdatedAt:     logistic.UpdatedAt,
// 		Category:      web.LogisticCategory(logistic.Category),
// 		User:          user,
// 	}
// }

// func ToLogisticResponses(logistic []domain.Logistic) []web.LogisticResponse {
// 	var logisticResponses []web.LogisticResponse
// 	for _, logistic := range logistic {
// 		logisticResponses = append(logisticResponses, ToLogisticResponse(logistic))
// 	}
// 	return logisticResponses
// }

// func ToUserLogistic(user domain.User) web.LogisticUser {
// 	return web.LogisticUser{
// 		Id:   user.Id,
// 		Name: user.Name,
// 	}
// }

// func ToLogisticSumResponse(logistic domain.ProgressLogistic) web.ProgressLogistic {
// 	return web.ProgressLogistic{
// 		RecruitProgress: logistic.RecruitProgress,
// 		RecruitGoal:     logistic.RecruitGoal,
// 	}
// }

// // Issue
// func ToIssueResponse(issue domain.Issue) web.IssueResponse {
// 	years := ToIssueYearResponses(issue.Years)
// 	return web.IssueResponse{
// 		ID:    issue.ID,
// 		Label: issue.Label,
// 		Value: issue.Value,
// 		Color: issue.Color,
// 		Years: years,
// 	}
// }

// // Issue PinPoint
// func ToIssuePinPointResponse(issue domain.IssueResponsePoints) web.IssueResponsePoints {
// 	return web.IssueResponsePoints{
// 		ID:         issue.ID,
// 		Year:       issue.Year,
// 		ProvinceId: issue.ProvinceId,
// 		SubIssueId: issue.SubIssueId,
// 		Data:       web.DataPinPoint(issue.Data),
// 	}
// }

// // Issue Rank
// func ToIssueResponseRank(issue web.IssueDataResponseRank) web.IssueDataResponseRank {
// 	return web.IssueDataResponseRank{
// 		Wilayah:   issue.Wilayah,
// 		ID:        issue.ID,
// 		Rank:      issue.Rank,
// 		MetaData:  issue.MetaData,
// 		SubIssues: issue.SubIssues,
// 	}
// }

// // IssueData
// func ToIssueDataResponse(issue domain.IssueData) web.IssueDataResponse {
// 	return web.IssueDataResponse{
// 		ID:        issue.ID,
// 		IssueType: issue.IssueType,
// 		Year:      issue.Year,
// 	}
// }

// // IssueDataResponse
// func ToIssueDataResponses(issue domain.IssueData) domain.IssueData {
// 	return domain.IssueData{
// 		ID:        issue.ID,
// 		IssueType: issue.IssueType,
// 		Year:      issue.Year,
// 		MetaData:  issue.MetaData,
// 		Data:      issue.Data,
// 	}
// }

// func ToIssueResponseRanks(issues []web.IssueDataResponseRank) []web.IssueDataResponseRank {
// 	var issuesResponse []web.IssueDataResponseRank
// 	for _, issue := range issues {
// 		issuesResponse = append(issuesResponse, ToIssueResponseRank(issue))
// 	}
// 	return issuesResponse
// }

// func ToIssueResponses(issues []domain.Issue) []web.IssueResponse {
// 	var issuesResponse []web.IssueResponse
// 	for _, issue := range issues {
// 		issuesResponse = append(issuesResponse, ToIssueResponse(issue))
// 	}
// 	return issuesResponse
// }

// func ToIssuePinPointsResponses(issues []domain.IssueResponsePoints) []web.IssueResponsePoints {
// 	var issuesResponse []web.IssueResponsePoints
// 	for _, issue := range issues {
// 		issuesResponse = append(issuesResponse, ToIssuePinPointResponse(issue))
// 	}
// 	return issuesResponse
// }

// // Issue Year
// func ToIssueYearResponse(issueYear domain.IssueYear) web.IssueYearResponse {
// 	issueSubs := ToIssueSubResponses(issueYear.IssueSubs)
// 	return web.IssueYearResponse{
// 		ID:        issueYear.ID,
// 		Label:     issueYear.Year,
// 		Value:     issueYear.Year,
// 		IssueSubs: issueSubs,
// 	}
// }

// func ToIssueYearResponses(issuesYear []domain.IssueYear) []web.IssueYearResponse {
// 	var issuesYearResponse []web.IssueYearResponse
// 	for _, issueYear := range issuesYear {
// 		issuesYearResponse = append(issuesYearResponse, ToIssueYearResponse(issueYear))
// 	}
// 	return issuesYearResponse
// }

// // Issue Sub
// func ToIssueSubResponse(issue domain.IssueSub) web.IssueSubResponse {
// 	var linkImage string
// 	if issue.ImageUrl != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), issue.ImageUrl)
// 	}

// 	return web.IssueSubResponse{
// 		ID:       issue.ID,
// 		IssueID:  issue.IssueID,
// 		Label:    issue.Label,
// 		Value:    issue.Value,
// 		ImageUrl: linkImage,
// 	}
// }

// func ToIssueSubResponses(issues []domain.IssueSub) []web.IssueSubResponse {
// 	var issuesResponse []web.IssueSubResponse
// 	for _, issue := range issues {
// 		issuesResponse = append(issuesResponse, ToIssueSubResponse(issue))
// 	}
// 	return issuesResponse
// }

// // Permission
// func ToPermissionResponse(permission domain.Permission) web.PermissionResponse {
// 	return web.PermissionResponse{
// 		ID:    permission.ID,
// 		Label: permission.Label,
// 		Value: permission.Value,
// 	}
// }

// func ToPermissionResponses(permissions []domain.Permission) []web.PermissionResponse {
// 	var permissionsResponse []web.PermissionResponse
// 	for _, permission := range permissions {
// 		permissionsResponse = append(permissionsResponse, ToPermissionResponse(permission))
// 	}
// 	return permissionsResponse
// }

// // Role
// func ToRoleResponse(role domain.Role) web.RoleResponse {
// 	return web.RoleResponse{
// 		ID:         role.ID,
// 		Label:      role.Label,
// 		Value:      role.Value,
// 		Permission: role.Permission,
// 	}
// }

// func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
// 	var rolesResponse []web.RoleResponse
// 	for _, role := range roles {
// 		rolesResponse = append(rolesResponse, ToRoleResponse(role))
// 	}
// 	return rolesResponse
// }

// // Access
// func ToAccessResponse(access domain.Access) web.AccessResponse {
// 	return web.AccessResponse{
// 		ID:    access.ID,
// 		Label: access.Label,
// 		Value: access.Value,
// 	}
// }

// func ToAccessResponses(accesses []domain.Access) []web.AccessResponse {
// 	var accessesResponse []web.AccessResponse
// 	for _, access := range accesses {
// 		accessesResponse = append(accessesResponse, ToAccessResponse(access))
// 	}
// 	return accessesResponse
// }

// // Talkwalker
// func ToTalkwalkerResponse(issue domain.Talkwalker) web.TalkwalkerResponse {
// 	var linkImage string
// 	if issue.FileUrl != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), issue.FileUrl)
// 	}

// 	return web.TalkwalkerResponse{
// 		ID:        issue.ID,
// 		Name:      issue.Name,
// 		FileUrl:   linkImage,
// 		CreatedAt: issue.CreatedAt,
// 	}
// }

// func ToTalkwalkerResponses(issues []domain.Talkwalker) []web.TalkwalkerResponse {
// 	var issuesResponse []web.TalkwalkerResponse
// 	for _, issue := range issues {
// 		issuesResponse = append(issuesResponse, ToTalkwalkerResponse(issue))
// 	}
// 	return issuesResponse
// }

// // IssueManagementData
// func ToIssueManagementDataResponse(issue domain.IssueManagementData) web.IssueManagementDataResponse {
// 	var linkImage string
// 	if issue.FileUrl != "" {
// 		linkImage = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PREFIX"), issue.FileUrl)
// 	}

// 	return web.IssueManagementDataResponse{
// 		ID:        issue.ID,
// 		Name:      issue.Name,
// 		FileUrl:   linkImage,
// 		Year:      issue.Year,
// 		CreatedAt: issue.CreatedAt,
// 		IssueType: issue.IssueType,
// 		SubIssue:  issue.SubIssue,
// 	}
// }

// func ToIssueManagementDataResponses(issues []domain.IssueManagementData) []web.IssueManagementDataResponse {
// 	var issuesResponse []web.IssueManagementDataResponse
// 	for _, issue := range issues {
// 		issuesResponse = append(issuesResponse, ToIssueManagementDataResponse(issue))
// 	}
// 	return issuesResponse
// }

// Product
func ToProductResponse(permission domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:   permission.ID,
		Name: permission.Name,
	}
}

func ToProductResponses(permissions []domain.Product) []web.ProductResponse {
	var permissionsResponse []web.ProductResponse
	for _, permission := range permissions {
		permissionsResponse = append(permissionsResponse, ToProductResponse(permission))
	}
	return permissionsResponse
}
