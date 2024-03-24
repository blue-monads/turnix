import axios from "axios"
import { browser } from '$app/environment'
import type { AxiosInstance } from "axios"
import { BASE_URL } from "./common"



export class API {
    client: AxiosInstance

    constructor() {
        let access_token;

        if (browser) {
            access_token = localStorage.getItem("access_token")
            if (!access_token) {
                window.location.pathname = "/auth/login"
            }
        }

        this.client = axios.create({
            baseURL: BASE_URL + "api/v1/",
            headers: {
                "Authorization": access_token
            }
        })
    }


    // projects

    listProjects = () => {
        return this.client.get("/project")
    }

    addProject = (data: object) => {
        return this.client.post("/project", data)
    }

    updateProject = (pid: number, data: object) => {
        return this.client.post(`/project/${pid}`, data)
    }

    getProject = (pid: number) => {
        return this.client.get(`/project/${pid}`)
    }

    removeProject = (pid: number) => {
        return this.client.delete(`/project/${pid}`)
    }

    inviteUserToPoject = (pid: string) => {
        return this.client.post(`/project/${pid}/user`)
    }

    removeUserFromPoject = (pid: string) => {
        return this.client.delete(`/project/${pid}/user`)
    }

    // templates

    listTemplates = (pid: string) => {
        return this.client.get(`/pt/laction/template/${pid}`)
    }

    addTemplate = (pid: string) => {
        return this.client.post(`/pt/laction/template/${pid}`)
    }

    updateTemplate = (pid: string, tid: string, data: object) => {
        return this.client.post(`/pt/laction/template/${pid}/${tid}`, data)
    }

    getTemplate = (pid: string, tid: string) => {
        return this.client.get(`/pt/laction/template/${pid}/${tid}`)
    }

    removeTemplate = (pid: string, tid: string) => {
        return this.client.delete(`/pt/laction/template/${pid}/${tid}`)
    }

    // queue

    pushQueueMessage = (pid: string, tid: string, data: object) => {
        return this.client.post(`/pt/laction/template/${pid}/${tid}/push`, data)
    }

    listQueueMessages = (pid: string) => {
        return this.client.get(`/pt/laction/queue/${pid}`)
    }


    addQueueMessage = (pid: string, data: object) => {
        return this.client.post(`/pt/laction/queue/${pid}`, data)
    }

    updateQueueMessage = (pid: string, qid: string, data: object) => {
        return this.client.post(`/pt/laction/queue/${pid}/${qid}`, data)
    }

    getQueueMessage = (pid: string, qid: string) => {
        return this.client.get(`/pt/laction/queue/${pid}/${qid}`)
    }

    removeUpdateQueueMessage = (pid: string, qid: string) => {
        return this.client.delete(`/pt/laction/queue/${pid}/${qid}`)
    }

    queryQueueMessage = (pid: string) => {
        return this.client.post(`/pt/laction/queue/${pid}`)
    }

}



