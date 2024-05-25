import { derived, writable } from "svelte/store";
import { getStores, navigating, page, updated } from '$app/stores';

export const params = derived([navigating, page], ([ndata, p]) => {
    let params
    if (ndata && ndata.to) {
        params = ndata.to.url.searchParams
    } else {
        params = new URLSearchParams(location.search)
    }

    let o: Record<string, string> = {}
    if (p.params) {
        o["pid"] = p.params["pid"]
    }

    
    params.forEach((v, k) => { o[k] = v })
    return o
})


params.subscribe((v) => {
    console.log("@params", v)
})