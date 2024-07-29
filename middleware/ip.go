package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPData struct {
	City       string  `json:"city"`
	Country    string  `json:"country"`
	Query      string  `json:"query"`
	RegionName string  `json:"regionName"`
	ISP        string  `json:"isp"`
	Zip        string  `json:"zip"`
	Lat        float32 `json:"lat"`
	Lon        float32 `json:"lon"`
	Timezone   string  `json:"timezone"`
}

func IPInfoMiddleware(c *gin.Context) {

	client := &http.Client{}
	url := "http://ip-api.com/json"

	ip := c.ClientIP()
	url = fmt.Sprintf("%s/%s", url, ip)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		c.Abort()
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve IP information"})
		c.Abort()
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var ipData IPData

	if err := json.Unmarshal(body, &ipData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse IP information"})
		c.Abort()
		return
	}
	c.Set("ipData", ipData)
	c.Next()
}
