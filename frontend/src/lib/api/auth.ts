import axios from "axios"


export const signUpDirect = (data: any) => {
    return axios.post(location.origin + "/z/auth/signup/direct", data)
}


export const signUpInvite = (data: any) => {
    return axios.post(location.origin + "/z/auth/signup/invite", data)
}

export const login = (data: any) => {
    return axios.post(location.origin + "/z/auth/login", data)
}

