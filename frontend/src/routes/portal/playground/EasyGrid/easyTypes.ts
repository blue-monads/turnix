

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


export interface FilterModel {
    key: string
    fiterType: "string" | "number" | "date" | "boolean"
    operator: OperatorType
    value: string
}

export type OperatorType = 
    "equal" | 
    "not_equal" | 
    "contains" | 
    "not_contains" | 
    "starts_with" | 
    "ends_with" | 
    "is_null" | 
    "is_not_null" |
    "between" | 
    "not_between" | 
    "greater_than" | 
    "less_than" | 
    "greater_than_or_equal" | 
    "less_than_or_equal"


export interface Column {
    title?: string
    key: string
    renderer?: RendererType
    rendererOptions?: Record<string, any>
    enableSort?: boolean

    cssClasses?: string
    styles?: Record<string, any>
    minWidth?: string
    maxWidth?: string
}



export type RendererType = "text" | "date" | "number" | "currency" | "autocolor" | "relativedate" | "lookup" | string
