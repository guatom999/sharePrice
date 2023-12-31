package sharePriceEntity

import "time"

type (
	InsertSharePrice struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	SharePrice struct {
		Id        uint32    `json:"id"`
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ShareTrackReq struct {
		Name string `json:"name" validate:"required,max=64"`
	}
)

// type (
// 	InsertCockroachDto struct {
// 		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
// 		Amount    uint32    `json:"amount"`
// 		CreatedAt time.Time `json:"createdAt"`
// 	}

// 	Cockroach struct {
// 		Id        uint32    `json:"id"`
// 		Amount    uint32    `json:"amount"`
// 		CreatedAt time.Time `json:"createdAt"`
// 	}

// 	CockroachPushNotificationDto struct {
// 		Title        string `json:"title"`
// 		Amount       uint32 `json:"amount"`
// 		ReportedTime string `json:"createdAt"`
// 	}
// )
