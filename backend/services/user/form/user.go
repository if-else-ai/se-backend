package form

type User struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Email       string `bson:"email"`
	Password    string `bson:"password"`
	TelNo       string `bson:"telNo"`
	Address     string `bson:"address"`
	DateOfBirth string `bson:"dateOfBirth"`
	Gender      string `bson:"gender"`
}
