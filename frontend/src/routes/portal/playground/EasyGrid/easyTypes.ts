

export interface GridOptions {
    columns: Column[]
    onLoad: LoaderFunc
    actions?: Action[]
    key: string
    enableSort?: boolean
}

export type LoaderFunc = (params: LoaderParams) => Promise<Record<string, any>[]> 


export interface Action {
    name: string
    icon?: string
    action: (id: any, data: any) => Promise<void> | void    
}


export interface LoaderParams {    
    loadType: "next" | "initial" | "previous"
    maxId?: number
    minId?: number
    activeColumns: string[]
    filterModels: Record<string, any>
    orderBy?: string
    orderDirection?: "asc" | "desc"
}


export interface Column {
    title?: string
    key: string
    renderer?: RendererType
    rendererOptions?: Record<string, any>
    enableSort?: boolean
}



export type RendererType = "text" | "date" | "number" | "currency" | "autocolor" | "relativedate" | "lookup" | string
