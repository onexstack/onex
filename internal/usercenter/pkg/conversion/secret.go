package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/usercenter/model"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// SecretMToSecretV1 converts a SecretM object from the internal model
// to a Secret object in the v1 API format.
func SecretMToSecretV1(secretModel *model.SecretM) *v1.Secret {
	var secret v1.Secret
	_ = core.CopyWithConverters(&secret, secretModel)
	return &secret
}

// SecretV1ToSecretM converts a Secret object from the v1 API format
// to a SecretM object in the internal model.
func SecretV1ToSecretM(secret *v1.Secret) *model.SecretM {
	var secretModel model.SecretM
	_ = core.CopyWithConverters(&secretModel, secret)
	return &secretModel
}
