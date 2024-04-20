import axios from "axios"
import { browser } from '$app/environment'
import type { AxiosInstance } from "axios"
import { BASE_URL } from "./common"



export class RootAPI {
    client: AxiosInstance

    constructor() {
        let access_token;

        if (browser) {
            access_token = localStorage.getItem("access_token")
            if (!access_token) {
                window.location.pathname = "/z/pages/auth/login"
            }
        }

        this.client = axios.create({
            baseURL: BASE_URL + "api/v1/",
            headers: {
                "Authorization": access_token
            }
        })

        if (window !== undefined) {
            (window as any)["_turnis_api_"] = this
        }
    }


    // projects

    listProjectTypes = () => {
        return this.client.get(`/project_types`)
    }


    getProjectType = (ptype: string) => {
        return this.client.get(`/project_types/${ptype}`)
    }


    listProjects = (ptype?: string) => {
        return this.client.get(`/project?ptype=${ptype || ""}`)
    }

    addProject = (data: object) => {
        return this.client.post("/project", data)
    }

    updateProject = (pid: string, data: object) => {
        return this.client.post(`/project/${pid}`, data)
    }

    getProject = (pid: string) => {
        return this.client.get(`/project/${pid}`)
    }

    removeProject = (pid: string) => {
        return this.client.delete(`/project/${pid}`)
    }

    inviteUserToPoject = (pid: string) => {
        return this.client.post(`/project/${pid}/user`)
    }

    removeUserFromPoject = (pid: string) => {
        return this.client.delete(`/project/${pid}/user`)
    }

}



