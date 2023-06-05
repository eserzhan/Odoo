package repository

type Bucket string 


const (
	RequestToken Bucket = "RequestToken"
	AccessToken Bucket = "AccessToken"
)

type TokenRepository interface {
	Save(chatId int64, token string, bucket Bucket) error
	Get(chatId int64, bucket Bucket) (string, error)
}