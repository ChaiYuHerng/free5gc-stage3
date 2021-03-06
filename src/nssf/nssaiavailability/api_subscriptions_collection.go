/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssaiavailability

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"free5gc/lib/http_wrapper"
	. "free5gc/lib/openapi/models"
	"free5gc/src/nssf/handler"
	"free5gc/src/nssf/handler/message"
	"free5gc/src/nssf/util"
)

func ApiSubscriptionsCollection(c *gin.Context) {
	var request NssfEventSubscriptionCreateData
	err := c.ShouldBindJSON(&request)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		d := ProblemDetails{
			Title:  util.MALFORMED_REQUEST,
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		c.JSON(http.StatusBadRequest, d)
		return
	}
	req := http_wrapper.NewRequest(c.Request, request)

	msg := message.NewMessage(message.NSSAIAvailabilityPost, req)

	handler.SendMessage(msg)
	rsp := <-msg.ResponseChan

	httpResponse := rsp.HttpResponse
	c.JSON(httpResponse.Status, httpResponse.Body)
}
