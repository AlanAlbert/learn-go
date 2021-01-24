package models

type Tag struct {
	Model
	
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, where interface{}) []Tag {
	var tags []Tag
	db.Where(where).Offset(pageNum).Limit(pageSize).Find(&tags)

	return tags
}

func GetTagTotal(where interface{}) int {
	var total int
	db.Model(&Tag{}).Where(where).Count(&total)

	return total
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
		ModifiedBy: createdBy,
	})
}