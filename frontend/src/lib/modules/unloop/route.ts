import { goto } from '$app/navigation';


export const gotoProjectOnloopTemplates = (pid: string) => {
    goto(`/z/pages/portal/projects/unloop/template?pid=${pid}`)
}

export const gotoProjectOnloopTemplateAdd = (pid: string) => {
    goto(`/z/pages/portal/projects/unloop/template/add?pid=${pid}`)
}

export const gotoProjectOnloopTemplateEdit = (pid: string, tid: string) => {
    goto(`/z/pages/portal/projects/unloop/template/edit?pid=${pid}&tid=${tid}`)
}

export const gotoProjectOnloopTemplateWatch = (pid: string, tid: string) => {
    goto(`/z/pages/portal/projects/unloop/template/watch?pid=${pid}&tid=${tid}`)
}


export const gotoProjectOnloopQueues = (pid: string, tid: string) => {
    goto(`/z/pages/portal/projects/unloop/queue?pid=${pid}&tid=${tid}`)
}

export const gotoProjectOnloopQueueAdd = (pid: string, tid: string) => {
    goto(`/z/pages/portal/projects/unloop/queue/add?pid=${pid}&tid=${tid}`)
}


export const gotoProjectOnloopQueuePreview = (pid: string, tid: string, mid: string) => {
    goto(`/z/pages/portal/projects/unloop/queue/preview?pid=${pid}&tid=${tid}&mid=${mid}`)
}
