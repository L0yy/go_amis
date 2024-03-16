package handles

import "github.com/gofiber/fiber/v2"

type RespendJson struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func HandleGetConfig(c *fiber.Ctx) error {
	resp := RespendJson{
		Status: 0,   大厦西侧
		Msg:    "susc ok",
		Data: []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}{
			{Id: 1, Name: "Alice"},
			{Id: 2, Name: "Bob"},
			{Id: 3, Name: "Bobcc"},
		},
	}
	return c.JSON(resp)
}

func HandleRootPathIndex(c *fiber.Ctx) error {
	return c.SendFile("D:\\Code\\GO\\amis_fiber\\static\\index.site")
}
func HandleRootPath(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"pageTitle":     "index",
		"pageSchemaApi": "GET:/index",
	})
}
