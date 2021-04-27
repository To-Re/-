package method

import (
	"bishe/backend/util"
	"fmt"
)

func Login() {
	ret, _ := util.CreteToken(1, 1)
	fmt.Println("token : " + ret)
	tmp, _ := util.ParseToken(ret)

	fmt.Println(tmp)
}
