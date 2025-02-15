import type { RootAPI } from "$lib/api";
import type { AxiosInstance } from "axios";

export interface Post {
    id: number
    title: string
    slug: string
    excerpt: string
    content: string
    author_id: number
    created_at: string
    updated_at: string
    is_published: boolean
}


export const NewZblogAPI = (api: RootAPI) => {
    return new ZBlogAPI(api.projectClient)
}

export class ZBlogAPI {
    client: AxiosInstance

    constructor(client: AxiosInstance) {
        this.client = client
    }

    // posts

    listPost = (pid: string) => {
        return this.client.get<Post[]>(`zblog/api/${pid}/post`)
    }

    addPost = (pid: string, data: Partial<Post>) => {
        return this.client.post(`zblog/api/${pid}/post`, data)
    }

    getPost = (pid: string, id: number) => {
        return this.client.get<Post>(`zblog/api/${pid}/post/${id}`)
    }

    updatePost = (pid: string, id: number, data: Partial<Post>) => {
        return this.client.post(`zblog/api/${pid}/post/${id}`, data)
    }

    deletePost = (pid: string, id: number) => {
        return this.client.delete(`zblog/api/${pid}/post/${id}`)
    }

}


