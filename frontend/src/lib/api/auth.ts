import axios from "axios"
import { BASE_URL } from "./common"


export const signUpDirect = (data: any) => {
    return axios.post(BASE_URL + "auth/signup/direct", data)
}


export const signUpInvite = (data: any) => {
    return axios.post(BASE_URL + "auth/signup/invite", data)
}

export const login = (data: any) => {
    return axios.post(BASE_URL + "auth/login", data)
}

