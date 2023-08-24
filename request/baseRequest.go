package request

import (
	"go-web-example/errorcode"
	"strings"

	"github.com/gin-gonic/gin"
)

// PageOffset Paging parameters
type PageOffset struct {
	Page   int `json:"page" form:"page" validate:"required" comment:"页码"`
	Size   int `json:"size" form:"size" validate:"required" comment:"条数"`
	Offset int
}

// CheckInputParams  check page size
func (po *PageOffset) CheckInputParams(c *gin.Context) error {
	if err := c.ShouldBind(po); err != nil {
		return err
	}

	//validate request params
	if errs, err := Validate(po); err != nil {
		return errorcode.New(errorcode.ErrParams, strings.Join(errs, ","), nil)
	}

	po.Offset = (po.Page - 1) * po.Size
	return nil
}
