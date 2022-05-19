package model

// type User struct {
// 	ID       uint64 `gorm:"primaryKey"`
// 	Username string
// 	Join     time.Time
// 	About    string
// 	UserStats
// 	UserContacts
// }

// type UserStats struct {
// 	Group     string
// 	Status    string
// 	Age       string
// 	Birthday  string
// 	Sex       string
// 	City      string
// 	Interests string
// }

// type UserContacts struct {
// 	Site    string
// 	ICQ     string
// 	Jabber  string
// 	Byond   string
// 	Discord string
// 	Skype   string
// }

// type Status struct {
// 	ID        uint64    `gorm:"primaryKey;not null"`
// 	UserID    uint64    `gorm:"not null"`
// 	Text      string    `gorm:"not null"`
// 	CreatedAt time.Time `gorm:"not null"`
// }

// type StatusReply struct {
// 	ID        uint64    `gorm:"primaryKey;not null"`
// 	StatusID  uint64    `gorm:"not null"`
// 	UserID    uint64    `gorm:"not null"`
// 	Text      string    `gorm:"not null"`
// 	CreatedAt time.Time `gorm:"not null"`
// }

// type RepliesHTML struct {
// 	HTML          string `json:"html"`
// 	Status        string `json:"status"`
// 	StatusReplies int    `json:"status_replies"`
// }

// type Forum struct {
// 	ID          uint64 `gorm:"primaryKey;not null"`
// 	Name        string `gorm:"not null"`
// 	Description string
// }

// type Topic struct {
// 	ID        uint64    `gorm:"primaryKey;not null"`
// 	UserID    uint64    `gorm:"not null"`
// 	ForumID   uint64    `gorm:"not null"`
// 	Title     string    `gorm:"not null"`
// 	CreatedAt time.Time `gorm:"not null"`
// }

// type Post struct {
// 	ID        uint64    `gorm:"primaryKey;not null"`
// 	TopicID   uint64    `gorm:"not null"`
// 	UserID    uint64    `gorm:"not null"`
// 	Text      string    `gorm:"not null"`
// 	CreatedAt time.Time `gorm:"not null"`
// }
