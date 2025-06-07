package uploaderror

import (
	"fmt"

	"github.com/kartFr/Asset-Reuploader/internal/roblox/develop"
)

func New(idsProcessed, totalIDs int, msg string, assetInfo *develop.AssetInfo, err any) error {
	if msg != "" {
		msg += " "
	}
	return fmt.Errorf("[%d/%d] %s%s(%d): %v",
		idsProcessed,
		totalIDs,
		msg,
		assetInfo.Name,
		assetInfo.ID,
		err,
	)
}

func NewBatch(idsStart, idsEnd, totalIDs int, msg, err any) error {
	return fmt.Errorf("[%d-%d/%d] %s: %v", idsStart, idsEnd, totalIDs, msg, err)
}
