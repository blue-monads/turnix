package simplerat

import (
	"encoding/binary"
	"encoding/json"
	"strconv"
	"time"

	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/gin-gonic/gin"
	"github.com/upper/db/v4"
)

type ECPServer struct {
	app    xtypes.App
	db     *database.DB
	signer *signer.Signer
}

func (e *ECPServer) register(group *gin.RouterGroup) error {

	mw := e.app.AsApiAction

	device := group.Group("/:pid/device")

	device.POST("/new", mw("addDevice", e.apiAddDevice))
	device.DELETE("/remove/:id", mw("removeDevice", e.apiRemoveDevice))
	device.GET("/list", mw("listDevice", e.apiListDevice))
	device.POST("/finish-setup", e.apiFinishDeviceSetup)
	device.POST("/refresh", e.apiRefreshDevice)

	return nil
}

type NewDevice struct {
	Name            string `json:"name"`
	Dtype           string `json:"dtype"`
	IsBeaconEnabled int    `json:"beacon_enabled"`
}

func (e *ECPServer) apiAddDevice(ctx xtypes.ContextPlus) (any, error) {

	var data NewDevice
	err := ctx.Http.Bind(&data)
	if err != nil {
		return nil, err
	}

	did, err := e.dbAddDevice(ctx.ProjectId(), ctx.Claim.UserId, &Device{
		Name:            data.Name,
		RegisterToken:   "",
		Dtype:           data.Dtype,
		Status:          "inactive",
		IsBeaconEnabled: data.IsBeaconEnabled,
	})
	if err != nil {
		return nil, err
	}

	deviceIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(deviceIdBytes, uint64(did))

	return e.signer.SignProjectAdvisiery(ctx.ProjectId(), deviceIdBytes)
}

func (e *ECPServer) apiRemoveDevice(ctx xtypes.ContextPlus) (any, error) {

	err := e.dbDeleteDevice(ctx.ProjectId(), ctx.Claim.UserId, ctx.Claim.UserId)

	return nil, err
}

func (e *ECPServer) apiListDevice(ctx xtypes.ContextPlus) (any, error) {

	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)

	return e.dbListDevice(ctx.ProjectId(), ctx.Claim.UserId, offset)
}

type FinishDeviceSetup struct {
	Token string `json:"token"`
}

type DeviceClaim struct {
	Type     uint8 `json:"type"`
	DeviceId int64 `json:"deviceId"`
	Expires  int64 `json:"expires"`
}

const (
	DeviceClaimTypeSetup   = 1
	DeviceClaimTypeRefresh = 2
	DeviceClaimTypeSession = 3
)

func (e *ECPServer) apiFinishDeviceSetup(ctx *gin.Context) {

	var data FinishDeviceSetup
	err := ctx.Bind(&data)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		panic(err)
	}

	dbytes, err := e.signer.ParseProjectAdvisiery(pid, data.Token)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	deviceId := int64(binary.BigEndian.Uint64(dbytes))

	e.deviceTable(pid).
		Find(db.Cond{"id": deviceId}).
		Update(map[string]any{
			"status": "active",
			"lastIp": ctx.ClientIP(),
			"lastAt": db.Raw("NOW()"),
		})

	rclaim := &DeviceClaim{
		Type:     DeviceClaimTypeRefresh,
		DeviceId: deviceId,
		Expires:  time.Now().Add(time.Hour * 24 * 365).Unix(),
	}

	rclaimBytes, err := json.Marshal(rclaim)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := e.signer.SignProjectAdvisiery(pid, rclaimBytes)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sclaim := &DeviceClaim{
		Type:     DeviceClaimTypeSession,
		DeviceId: deviceId,
		Expires:  time.Now().Add(time.Hour).Unix(),
	}

	sclaimBytes, err := json.Marshal(sclaim)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sessionToken, err := e.signer.SignProjectAdvisiery(pid, sclaimBytes)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"refreshToken": refreshToken, "sessionToken": sessionToken})
}

func (e *ECPServer) apiRefreshDevice(ctx *gin.Context) {

}
