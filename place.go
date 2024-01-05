package dagestan

type Place struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	TypeId      int     `json:"typeId"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Type struct {
	Id       int    `json:"id"`
	NameType string `json:"nameType"`
}

type UserPlace struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	PlaceId int `json:"placeId"`
}

type CoverPlace struct {
	Id      int `json:"Id"`
	PlaceId int `json:"PlaceId"`
	CoverId int `json:"coverId"`
}

type Cover struct {
	Id       int    `json:"Id"`
	CoverUrl string `json:"coverUrl"`
}
