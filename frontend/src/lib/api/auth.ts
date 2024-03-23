import axios from "axios"

export const signUpDirect = (data: any) => {
    return axios.post("/auth/signup/direct", data)
}


export const signUpInvite = (data: any) => {
    return axios.post("/auth/signup/invite", data)
}

export const login = (data: any) => {
    return axios.post("/auth/login", data)
}

