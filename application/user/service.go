package user

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/infrastructure"
	"github.com/chaihaobo/boice-blog-api/infrastructure/store/repository/user"
	userdto "github.com/chaihaobo/boice-blog-api/model/dto/user"
	"github.com/chaihaobo/boice-blog-api/resource"
	"github.com/chaihaobo/boice-blog-api/utils/crypto"
	"github.com/chaihaobo/boice-blog-api/utils/jwt"
)

type (
	Service interface {
		Login(ctx context.Context, request *userdto.LoginRequest) (*userdto.LoginResponse, error)
		TokenManger() jwt.TokenManager
	}

	service struct {
		res          resource.Resource
		infra        infrastructure.Infrastructure
		passwordHash crypto.Hash
		tokenManger  jwt.TokenManager
	}
)

func (s *service) TokenManger() jwt.TokenManager {
	return s.tokenManger
}

func (s *service) userRepository() user.Repository {
	return s.infra.Store().Repository().User()
}

func NewService(res resource.Resource, infra infrastructure.Infrastructure) Service {
	jwtConfig := res.Configuration().JWT
	return &service{
		res:   res,
		infra: infra,
		passwordHash: crypto.NewArgon2IDHash(&crypto.GeneratePwdParams{
			Memory:      defaultMemory,
			Iterations:  defaultIterations,
			Parallelism: defaultParallelism,
			SaltLength:  defaultSaltLength,
			KeyLength:   defaultKeyLength,
		}),

		tokenManger: jwt.NewJWTManager(
			jwtConfig.AccessTokenSecretKey,
			jwtConfig.RefreshTokenSecretKey,
			jwtConfig.AccessTokenDuration,
			jwtConfig.RefreshTokenDuration,
		),
	}
}
