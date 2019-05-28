package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	db := Connect()

	r := gin.Default()
	studentGroup := r.Group("students")
	{
		studentGroup.GET(":id", func(c *gin.Context) {
			studentID := c.Param("id")
			if studentID == "" {
				errRsp(c, 400, "MISSING_STUDENT_ID", "student id is mandatory")
				return
			}
			uStudentID, err := strconv.ParseUint(studentID, 10, 32)
			if err != nil {
				errRsp(c, 400, "MISSING_STUDENT_ID", "student id is mandatory")
				return
			}
			st := &Student{}
			if err := db.First(st, uStudentID).Error; err != nil {
				errRsp(c, 500, "", "")
				return
			}
			c.JSON(200, fromStudent(st))
		})

		studentGroup.POST("", func(c *gin.Context) {
			st := &StudentDto{}
			if err := c.ShouldBind(st); err != nil {
				errRsp(c, 400, "INVALID_PARAM", err.Error())
				return
			}

			stDB := st.ToDB()
			if err := db.Create(stDB).Error; err != nil {
				errRsp(c, 500, "", "")
				return
			}

			c.JSON(200, fromStudent(stDB))
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func errRsp(c *gin.Context, httpCode int, code, desc string) {
	if httpCode == 500 {
		c.JSON(httpCode, gin.H{
			"code": "INTERNAL_SERVER_ERROR",
			"desc": "internal server error",
		})
		return
	}
	c.JSON(httpCode, gin.H{
		"code": code,
		"desc": desc,
	})
}

type StudentDto struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

func fromStudent(st *Student) *StudentDto {
	return &StudentDto{
		ID:    int64(st.ID),
		Name:  st.Name,
		Class: st.Class,
	}
}

func (st *StudentDto) ToDB() *Student {
	return &Student{
		Name:  st.Name,
		Class: st.Class,
	}
}
