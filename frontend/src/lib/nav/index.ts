
import { goto } from '$app/navigation';


export const gotoPorjects = () => {
    goto("/portal/projects")
}

export const gotoAddProject = () => {
    goto("/portal/projects/add")
}

export const gotoEditProject = () => {
    goto("/portal/projects/edit")
}

export const gotoProjectLactions = () => {
    goto("/portal/projects/ptype/laction")
}

export const gotoAddProjectLactions = () => {
    goto("/portal/projects/ptype/laction/add")
}

export const gotoEditProjectLactions = () => {
    goto("/portal/projects/ptype/laction/edit")
}