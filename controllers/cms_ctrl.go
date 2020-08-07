package controllers

import (
	"github.com/astaxie/beego/pkg"
)

// Cms_ctrlController operations for Cms_ctrl
type Cms_ctrlController struct {
	beego.Controller
}

// URLMapping ...
func (c *Cms_ctrlController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetTwo", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Cms_ctrl
// @Param	body		body 	models.Cms_ctrl	true		"body for Cms_ctrl content"
// @Success 201 {object} models.Cms_ctrl
// @Failure 403 body is empty
// @router / [post]
func (c *Cms_ctrlController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Cms_ctrl by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Cms_ctrl
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Cms_ctrlController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Cms_ctrl
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Cms_ctrl
// @Failure 403
// @router / [get]
func (c *Cms_ctrlController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Cms_ctrl
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Cms_ctrl	true		"body for Cms_ctrl content"
// @Success 200 {object} models.Cms_ctrl
// @Failure 403 :id is not int
// @router /:id [put]

func (c *Cms_ctrlController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Cms_ctrl
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Cms_ctrlController) Delete() {

}