
import { goto } from '$app/navigation';


export const gotoPorjects = () => {
    goto("/z/pages/portal/project")
}

export const gotoAddProject = () => {
    goto("/z/pages/portal/project/add")
}


export const gotoEditProject = (pid: string) => {
    goto(`/z/pages/portal/project/edit?pid=${pid}`)
}


export const gotoAddProjectHook = (ptype: string, pid: string) => {
    goto(`/z/pages/portal/hooks/new?ptype=${ptype}&pid=${pid}`)
}

export const gotoEditProjectHook = (ptype: string, pid: string, hid: string) => {
    goto(`/z/pages/portal/hooks/edit?ptype=${ptype}&pid=${pid}&hid=${hid}`)
}

export const gotoProjectHooks = (ptype: string, pid: string) => {
    goto(`/z/pages/portal/hooks?ptype=${ptype}&pid=${pid}`)
}
