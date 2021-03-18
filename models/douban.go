package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zhangyusheng/data-center/pkg/logging"
)

type Movie struct {
	Model
	Title    string  `json:"title"`
	Subtitle string  `json:"subtitle"`
	Other    string  `json:"other"`
	Desc     string  `json:"desc"`
	Year     string  `json:"desc"`
	Area     string	 `json:"desc"`
	Tag      string  `json:"desc"`
	Star     string  `json:"desc"`
	Comment  string  `json:"desc"`
	Quote    string  `json:"desc"`
}

func AddMovie(movie Movie) error {
	if err := db.Create(&movie).Error; err != nil {
		return err
	}

	return nil
}

func UpsertMovie(movieIns Movie) error {
	var movie Movie
	err := db.Select("id").Where("title = ?", movieIns.Title).First(&movie).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logging.Logger.Info(err.Error())
			results := db.Create(&movieIns).Error
			if results != nil {
				return results
			}
		}
	} else {
		if err := db.Model(&Movie{}).Where("title = ?", movieIns.Title).Updates(movieIns).Error; err != nil {
			return err
		}
	}
	return nil
}