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


export interface File {
    id: number;
    name: string;
    ftype: string;
    path: string;
    size: number;
    mime: string;
    hash: string;
    external: boolean;
    is_public: boolean;
    is_folder: boolean;
    owner_user_id: number;
    owner_proj_id: number;
    created_at: string;
}

export interface ProjectSQLExec {
    input: string;
    name: string;
    data: any[];
}

export interface User {
    id: number;
    name: string;
    utype: string;
    email: string;
    bio: string;
    phone: string;

    password: string;
    email_verified: boolean;
    owner_user_id: number;
    owner_project_id: number;
    disabled: boolean;
    msg_read_head: number;
    extrameta: any;
    created_at: string;
}

export interface ProjectPlugin {
    id: number;
    name: string;
    ptype: string;
    project_id: number;
    server_code: string;
    client_code: string;
    created_by: number;
    updated_by: number;
    created_at: string;
    updated_at: string;
}

export interface TableColumn {
    name: string;
    type: string;
    not_null: boolean;
    default_value: string;
    primary_key: boolean;
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
            (window as any)["_turnix_api_"] = this
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

    addProjectHook = (pid: string, data: Partial<ProjectHook>) => {
        return this.client.post(`/project/${pid}/hook`, data)
    }


    getProjectHook = (pid: string, id: string) => {
        return this.client.get<ProjectHook>(`/project/${pid}/hook/${id}`)
    }


    updateProjectHook = (pid: string, id: string, data: Partial<ProjectHook>) => {
        return this.client.post(`/project/${pid}/hook/${id}`, data)
    }

    removeProjectHook = (pid: string, id: string) => {
        return this.client.delete(`/project/${pid}/hook/${id}`)
    }

    runProjectSQL = (pid: string, data: ProjectSQLExec) => {
        return this.client.post(`/project/${pid}/sqlexec`, data)
    }

    runProjectSQL2 = (pid: string, data: { qstr: string, data: any[] }) => {
        return this.client.post(`/project/${pid}/sqlexec2`, data)
    }


    listProjectTables = (pid: string) => {
        return this.client.get<string[]>(`/project/${pid}/tables`)
    }


    listProjectTableColumns = (pid: string, table: string) => {
        return this.client.get<TableColumn[]>(`/project/${pid}/tables/${table}/columns`)
    }

    // 	apiv1.POST("/project/:pid/autoquery", a.accessMiddleware(a.autoQueryProjectTable))

    autoQueryProjectTable = (pid: string, table: string, data: any) => {
        return this.client.post(`/project/${pid}/autoquery?table=${table}`, data)
    }

    // project files

    listProjectFiles = (pid: string, path?: string) => {
        return this.client.get<File[]>(`/project/${pid}/files?path=${path}`)
    }

    addProjectFile = (pid: string, name: string, path: string, data: any) => {
        return this.client.post(`/project/${pid}/files?name=${name}&path=${path}`, data, {
            headers: {
                "Content-Type": "application/octet-stream"
            }
        })
    }

    addProjectFolder = (pid: string, name: string, path: string) => {
        return this.client.put(`/project/${pid}/files`, {
            name,
            path
        })
    }

    getProjectFile = (pid: string, id: string) => {
        return this.client.get(`/project/${pid}/files/${id}`, {
            responseType: "blob"
        })
    }

    deleteProjectFile = (pid: string, id: string) => {
        return this.client.delete(`/project/${pid}/files/${id}`)
    }

    // self users

    listSelfUsers = () => {
        return this.client.get<User[]>(`/self/users`)
    }

    addSelfUser = (data: Partial<User>) => {
        return this.client.post(`/self/users`, data)
    }

    getSelfUser = (uid: string) => {
        return this.client.get<User>(`/self/users/${uid}`)
    }

    updateSelfUser = (uid: string, data: Partial<User>) => {
        return this.client.post(`/self/users/${uid}`, data)
    }

    deleteSelfUser = (uid: string) => {
        return this.client.delete(`/self/users/${uid}`)
    }

    getSelfSelf = () => {
        return this.client.get<User>(`/self/self`)
    }



    // plugins

    listProjectPlugins = (pid: string) => {
        return this.client.get<ProjectPlugin[]>(`/project/${pid}/plugins`)
    }

    addProjectPlugin = (pid: string, data: Partial<ProjectPlugin>) => {
        return this.client.post(`/project/${pid}/plugins`, data)
    }

    removeProjectPlugin = (pid: string, id: string) => {
        return this.client.delete(`/project/${pid}/plugins/${id}`)
    }

    updateProjectPlugin = (pid: string, id: string, data: Partial<ProjectPlugin>) => {
        return this.client.post(`/project/${pid}/plugins/${id}`, data)
    }

    getProjectPlugin = (pid: string, id: string) => {
        return this.client.get<ProjectPlugin>(`/project/${pid}/plugins/${id}`)
    }



    // self files

    listSelfFiles = (path?: string) => {
        return this.client.get<File[]>(`/self/files?path=${path}`)
    }

    addSelfFile = (name: string, path: string, data: any) => {
        return this.client.post(`/self/files?name=${name}&path=${path}`, data, {
            headers: {
                "Content-Type": "application/octet-stream"
            }
        })
    }

    addSelfFolder = (path: string, name: string) => {
        return this.client.put(`/self/files`, {
            name,
            path
        })
    }

    getSelfFile = (id: string) => {
        return this.client.get(`/self/files/${id}`, {
            responseType: "blob"
        })
    }

    deleteSelfFile = (id: string) => {
        return this.client.delete(`/self/files/${id}`)
    }


    // messages

    listUserMessages = (count?: number, cursor?: number) => {
        return this.client.get<any[]>(`/self/messages?count=${count || ""}&cursor=${cursor || ""}`)
    }
    
    messageUser = (uid: string, data: any) => {
        return this.client.post(`/self/messages/${uid}`, data)
    }

    userInfo = (uid: string) => {
        return this.client.get<Partial<User>>(`/user/${uid}`)
    }

    getSharedFile = (fid: string) => {
        return this.client.get(`/file/shared/${fid}`)
    }

    getSharedFileURL = (fid: string): string => {
        return `${location.origin}/z/api/v1/file/shared/${fid}`
    }

    sharedFile = (fid: string, pid?: string) => {
        return this.client.post(`/file/shared/${fid}?pid=${pid}`, {})
    }


    deleteShareFile = (fid: string) => {
        return this.client.delete(`/file/shared/${fid}`)
    }


    getFileShortKey = (fid: string) => {
        return this.client.post(`/file/${fid}/shortkey`, {})
    }

    getFileWithShortKeyURL = (shortkey: string): string => {

        return `${location.origin}/z/api/v1/file/shortKey/${shortkey}` 
    }

}



