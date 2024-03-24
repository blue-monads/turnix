
import { goto } from '$app/navigation';


export const gotoPorjects = () => {
    goto("/portal/projects")
}

export const gotoAddProject = () => {
    goto("/portal/projects/add")
}

export const gotoEditProject = (pid: string) => {
    goto(`/portal/projects/edit?pid=${pid}`)
}


// ptypes

export const gotoProjectLactions = (pid: string) => {
    goto(`/portal/ptypes/onloop?pid=${pid}`)
}

export const gotoAddProjectLactions = (pid: string) => {
    goto(`/portal/ptypes/onloop/add?pid=${pid}`)
}

export const gotoEditProjectLactions = (pid: string) => {
    goto(`/portal/ptypes/onloop/edit`)
}