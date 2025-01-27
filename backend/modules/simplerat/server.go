package simplerat

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"strconv"
	"sync"
	"time"

	"github.com/blue-monads/turnix/backend/modules/simplerat/ratws"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/utils/libx/httpx"
	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

type ECPServer struct {
	app    xtypes.App
	db     *database.DB
	signer *signer.Signer

	websocket map[int64]*ratws.ECPWebsocket
	wLock     sync.RWMutex
}

func (e *ECPServer) register(group *gin.RouterGroup) error {

	mw := e.app.AsApiAction

	device := group.Group("/:pid/device")

	device.POST("/new", mw("addDevice", e.apiAddDevice))
	device.DELETE("/remove/:id", mw("removeDevice", e.apiRemoveDevice))
	device.GET("/list", mw("listDevice", e.apiListDevice))

	device.POST("/action/:did", mw("performDeviceAction", e.performDeviceAction))
	device.GET("/room/new/:did", (e.browserServiceWS))
	device.GET("/room/join/:rid", e.deviceServiceWS)

	device.POST("/finish-setup", e.apiFinishDeviceSetup)
	device.POST("/refresh", e.apiRefreshDevice)
	device.GET("/device-ws", e.deviceWS)

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

	return e.signer.SignProjectAdvisiery(ctx.ProjectId(), &DeviceClaim{
		Type:     DeviceClaimTypeSetup,
		DeviceId: did,
		Expires:  time.Now().Add(time.Hour * 24).Unix(),
	})
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
		pp.Println("@1", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		panic(err)
	}

	dbytes, err := e.signer.ParseProjectAdvisiery(pid, data.Token)
	if err != nil {
		pp.Println("@2", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var claim DeviceClaim
	err = json.Unmarshal(dbytes, &claim)
	if err != nil {
		pp.Println("@3", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	deviceId := claim.DeviceId

	t := time.Now()

	e.deviceTable(pid).
		Find(db.Cond{"id": deviceId}).
		Update(map[string]any{
			"status":      "active",
			"last_ip":     ctx.ClientIP(),
			"last_active": &t,
		})

	rclaim := &DeviceClaim{
		Type:     DeviceClaimTypeRefresh,
		DeviceId: deviceId,
		Expires:  time.Now().Add(time.Hour * 24 * 365).Unix(),
	}

	refreshToken, err := e.signer.SignProjectAdvisiery(pid, rclaim)
	if err != nil {
		pp.Println("@4", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sclaim := &DeviceClaim{
		Type:     DeviceClaimTypeSession,
		DeviceId: deviceId,
		Expires:  time.Now().Add(time.Hour * 24).Unix(),
	}

	pp.Println("@5", sclaim)

	sessionToken, err := e.signer.SignProjectAdvisiery(pid, sclaim)
	if err != nil {
		pp.Println("@6", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"refreshToken": refreshToken,
		"sessionToken": sessionToken,
	})
}

func (e *ECPServer) apiRefreshDevice(ctx *gin.Context) {}

func (e *ECPServer) readDeviceClaim(pid int64, ctx *gin.Context) (*DeviceClaim, error) {

	token := ctx.Query("token")
	claimBytes, err := e.signer.ParseProjectAdvisiery(pid, token)
	if err != nil {
		return nil, err
	}

	var claim DeviceClaim
	err = json.Unmarshal(claimBytes, &claim)
	if err != nil {
		return nil, err
	}

	return &claim, nil

}

func (e *ECPServer) deviceWS(ctx *gin.Context) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		pp.Println("@0", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	claim, err := e.readDeviceClaim(pid, ctx)
	if err != nil {
		pp.Println("@1", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if claim.Type != DeviceClaimTypeSession {
		pp.Println("@3", "invalid claim type")
		ctx.JSON(400, gin.H{"error": "invalid claim type"})
		return
	}

	pp.Println("@4", claim)

	deviceId := claim.DeviceId

	ws := e.getProjectWS(pid, true)
	if ws == nil {
		pp.Println("@5", "no websocket")
		ctx.JSON(400, gin.H{"error": "no websocket"})
		return
	}

	slog.Info("handling ws connection", "pid", pid, "deviceId", deviceId)

	ws.HandleAgentWS(deviceId, ctx)
}

func (e *ECPServer) getProjectWS(pid int64, create bool) *ratws.ECPWebsocket {
	e.wLock.RLock()
	ws := e.websocket[pid]
	e.wLock.RUnlock()

	if ws != nil {
		return ws
	}

	if !create {
		return nil
	}

	e.wLock.Lock()
	defer e.wLock.Unlock()
	ws = ratws.NewECPWebsocket(pid)

	e.websocket[pid] = ws

	go ws.Run()

	return ws

}

func (e *ECPServer) performDeviceAction(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	ws := e.getProjectWS(pid, false)

	if ws == nil {
		slog.Error("no websocket for project")
		return nil, errors.New("no websocket for project")
	}

	did := ctx.ParamInt64("did")

	out, err := io.ReadAll(ctx.Http.Request.Body)
	if err != nil {
		return nil, err
	}

	ctx.Http.Set("raw_mode", true)

	rbytes := ws.SendAgentMessage(ctx.Http.Request.Context(), did, out, true)

	if rbytes != nil {
		pp.Println("@rbytes", string(rbytes))
		ctx.Http.Request.Header.Set("Content-Type", "application/json")
		ctx.Http.Writer.WriteHeader(200)
		ctx.Http.Writer.Write(rbytes)

	} else {
		ctx.Http.JSON(400, gin.H{"error": "no response"})
	}

	return nil, nil
}

type ServiceRoomMessage struct {
	MType string         `json:"mtype"`
	Data  map[string]any `json:"data"`
}

func (e *ECPServer) browserServiceWS(ctx *gin.Context) {
	tok := ctx.Query("token")
	claim, err := e.signer.ParseAccess(tok)
	if err != nil {
		httpx.WriteAuthErr(ctx, err)
		return
	}

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		panic(err)
	}

	did, err := strconv.ParseInt(ctx.Param("did"), 10, 64)
	if err != nil {
		panic(err)
	}

	ws := e.getProjectWS(pid, true)
	if ws == nil {
		ctx.JSON(400, gin.H{"error": "no websocket"})
		return
	}

	wsroom := ws.AddRoom(claim.UserId, did)

	ok := wsroom.AddBrowserToRoom(claim.UserId, ctx)
	if !ok {
		ctx.AbortWithStatus(400)
		return
	}

	mtype := ctx.Query("mtype")
	if mtype == "" {
		mtype = "service.hello"
	}

	data := ServiceRoomMessage{
		MType: mtype,
		Data:  map[string]any{},
	}

	queryValues := ctx.Request.URL.Query()
	for k, v := range queryValues {
		if k == "mtype" {
			continue
		}

		if k == "token" {
			continue
		}

		data.Data[k] = v
	}

	data.Data["room_id"] = wsroom.GetRoomId()

	jdata, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	ws.SendAgentMessage(ctx.Request.Context(), did, jdata, false)

}

func (e *ECPServer) deviceServiceWS(ctx *gin.Context) {

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		panic(err)
	}

	claim, err := e.readDeviceClaim(pid, ctx)
	if err != nil {
		pp.Println("@1", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if claim.Type != DeviceClaimTypeSession {
		pp.Println("@3", "invalid claim type")
		ctx.JSON(400, gin.H{"error": "invalid claim type"})
		return
	}

	rid, err := strconv.ParseInt(ctx.Param("rid"), 10, 64)
	if err != nil {
		pp.Println("@4", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	ws := e.getProjectWS(pid, false)

	if ws == nil {
		pp.Println("@5", "no websocket")
		ctx.JSON(400, gin.H{"error": "no websocket"})
		return
	}

	wsroom := ws.GetRoom(rid)
	if wsroom == nil {
		pp.Println("@6", "no room")
		ctx.JSON(400, gin.H{"error": "no room"})
		return
	}

	ok := wsroom.AddAgentToRoom(claim.DeviceId, ctx)
	if !ok {
		pp.Println("@7", "add agent failed")
		ctx.JSON(400, gin.H{"error": "add agent failed"})
		return
	}

}
