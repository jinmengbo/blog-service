package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (T Article) Get(c *gin.Context)  {}
func (T Article) List(c *gin.Context)  {}
func (T Article) Create(c *gin.Context)  {}
func (T Article) Update(c *gin.Context)  {}
func (T Article) Delete(c *gin.Context)  {}
