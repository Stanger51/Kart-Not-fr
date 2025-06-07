package assetutils

import (
	"strconv"

	"github.com/kartFr/Asset-Reuploader/internal/roblox/assets"
)

func NewPermissionBodyFromIds(universeIDs []int64) assets.PermissionRequest {
	requests := make([]assets.PermissionRequestItem, len(universeIDs))

	for i, universeID := range universeIDs {
		requests[i] = assets.PermissionRequestItem{
			SubjectType: "Universe",
			SubjectID:   strconv.FormatInt(universeID, 10),
			Action:      "Use",
		}
	}

	return assets.PermissionRequest{
		Requests: requests,
	}
}
