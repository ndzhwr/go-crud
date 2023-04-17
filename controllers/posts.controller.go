package controllers

import (
	"go-crud/models"
	"go-crud/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func MainRouteController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

type Body struct {
	Body  string
	Title string
}

func PostCreate(c *gin.Context) {
	var body Body

	c.Bind(&body)

	if strings.TrimSpace(body.Title) == "" {
		c.JSON(406, gin.H{
			"success": false,
			"message": "post lacks the title parameter",
		})
		return
	} else if strings.TrimSpace(body.Body) == "" {
		c.JSON(406, gin.H{
			"success": false,
			"message": "post lacks the body parameter",
		})
		return
	}
	// creating and saving the post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := utils.DB.Create(&post)

	if result.Error != nil {
		log.Fatal("create error", result.Error)
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	} else {

		// sending the post

		c.JSON(200, gin.H{
			"message": "Post created successfully",
			"success": true,
			"post":    post,
		})

	}
}

func AllPosts(c *gin.Context) {
	var posts []models.Post
	result := utils.DB.Find(&posts)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	} else {

		// sending the post

		c.JSON(200, gin.H{
			"success": true,
			"posts":   posts,
		})

	}
}

func SinglePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	utils.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(404, gin.H{
			"message": "post not found",
			"success": false,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"post":    post,
		})

	}

}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		c.JSON(406, gin.H{
			"message": "postid is missing in the url params",
			"success": false,
		})
		return

	}

	var body Body
	c.Bind(&body)

	var post models.Post
	result := utils.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(404, gin.H{
			"message": "post not found",
			"success": false,
		})
		return
	} else if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	}

	res := utils.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "post updated successfully",
		"success": true,
		"post":    post,
	})

}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		c.JSON(406, gin.H{
			"message": "postid is missing in the url params",
			"success": false,
		})
		return
	}

	result := utils.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "post deleted successfully",
		"success": true,
	})

}
