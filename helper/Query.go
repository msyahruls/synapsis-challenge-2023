package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func QuerySurvey(searchSurveyName, status string, pipeline *mongo.Pipeline) {

// 	if searchSurveyName != "" {
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "survey_name", Value: primitive.Regex{Pattern: searchSurveyName, Options: "i"}}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}

// 	statuInt, _ := strconv.Atoi(status)

// 	if status != "" {
// 		matchStatusStage := bson.D{{Key: "$match", Value: bson.D{{Key: "status", Value: statuInt}}}}
// 		*pipeline = append(*pipeline, matchStatusStage)
// 	}

// }

// func QueryResponse(surveyId, respondentId, recruiterId string, pipeline *mongo.Pipeline) {

// 	if surveyId != "" {
// 		matchQuestionIdStage := bson.D{{Key: "$match", Value: bson.D{{Key: "survey_id", Value: surveyId}}}}
// 		*pipeline = append(*pipeline, matchQuestionIdStage)
// 	}

// 	if respondentId != "" {
// 		matchQuestionIdStage := bson.D{{Key: "$match", Value: bson.D{{Key: "respondent_id", Value: respondentId}}}}
// 		*pipeline = append(*pipeline, matchQuestionIdStage)
// 	}
// 	if recruiterId != "" {
// 		matchQuestionIdStage := bson.D{{Key: "$match", Value: bson.D{{Key: "recruiter_id", Value: recruiterId}}}}
// 		*pipeline = append(*pipeline, matchQuestionIdStage)
// 	}

// }

// func QueryQuestion(surveyId string) primitive.M {
// 	query := bson.M{}

// 	if surveyId != "" {
// 		query["survey_id"] = surveyId
// 	}

// 	return query
// }

// func QueryOption(questionId string) primitive.M {
// 	query := bson.M{}

// 	if questionId != "" {
// 		query["question_id"] = questionId
// 	}

// 	return query
// }

// func QueryAnswer(questionId, responseId string, pipeline *mongo.Pipeline) {
// 	if questionId != "" {
// 		matchQuestionIdStage := bson.D{{Key: "$match", Value: bson.D{{Key: "question_id", Value: questionId}}}}
// 		*pipeline = append(*pipeline, matchQuestionIdStage)
// 	}
// 	if responseId != "" {
// 		matchResponseIdStage := bson.D{{Key: "$match", Value: bson.D{{Key: "response_id", Value: responseId}}}}
// 		*pipeline = append(*pipeline, matchResponseIdStage)
// 	}
// }

// func QueryUser(level, districtId, searchName string, pipeline *mongo.Pipeline) {

// 	if searchName != "" {
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "name", Value: primitive.Regex{Pattern: searchName, Options: "i"}}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}

// 	if level != "" {
// 		intLevel, _ := strconv.ParseInt(level, 10, 32)
// 		matchLevelStage := bson.D{{Key: "$match", Value: bson.D{{Key: "occupation.level", Value: intLevel}}}}
// 		*pipeline = append(*pipeline, matchLevelStage)
// 	}

// 	if districtId != "" {
// 		matchDistrictStage := bson.D{{Key: "$match", Value: bson.D{{Key: "district_id", Value: districtId}}}}
// 		*pipeline = append(*pipeline, matchDistrictStage)
// 	}
// }

// func QueryProvince(name string) primitive.M {
// 	query := bson.M{}

// 	if name != "" {
// 		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
// 	}

// 	return query
// }

// func QueryRegency(provinceId, name string) primitive.M {
// 	query := bson.M{}

// 	if provinceId != "" {
// 		id, _ := strconv.Atoi(provinceId)
// 		query["province_id"] = id
// 	}

// 	if name != "" {
// 		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
// 	}

// 	return query
// }

// func QueryDistricts(regencyId, name string) primitive.M {
// 	query := bson.M{}

// 	if regencyId != "" {
// 		id, _ := strconv.Atoi(regencyId)
// 		query["province_id"] = id
// 	}

// 	if name != "" {
// 		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
// 	}

// 	return query
// }

// func QueryVillage(districtId, name string) primitive.M {
// 	query := bson.M{}

// 	if districtId != "" {
// 		id, _ := strconv.Atoi(districtId)
// 		query["province_id"] = id
// 	}

// 	if name != "" {
// 		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
// 	}

// 	return query
// }

// func QueryVillageAggregat(name string, pipeline *mongo.Pipeline) {

// 	if name != "" {
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "name", Value: primitive.Regex{Pattern: name, Options: "i"}}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}

// }

// func QueryFilterDistrict(districtId string, pipeline *mongo.Pipeline) {
// 	if districtId != "" {
// 		matchDistrictStage := bson.D{{Key: "$match", Value: bson.D{{Key: "district_id", Value: districtId}}}}
// 		*pipeline = append(*pipeline, matchDistrictStage)
// 	}
// }

// func QueryUserDpt(pollstationId, villageId, searchName string, pipeline *mongo.Pipeline) {

// 	if searchName != "" {
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "name", Value: primitive.Regex{Pattern: searchName, Options: "i"}}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}

// 	if villageId != "" {
// 		matchLevelStage := bson.D{{Key: "$match", Value: bson.D{{Key: "village_id", Value: villageId}}}}
// 		*pipeline = append(*pipeline, matchLevelStage)
// 	}

// 	if pollstationId != "" {
// 		matchDistrictStage := bson.D{{Key: "$match", Value: bson.D{{Key: "polling_station_id", Value: pollstationId}}}}
// 		*pipeline = append(*pipeline, matchDistrictStage)
// 	}
// }

// func QueryRecruitUserId(userId string, pipeline *mongo.Pipeline) {

// 	if userId != "" {
// 		matchUserStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userId}}}}
// 		*pipeline = append(*pipeline, matchUserStage)
// 	}
// }

// func QueryRecruit(searchName, level string, pipeline *mongo.Pipeline) {
// 	if searchName != "" {
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user.name", Value: primitive.Regex{Pattern: searchName, Options: "i"}}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}

// 	if level != "" {
// 		intlevel, err := strconv.Atoi(level)
// 		if err != nil {
// 			panic(exception.NewError(400, "error in query level"))
// 		}
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user.occupation.level", Value: intlevel}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}
// }

// func QueryComplaint(userid, complaintStatus, categoryId string, pipeline *mongo.Pipeline) {
// 	if userid != "" {
// 		matchUser := bson.D{{Key: "$match", Value: bson.D{{Key: "sender_id", Value: userid}}}}
// 		*pipeline = append(*pipeline, matchUser)
// 	}

// 	if complaintStatus != "" {
// 		matchStatus := bson.D{{Key: "$match", Value: bson.D{{Key: "complaint_status_id", Value: complaintStatus}}}}
// 		*pipeline = append(*pipeline, matchStatus)
// 	}

// 	if categoryId != "" {
// 		matchCategory := bson.D{{Key: "$match", Value: bson.D{{Key: "category_id", Value: categoryId}}}}
// 		*pipeline = append(*pipeline, matchCategory)
// 	}
// }

// func QueryProgresSurvey(searchName string, pipeline *mongo.Pipeline) {
// 	if searchName != "" {
// 		matchNameStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user.name", Value: primitive.Regex{Pattern: searchName, Options: "i"}}}}}
// 		*pipeline = append(*pipeline, matchNameStage)
// 	}
// }

// func QueryProgressSurveyUserId(userId string, pipeline *mongo.Pipeline) {

// 	if userId != "" {
// 		matchUserStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userId}}}}
// 		*pipeline = append(*pipeline, matchUserStage)
// 	}
// }

// func QueryCandidate(partaiid, electionTypeId, daerahPemiluId string, pipeline *mongo.Pipeline) {

// 	if partaiid != "" {
// 		matchUserStage := bson.D{{Key: "$match", Value: bson.D{{Key: "partai_id", Value: partaiid}}}}
// 		*pipeline = append(*pipeline, matchUserStage)
// 	}

// 	if electionTypeId != "" {
// 		matchUserStage := bson.D{{Key: "$match", Value: bson.D{{Key: "election_type_id", Value: electionTypeId}}}}
// 		*pipeline = append(*pipeline, matchUserStage)
// 	}

// 	if daerahPemiluId != "" {
// 		matchUserStage := bson.D{{Key: "$match", Value: bson.D{{Key: "daerah_pemilu_id", Value: daerahPemiluId}}}}
// 		*pipeline = append(*pipeline, matchUserStage)
// 	}
// }

// func QuerySurveyResult(questionId, villageId, districtId, regencyId string, pipeline *mongo.Pipeline) {

// 	if questionId != "" {
// 		matchQuestionStage := bson.D{{Key: "$match", Value: bson.D{{Key: "questions._id", Value: questionId}}}}
// 		*pipeline = append(*pipeline, matchQuestionStage)
// 	}

// 	if villageId != "" {
// 		matchVillageStage := bson.D{{Key: "$match", Value: bson.D{{Key: "village._id", Value: villageId}}}}
// 		*pipeline = append(*pipeline, matchVillageStage)
// 	}

// 	if districtId != "" {
// 		matchDistrictStage := bson.D{{Key: "$match", Value: bson.D{{Key: "district._id", Value: districtId}}}}
// 		*pipeline = append(*pipeline, matchDistrictStage)
// 	}

// 	if regencyId != "" {
// 		matchRegencyStage := bson.D{{Key: "$match", Value: bson.D{{Key: "regency._id", Value: regencyId}}}}
// 		*pipeline = append(*pipeline, matchRegencyStage)
// 	}

// }

// func QueryLogistic(userid, categoryId string, pipeline *mongo.Pipeline) {
// 	if userid != "" {
// 		matchUser := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userid}}}}
// 		*pipeline = append(*pipeline, matchUser)
// 	}

// 	if categoryId != "" {
// 		matchCategory := bson.D{{Key: "$match", Value: bson.D{{Key: "category_id", Value: categoryId}}}}
// 		*pipeline = append(*pipeline, matchCategory)
// 	}
// }

// // Event
// func QueryEvent(status string) primitive.M {
// 	query := bson.M{}

// 	if status != "" {
// 		if status == "true" {
// 			query["status"] = true
// 		} else if status == "false" {
// 			query["status"] = false
// 		}

// 	}

// 	return query
// }

// // Issue
// func QueryIssue(issueId string, label string, value string) primitive.M {
// 	query := bson.M{}

// 	if issueId != "" {
// 		query["_id"] = primitive.Regex{Pattern: issueId, Options: "i"}
// 	}

// 	if label != "" {
// 		query["label"] = primitive.Regex{Pattern: label, Options: "i"}
// 	}

// 	if value != "" {
// 		query["value"] = primitive.Regex{Pattern: value, Options: "i"}
// 	}

// 	return query
// }

// // Issue Sub
// func QueryIssueSub(issueId string, label string, value string) primitive.M {
// 	query := bson.M{}

// 	if issueId != "" {
// 		query["issue_id"] = primitive.Regex{Pattern: issueId, Options: "i"}
// 	}

// 	if label != "" {
// 		query["label"] = primitive.Regex{Pattern: label, Options: "i"}
// 	}

// 	if value != "" {
// 		query["value"] = primitive.Regex{Pattern: value, Options: "i"}
// 	}

// 	return query
// }

// // Issue Exist
// func QueryIssueExist(label string, value string) primitive.M {
// 	query := bson.M{
// 		"$or": []interface{}{
// 			bson.M{"label": label},
// 			bson.M{"value": value},
// 		},
// 	}

// 	return query
// }

// // Issue Year
// func QueryIssueYear(year string, issueType string) primitive.M {
// 	query := bson.M{}

// 	if year != "" {
// 		query["year"] = primitive.Regex{Pattern: year, Options: "i"}
// 	}

// 	if issueType != "" {
// 		query["issue_type"] = primitive.Regex{Pattern: issueType, Options: "i"}
// 	}

// 	return query
// }

// // Talkwalker
// func QueryTalkwalker(name string) primitive.M {
// 	query := bson.M{}

// 	if name != "" {
// 		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
// 	}

// 	return query
// }

// // Issue Management Data
// func QueryIssueManagementData(name string, issue string, sub_issue string, year string, created_at string) primitive.M {
// 	query := bson.M{}

// 	if name != "" {
// 		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
// 	}

// 	if issue != "" {
// 		query["issue"] = primitive.Regex{Pattern: issue, Options: "i"}
// 	}

// 	if sub_issue != "" {
// 		query["sub_issue"] = primitive.Regex{Pattern: sub_issue, Options: "i"}
// 	}

// 	if year != "" {
// 		query["year"] = primitive.Regex{Pattern: year, Options: "i"}
// 	}

// 	if created_at != "" {
// 		query["created_at"] = primitive.Regex{Pattern: created_at, Options: "i"}
// 	}

// 	return query
// }

// // Issue Log
// func QueryIssueLog(provinceId, regencyId, issueType string, subIssue string, year string) primitive.M {
// 	query := bson.M{}

// 	// query := bson.M{"_id": provinceIdInt}

// 	if provinceId != "" {
// 		parseInt, _ := strconv.Atoi(provinceId)
// 		query["province_id"] = parseInt
// 	}

// 	if regencyId != "" {
// 		parseInt, _ := strconv.Atoi(regencyId)
// 		query["province_id"] = parseInt
// 	}

// 	if issueType != "" {
// 		query["data.type_issue"] = primitive.Regex{Pattern: issueType, Options: "i"}
// 	}

// 	if subIssue != "" {
// 		query["data.sub_issue"] = primitive.Regex{Pattern: subIssue, Options: "i"}
// 	}

// 	if year != "" {
// 		query["year"] = primitive.Regex{Pattern: year, Options: "i"}
// 	}

// 	return query
// }

// Category
func QueryCategory(issueId string, name string) primitive.M {
	query := bson.M{}

	if issueId != "" {
		query["_id"] = primitive.Regex{Pattern: issueId, Options: "i"}
	}

	if name != "" {
		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
	}

	return query
}

// Category Exist
func QueryCategoryExist(id string, name string) primitive.M {
	query := bson.M{
		"$or": []interface{}{
			bson.M{"_id": id},
			bson.M{"name": name},
		},
	}

	return query
}

// Product
func QueryProduct(issueId string, name string, category string) primitive.M {
	query := bson.M{}

	if issueId != "" {
		query["_id"] = primitive.Regex{Pattern: issueId, Options: "i"}
	}

	if name != "" {
		query["name"] = primitive.Regex{Pattern: name, Options: "i"}
	}

	if category != "" {
		query["category_id"] = primitive.Regex{Pattern: category, Options: "i"}
	}

	return query
}

// Product Exist
func QueryProductExist(name string) primitive.M {
	query := bson.M{
		"$or": []interface{}{
			bson.M{"name": name},
		},
	}

	return query
}

// Cart
func QueryCart(UserID string, ProductID string) primitive.M {
	query := bson.M{}

	if UserID != "" {
		query["user_id"] = primitive.Regex{Pattern: UserID, Options: "i"}
	}

	if ProductID != "" {
		query["product_id"] = primitive.Regex{Pattern: ProductID, Options: "i"}
	}

	return query
}

// Cart Exist
func QueryCartExist(UserID string, ProductID string) primitive.M {
	// query := bson.M{
	// 	"$or": []interface{}{
	// 		bson.M{"_id": id},
	// 		bson.M{"user_id": UserID},
	// 		bson.M{"prodcut_id": ProductID},
	// 	},
	// }

	query := bson.M{
		"$and": []interface{}{
			bson.M{"user_id": UserID},
			bson.M{"prodcut_id": ProductID},
		},
	}

	return query
}
