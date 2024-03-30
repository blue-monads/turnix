
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

export const gotoProjectOnloopTemplates = (pid: string) => {
    goto(`/portal/ptypes/onloop/template?pid=${pid}`)
}

export const gotoProjectOnloopTemplateAdd = (pid: string) => {
    goto(`/portal/ptypes/onloop/template/add?pid=${pid}`)
}

export const gotoProjectOnloopTemplateEdit = (pid: string, tid: string) => {
    goto(`/portal/ptypes/onloop/template/edit?pid=${pid}&tid=${tid}`)
}

export const gotoProjectOnloopTemplateWatch = (pid: string, tid: string) => {
    goto(`/portal/ptypes/onloop/template/watch?pid=${pid}&tid=${tid}`)
}


export const gotoProjectOnloopQueues = (pid: string, tid: string) => {
    goto(`/portal/ptypes/onloop/queue?pid=${pid}&tid=${tid}`)
}

export const gotoProjectOnloopQueueAdd = (pid: string, tid: string) => {
    goto(`/portal/ptypes/onloop/queue/add?pid=${pid}&tid=${tid}`)
}


export const gotoProjectOnloopQueuePreview = (pid: string, tid: string, mid: string) => {
    goto(`/portal/ptypes/onloop/queue/preview?pid=${pid}&tid=${tid}&mid=${mid}`)
}

