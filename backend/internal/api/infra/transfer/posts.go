package transfer

import (
	"github.com/norun9/postmantest/internal/api/domain/model"
	"github.com/norun9/postmantest/pkg/dbmodels"
)

// ToPostDomain :
func ToPostDomain(m *dbmodels.Post) *model.Post {
	return (*model.Post)(m)
}
