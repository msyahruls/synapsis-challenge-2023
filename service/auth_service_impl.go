package service

import (
	"fmt"
	"strings"
	"time"

	"synapsis-challange/exception"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"
	"synapsis-challange/model/web"
	"synapsis-challange/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thanhpk/randstr"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	// OccupationRepository    repository.OccupationRepository
	// RecruitRepository       repository.RecruitRepository
	// RecruitedUserRepository repository.RecruitedUserRepository
	Validate *validator.Validate
}

type AuthService interface {
	Register(request web.RegisterUserRequest) web.UserResponse
	Login(request web.LoginRequest) (fiber.Cookie, web.LoginResponse)
	Logout() fiber.Cookie

	// ForgetPasswordEmail(email string)
	// ForgetPasswordPhone(phone string)
	// ResetPassword(email, phone, token string, request web.ResetPassword)
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Register(request web.RegisterUserRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// if !helper.IsNumeric(request.NIK) {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "NIK harus angka."))
	// }

	// if !helper.IsNumeric(request.Phone) {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Nomer telepon harus angka."))
	// }

	// if len([]rune(request.NIK)) != 16 {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "NIK harus 16 digit."))
	// }

	// if len([]rune(request.Phone)) < 10 || len([]rune(request.Phone)) > 13 {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Nomer telepon harus 10-13 digit"))
	// }

	// if nik, _ := service.AuthRepository.FindByQuery("nik", request.NIK); nik.Id != "" {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "NIK sudah terdaftar."))
	// }
	// if phone, _ := service.AuthRepository.FindByQuery("phone", request.Phone); phone.Id != "" {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Nomor telepon sudah terdaftar."))
	// }

	if email, _ := service.AuthRepository.FindByQuery("email", request.Email); email.ID != "" {
		panic(exception.NewError(fiber.StatusBadRequest, "Email sudah terdaftar."))
	}

	if err != nil {
		panic(exception.NewError(fiber.StatusInternalServerError, "internal server error"))
	}

	user := domain.User{
		ID:        strings.ToLower(randstr.String(10)),
		Name:      request.Name,
		Email:     request.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		//OccupationId: Occupation.ID,
		// NIK:    request.NIK,
		// Phone:  request.Phone,
		// Gender: request.Gender,
		//ReferralCode: randstr.String(7),
		// Longitude: request.Longitude,
		// Latitude:  request.Latitude,
	}

	user.SetPassword(request.Password)

	// //check refferal
	// if request.ReferralCode != "" {
	// 	referalUser, _ := service.AuthRepository.FindByQuery("referral_code", request.ReferralCode)

	// 	if referalUser.Id == "" {
	// 		panic(exception.NewError(fiber.StatusBadRequest, "Kode referral tidak valid."))
	// 	}

	// 	// check if level referral
	// 	OccupationReferral, _ := service.OccupationRepository.FindById(referalUser.OccupationId)

	// 	// if referral from admin > user koordinator
	// 	// if referral from koor > user relawan
	// 	// if referral from relawan > user general
	// 	if OccupationReferral.Level == 1 {
	// 		koordinatorUser, _ := strconv.Atoi(helper.LevelKoordinator)
	// 		Occupation, _ := service.OccupationRepository.FindByLevel("level", koordinatorUser)
	// 		user.OccupationId = Occupation.ID
	// 		user.Occupation = Occupation
	// 		user.ReferralCode = randstr.String(7)
	// 	} else if OccupationReferral.Level == 2 {
	// 		relawanUser, _ := strconv.Atoi(helper.LevelRelawan)
	// 		Occupation, _ := service.OccupationRepository.FindByLevel("level", relawanUser)
	// 		user.OccupationId = Occupation.ID
	// 		user.Occupation = Occupation
	// 		user.ReferralCode = randstr.String(7)
	// 	} else if OccupationReferral.Level == 3 {
	// 		generalUser, _ := strconv.Atoi(helper.LevelGeneralUser)
	// 		Occupation, _ := service.OccupationRepository.FindByLevel("level", generalUser)
	// 		user.OccupationId = Occupation.ID
	// 		user.Occupation = Occupation
	// 	} else {
	// 		generalUser, _ := strconv.Atoi(helper.LevelGeneralUser)
	// 		Occupation, _ := service.OccupationRepository.FindByLevel("level", generalUser)
	// 		user.OccupationId = Occupation.ID
	// 		user.Occupation = Occupation
	// 	}

	// 	// If referral admin or koordinator or Relawan
	// 	// Update Recruitment
	// 	if OccupationReferral.Level == 1 || OccupationReferral.Level == 2 || OccupationReferral.Level == 3 {
	// 		// update user recruit
	// 		userRecruitment, err := service.RecruitRepository.FindByUserId(referalUser.Id)

	// 		if err != nil || userRecruitment.ID == "" {
	// 			defaultGoal, _ := strconv.Atoi(helper.DefaultGoalRecruitment)

	// 			userRecruitment = domain.UserRecruit{
	// 				ID:              strings.ToLower(randstr.String(10)),
	// 				UserId:          referalUser.Id,
	// 				RecruitProgress: 1,
	// 				RecruitGoal:     defaultGoal,
	// 			}

	// 			service.RecruitRepository.Create(userRecruitment)
	// 		} else {
	// 			userRecruitment.RecruitProgress = userRecruitment.RecruitProgress + 1
	// 			service.RecruitRepository.Update(userRecruitment)
	// 		}

	// 		//add user to recruited user
	// 		recrutedUser := domain.RecruitedUser{
	// 			ID:            strings.ToLower(randstr.String(10)),
	// 			UserId:        user.Id,
	// 			UserRecruitId: userRecruitment.ID,
	// 			RecruiterId:   referalUser.Id,
	// 		}
	// 		service.RecruitedUserRepository.Create(recrutedUser)
	// 	}
	// } else {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Kode referral tidak boleh kosong"))
	// }

	service.AuthRepository.Register(user)

	// if user.Occupation.Level == 4 {
	// 	service.AuthRepository.CreateRespondent(user)
	// }

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) Login(request web.LoginRequest) (fiber.Cookie, web.LoginResponse) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// if request.Email == "" && request.Phone == "" {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Harap masukan email atau nomor telepon untuk login"))
	// }

	var user domain.User

	if request.Email != "" {
		user, err = service.AuthRepository.LoginEmail(request.Email)
		if err != nil {
			panic(exception.NewError(fiber.StatusInternalServerError, "server error"))
		}
	}
	//  else {
	// 	user, err = service.AuthRepository.LoginPhone(request.Phone)
	// 	if err != nil {
	// 		panic(exception.NewError(fiber.StatusInternalServerError, "server error"))
	// 	}
	// }

	if user.ID == "" {
		panic(exception.NewError(fiber.StatusNotFound, "Email tidak ditemukan."))
	}

	if err := user.ComparePassword(user.Password, request.Password); err != nil {
		panic(exception.NewError(fiber.StatusBadRequest, "Password salah."))
	}

	// occupation, _ := service.OccupationRepository.FindById(user.OccupationId)

	// if occupation.Level == 5 {
	// 	panic(exception.NewError(fiber.StatusUnauthorized, "User sudah di blacklist."))
	// }

	token, err := helper.GenerateJwt(user.ID, fmt.Sprintf("%d", 1))
	if err != nil {
		helper.PanicIfError(err)
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	resp := web.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		// OccupationLevel: occupation.Level,
		// OccupationName:  occupation.Name,
	}

	return cookie, resp
}

func (service *AuthServiceImpl) Logout() fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	return cookie
}

// func (service *AuthServiceImpl) ForgetPasswordEmail(email string) {
// 	_, err := service.AuthRepository.FindByQuery("email", email)

// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusBadRequest, "Email tidak terdaftar di database."))
// 	}

// 	var data domain.ResetPasswordToken

// 	token := strings.ToLower(randstr.String(30))
// 	data.Token = token
// 	data.Email = email
// 	data.CreatedAt = time.Now()

// 	err = helper.EmailSender(email, token)
// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusBadGateway, "request error"))
// 	}
// 	service.AuthRepository.CreateToken(data)
// }

// func (service *AuthServiceImpl) ForgetPasswordPhone(phone string) {
// 	_, err := service.AuthRepository.FindByQuery("phone", phone)

// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusBadRequest, "Nomor telepon tidak terdaftar di database."))
// 	}

// 	var data domain.ResetPasswordToken

// 	token := strings.ToLower(randstr.String(30))
// 	data.Token = token
// 	data.Phone = phone
// 	data.CreatedAt = time.Now()

// 	err = helper.WhatsappSender(phone, token)
// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusBadGateway, "request error"))
// 	}
// 	service.AuthRepository.CreateToken(data)
// }

// func (service *AuthServiceImpl) ResetPassword(email, phone, token string, request web.ResetPassword) {
// 	err := service.Validate.Struct(request)
// 	helper.PanicIfError(err)

// 	var decodedByte, _ = base64.StdEncoding.DecodeString(token)
// 	var resetToken = string(decodedByte)

// 	if request.Password != request.PasswordConfirm {
// 		panic(exception.NewError(fiber.StatusBadRequest, "Password konfirmasi tidak sama."))
// 	}

// 	checkToken, err := service.AuthRepository.CheckToken(resetToken)

// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusBadRequest, "token Invalid"))
// 	}

// 	if checkToken.Token != resetToken || checkToken.Email != email || checkToken.Phone != phone {
// 		panic(exception.NewError(fiber.StatusBadRequest, "token Invalid"))
// 	}

// 	var user domain.User

// 	if email != "" {
// 		user, err = service.AuthRepository.FindByQuery("email", email)
// 	} else {
// 		user, err = service.AuthRepository.FindByQuery("phone", phone)
// 	}

// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusNotFound, err.Error()))
// 	}

// 	user.UpdatedAt = time.Now()
// 	user.SetPassword(request.Password)

// 	service.AuthRepository.UpdatePassword(user)

// 	var encodedString = base64.StdEncoding.EncodeToString([]byte(user.Password))
// 	user.Password = encodedString

// 	err = service.AuthRepository.DeleteToken(resetToken)
// 	if err != nil {
// 		helper.PanicIfError(err)
// 	}
// }
