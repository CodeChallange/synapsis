package handler

import "synapsis/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Avatar   string `json:"avatar" form:"avatar"`
	HP       string `json:"hp" form:"hp"`
	Birth    string `json:"birth" form:"birth"`
}

type DeleteRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Username = cnv.Username
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Username = cnv.Username
		res.Password = cnv.Password
		res.Avatar = cnv.Avatar
		res.HP = cnv.HP
		res.Birth = cnv.Birth
	case DeleteRequest:
		cnv := data.(DeleteRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Username = cnv.Username
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
