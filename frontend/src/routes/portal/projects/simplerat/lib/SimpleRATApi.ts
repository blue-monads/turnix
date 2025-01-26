import type { AxiosInstance } from "axios"


export interface Device  {
    id: number
    name: string
    register_token: string
    dtype: string
    status: string
    last_ip: string
    beacon_enabled: number
    last_active: string
    created_at: string
}

export interface DeviceBeacon  {
    id: number
    title: string
    device_id: number
    lat: number
    lng: number
    beacon_type: string
    beacon_data: string
    created_at: string
}

export const NewSimpleRATApi = (rootApi: {
    projectClient: AxiosInstance
    access_token: string  | null
}) => {
    return new SimpleRATApi(rootApi.projectClient, rootApi.access_token!)
}



export class SimpleRATApi {
    client: AxiosInstance
    access_token: string 
    constructor(c: AxiosInstance, token: string) {
        this.client = c
        this.access_token = token
        
    }

    // /z/api/v1/project/:pid

    listDevices = (pid: string) => {
        return this.client.get<Device[]>(`simplerat/${pid}/device/list`)
    }

    addDevice = (pid: string, data: Partial<Device>) => {
        return this.client.post(`simplerat/${pid}/device/new`, data)
    }

    removeDevice = (pid: string, id: string) => {
        return this.client.delete(`simplerat/${pid}/device/remove/${id}`)
    }

    performDeviceAction = (pid: string, did: string, mtype: string, data: any) => {
        return this.client.post(`simplerat/${pid}/device/action/${did}`, {
            mtype,
            data
        })
    }

    getNewRoomUrl = (pid: string, did: number) => {
        return `${location.origin}/z/project/simplerat/${pid}/device/room/new/${did}?token=${this.access_token}`
    }

}
