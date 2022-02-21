package model

type User struct {
	UserName string `db:user_name`
	Password string `db:password`
	Emil string `db:emil`
	Introduction string `db:introduction`
	Phone int `db:phone`
	QQ int `db:qq`
	Gender string `db:gender`
	Birth string `db:birth`
}
type UserInfo struct {
	UserName string `db:user_name`
	Topic string `db:topic`
}
type Collection struct {
	Name string `db:name`
	TopicName string `db:topic_name`
	dislikeName string `db:dislike_name`
}