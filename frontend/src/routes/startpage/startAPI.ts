import axios from "axios"

export const buildApi = (base: string) => {

    const client = axios.create({
        baseURL: base,
        timeout: 10000,
    });

    client.interceptors.request.use(async (config) => {
        


        return config;
    },  (error) => {
        return Promise.reject(error);
    });

    


    return client;
}