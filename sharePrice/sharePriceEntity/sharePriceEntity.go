package sharePriceEntity

import "time"

type (
	SharePrice struct {
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		UpdatedAt time.Time `json:"updated_at"`
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
