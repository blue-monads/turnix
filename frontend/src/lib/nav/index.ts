
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
