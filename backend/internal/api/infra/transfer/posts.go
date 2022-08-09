package transfer

import (
	"github.com/norun9/postmantest/internal/api/domain/model"
	"github.com/norun9/postmantest/pkg/dbmodels"
	"github.com/sharedine/next/backend/internal/api/domain/model"
	"github.com/sharedine/next/backend/pkg/dbmodels"
)

// ToAddressConfirmationImageEntity :
func ToPostDomain(m *dbmodels.Post) *model.Post {
	return (*model.Post)(m)
}
