package method

import (
	"bishe/backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	ret, _ := util.CreteToken(1, 1)
	fmt.Println("token : " + ret)
	tmp, _ := util.ParseToken(ret)

	fmt.Println(tmp)
}
