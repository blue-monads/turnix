import { API } from "./api"

export let api: API


export const initApi = () => {
    if (!api) {
        api = new API()
    }

}