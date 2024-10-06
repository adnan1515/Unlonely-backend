package models


type Remm struct {
	Pagination  struct {
		Size int `json:"last_visible_page"`
		HasNext bool `json:"has_next_page"`
	} `json:"pagination"`
	Data[] struct {
		MalId string `json:"mal_id"`
		Entry[] struct {
			MalId int `json:"mal_id"`
			Url string `json:"url"`
			Images struct {
				Jpg struct {
					Image string `json:"image_url"`
					SmallImage string `json:"small_image_url"`
					LargeImage string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					Image string `json:"image_url"`
					SmallImage string `json:"small_image_url"`
					LargeImage string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			
		} `json:"entry"`
	} `json:"data"`	


}