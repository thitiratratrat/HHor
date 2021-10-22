package dto

type DormDTO struct {
	ID            string `json:"id"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
	StartingPrice int    `json:"starting_price"`
	Zone          string `json:"zone"`
}
