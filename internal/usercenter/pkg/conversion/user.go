package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/usercenter/model"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// UserMToUserV1 converts a UserM object from the internal model
// to a User object in the v1 API format.
func UserMToUserV1(userModel *model.UserM) *v1.User {
	var user v1.User
	_ = core.CopyWithConverters(&user, userModel)
	return &user
}

// UserV1ToUserM converts a User object from the v1 API format
// to a UserM object in the internal model.
func UserV1ToUserM(user *v1.User) *model.UserM {
	var userModel model.UserM
	_ = core.CopyWithConverters(&userModel, user)
	return &userModel
}
