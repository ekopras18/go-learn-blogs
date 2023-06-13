package models

import (
	"errors"
	"go-learn-blogs/config"
	"go-learn-blogs/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func Get() ([]*entities.Blogs, error) {

	var blogs []*entities.Blogs
	err := config.DB.Find(&blogs).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return blogs, nil

}

func Store(blog *entities.Blogs) error {

	err := validate.Struct(blog)
	if err != nil {
		return err
	}

	err = config.DB.Create(&blog).Error
	if err != nil {
		return err
	}

	return nil
}

func Show(id string) (*entities.Blogs, error) {

	var blog entities.Blogs
	err := config.DB.First(&blog, id).Error
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func Update(id string, blog *entities.Blogs) error {
	existingBlog := &entities.Blogs{}
	err := config.DB.First(existingBlog, id).Error
	if err != nil {
		return err
	}

	existingBlog.Title = blog.Title
	existingBlog.Author = blog.Author
	existingBlog.Tags = blog.Tags
	existingBlog.Content = blog.Content

	err = config.DB.Save(existingBlog).Error
	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {

	var blog entities.Blogs
	err := config.DB.First(&blog, id).Error
	if err != nil {
		return err
	}

	err = config.DB.Delete(&blog).Error
	if err != nil {
		return err
	}

	return nil
}
