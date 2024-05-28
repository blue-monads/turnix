import axios from "axios"
import { browser } from '$app/environment'
import type { AxiosInstance } from "axios"



export interface ProjectDef {
    name: string
    ptype: string
    slug: string
    info: string
    icon: string
    is_external: boolean
    event_types: string[]
}


export interface ProjectHook {
    id: number;
    event: string;
    order_id: number;
    runas_user_id: number;
    hook_type: string;
    hook_code: string;
    envs: string;
    project_id: number;
    extrameta?: string;
}



export class RootAPI {
    client: AxiosInstance
    projectClient: AxiosInstance

    constructor() {
        let access_token;

        if (browser) {
            access_token = localStorage.getItem("access_token")
            if (!access_token) {
                window.location.pathname = "/z/pages/auth/login"
            }
        }

        this.client = axios.create({
            baseURL: `${location.origin}/z/api/v1/`,
            headers: {
                "Authorization": access_token
            }
        })

        this.projectClient = axios.create({
            baseURL: `${location.origin}/z/project/`,
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


    getProjectTypeForm = (ptype: string) => {
        return this.client.get(`/project_types/${ptype}/form`)
    }

    getProjectType = (ptype: string) => {
        return this.client.get<ProjectDef>(`/project_types/${ptype}`)
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


    listProjectHooks = (pid: string) => {
        return this.client.get<ProjectHook[]>(`/project/${pid}/hook`)
    }

    addProjectHook = (pid: string, data: object) => {
        return this.client.post<Partial<ProjectHook>>(`/project/${pid}/hook`, data)
    }


    getProjectHook = (pid: string, id: string) => {
        return this.client.get<ProjectHook>(`/project/${pid}/hook/${id}`)
    }


    updateProjectHook = (pid: string, id: string, data: object) => {
        return this.client.post<Partial<ProjectHook>>(`/project/${pid}/hook/${id}`, data)
    }

    removeProjectHook = (pid: string, id: string) => {
        return this.client.delete(`/project/${pid}/hook/${id}`)
    }


}



