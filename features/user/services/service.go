package services

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	user "synapsis/features/user"
	helper "synapsis/helper"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {

	res, err := uuc.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := helper.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password tidak sesuai " + res.Password)
	}

	token, _ := helper.GenerateJWT(int(res.ID))

	return token, res, nil

}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {

	err := uuc.vld.Struct(newUser)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return user.Core{}, errors.New("validation error")
	}
	hashed, err := helper.GeneratePassword(newUser.Password)

	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}

	newUser.Password = string(hashed)
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(file multipart.FileHeader, token interface{}, updateData user.Core) (user.Core, error) {
	if updateData.Password != "" {
		hashed, err := helper.GeneratePassword(updateData.Password)
		if err != nil {
			log.Println("bcrypt error ", err.Error())
			return user.Core{}, errors.New("password process error")
		}
		updateData.Password = string(hashed)
	}

	id := helper.ExtractToken(token)

	if file.Size > 5000000 {
		return user.Core{}, errors.New("file size is too big")
	}

	formFile, err := file.Open()
	if err != nil {
		return user.Core{}, errors.New("open formheader error")
	}

	if !helper.TypeFile(formFile) {
		return user.Core{}, errors.New("use jpg or png type file")
	}
	defer formFile.Close()
	formFile, _ = file.Open()
	uploadUrl, err := helper.NewMediaUpload().AvatarUpload(helper.Avatar{Avatar: formFile})

	if err != nil {
		return user.Core{}, errors.New("server error")
	}

	updateData.Avatar = uploadUrl

	res, err := uuc.qry.Update(uint(id), updateData)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else if strings.Contains(err.Error(), "not valid") {
			msg = "format tidak sesuai"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Delete(token interface{}) (user.Core, error) {

	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("id user not found")
	}
	data, err := uuc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}
	return data, nil

}
