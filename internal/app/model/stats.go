package model

type Stats struct {
	TotalMembers int64  `json:"totalMembers"` // 社区人数
	TotalTask    uint64 `json:"totalTask"`    // 需求数量
	TotalNFT     uint64 `json:"totalNFT"`     // NFT数量
}
