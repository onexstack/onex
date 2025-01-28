package auth

//go:generate mockgen -destination mock_auth.go -package auth onex/internal/usercenter/biz/v1/auth AuthBiz

import (
	"context"

	"github.com/onexstack/onexstack/pkg/authn"
	"github.com/onexstack/onexstack/pkg/i18n"
	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/store/where"

	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/internal/usercenter/pkg/auth"
	"github.com/onexstack/onex/internal/usercenter/pkg/locales"
	"github.com/onexstack/onex/internal/usercenter/store"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

const (
	// MaxErrGroupConcurrency defines the maximum concurrency level
	// for error group operations.
	MaxErrGroupConcurrency = 100
)

// AuthBiz defines the interface that contains methods for handling auth requests.
type AuthBiz interface {
	// Login authenticates a user and returns a token.
	Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error)

	// Logout invalidates a token.
	Logout(ctx context.Context, rq *v1.LogoutRequest) (*v1.LogoutResponse, error)

	// RefreshToken refreshes an existing token and returns a new one.
	RefreshToken(ctx context.Context, rq *v1.RefreshTokenRequest) (*v1.LoginReply, error)

	// Authenticate validates an access token and returns the associated user ID.
	Authenticate(ctx context.Context, accessToken string) (*v1.AuthenticateResponse, error)

	// Authorize checks if a user has the necessary permissions to perform an action on an object.
	Authorize(ctx context.Context, sub, obj, act string) (*v1.AuthorizeResponse, error)

	// AuthExpansion defines additional methods for extended auth operations, if needed.
	AuthExpansion
}

// AuthExpansion defines additional methods for auth operations.
type AuthExpansion interface{}

// authBiz is the implementation of the AuthBiz.
type authBiz struct {
	store store.IStore
	authn authn.Authenticator
	auth  auth.AuthProvider
}

// Ensure that *authBiz implements the AuthBiz.
var _ AuthBiz = (*authBiz)(nil)

// New creates and returns a new instance of *authBiz.
func New(store store.IStore, authn authn.Authenticator, auth auth.AuthProvider) *authBiz {
	return &authBiz{store: store, authn: authn, auth: auth}
}

// Login authenticates a user and returns a token.
func (b *authBiz) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	// Retrieve user information from the data storage by username.
	userM, err := b.store.User().Get(ctx, where.F("username", rq.Username))
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to retrieve user by username")
		return nil, i18n.FromContext(ctx).E(locales.RecordNotFound)
	}

	// Compare the obtained user information and the input password.
	// Because the password `userM.Password` stored in the database is an
	// encrypted password and cannot be decrypted, the comparison here
	// actually involves encrypting the `rq.Password` string using the
	// same method, and then comparing the encrypted string with the
	// one stored in the database. If they match, the password is verified.
	if err := authn.Compare(userM.Password, rq.Password); err != nil {
		log.W(ctx).Errorw(err, "Password does not match")
		return nil, i18n.FromContext(ctx).E(locales.IncorrectPassword)
	}

	// If the comparison passes, it means the password is correct.
	// Call `b.authn.Sign` to generate a refresh token.
	refreshToken, err := b.authn.Sign(ctx, userM.UserID)
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to generate refresh token")
		return nil, i18n.FromContext(ctx).E(locales.JWTTokenSignFail)
	}

	// Generate an access token for resource access.
	accessToken, err := b.auth.Sign(ctx, userM.UserID)
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to generate access token")
		return nil, i18n.FromContext(ctx).E(locales.JWTTokenSignFail)
	}

	// Return
	return &v1.LoginReply{
		RefreshToken: refreshToken.GetToken(),
		AccessToken:  accessToken.GetToken(),
		Type:         accessToken.GetTokenType(),
		ExpiresAt:    accessToken.GetExpiresAt(),
	}, nil
}

// Logout invalidates a token.
func (b *authBiz) Logout(ctx context.Context, rq *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	if err := b.authn.Destroy(ctx, contextx.AccessToken(ctx)); err != nil {
		log.W(ctx).Errorw(err, "Failed to remove token from cache")
		return nil, err
	}
	return &v1.LogoutResponse{}, nil
}

// RefreshToken refreshes an existing token and returns a new one.
func (b *authBiz) RefreshToken(ctx context.Context, rq *v1.RefreshTokenRequest) (*v1.LoginReply, error) {
	// Because a new token is issued, the old token needs to be destroyed.
	_ = b.authn.Destroy(ctx, contextx.AccessToken(ctx))

	userID := contextx.UserID(ctx)
	refreshToken, err := b.authn.Sign(ctx, userID)
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to generate refresh token")
		return nil, i18n.FromContext(ctx).E(locales.JWTTokenSignFail)
	}

	accessToken, err := b.auth.Sign(ctx, userID)
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to generate access token")
		return nil, i18n.FromContext(ctx).E(locales.JWTTokenSignFail)
	}

	return &v1.LoginReply{
		RefreshToken: refreshToken.GetToken(),
		AccessToken:  accessToken.GetToken(),
		Type:         accessToken.GetTokenType(),
		ExpiresAt:    accessToken.GetExpiresAt(),
	}, nil
}

// Authenticate validates an access token and returns the associated user ID.
func (b *authBiz) Authenticate(ctx context.Context, accessToken string) (*v1.AuthenticateResponse, error) {
	userID, err := b.auth.Verify(accessToken)
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to verify access token")
		return nil, err
	}

	return &v1.AuthenticateResponse{UserID: userID}, nil
}

// Authorize checks if a user has the necessary permissions to perform an action on an object.
func (b *authBiz) Authorize(ctx context.Context, sub, obj, act string) (*v1.AuthorizeResponse, error) {
	allowed, err := b.auth.Authorize(sub, obj, act)
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to authorize")
		return nil, err
	}
	return &v1.AuthorizeResponse{Allowed: allowed}, nil
}
