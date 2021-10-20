package userDomains

type User struct {
	ID                string `json:"id" db:"id"`
	Name              string `json:"name" db:"name"`
	Email             string `json:"email" db:"email"`
	ProfilePictureURL string `json:"profilePictureURL" db:"profile_picture_url"`
}
