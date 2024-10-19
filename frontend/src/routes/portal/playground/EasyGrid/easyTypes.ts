

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
    filterModels: Record<string, FilterModel>
    orderBy?: string
    orderDirection?: "asc" | "desc"
}


export interface FilterModel {
    key: string
    fiterType: ValueType
    operator: OperatorType
    value: string
}

export interface OperatorValue {
    operator: OperatorType
    value: string
}


const OPERATORS_MAP = {
    "equal": "equal",
    "not_equal": "not_equal",
    "contains": "contains",
    "not_contains": "not_contains",
    "starts_with": "starts_with",
    "ends_with": "ends_with",
    "is_null": "is_null",
    "is_not_null": "is_not_null",
    "between": "between",
    "not_between": "not_between",
    "greater_than": "greater_than",
    "less_than": "less_than",
    "greater_than_or_equal": "greater_than_or_equal",
    "less_than_or_equal": "less_than_or_equal",
}

export const OPERATORS = Object.keys(OPERATORS_MAP);

export type OperatorType = keyof typeof OPERATORS_MAP;


export type ValueType = 
    "string" | 
    "number" | 
    "date" | 
    "boolean"


export interface Column {
    title?: string
    key: string
    type?: ValueType
    renderer?: RendererType
    rendererOptions?: Record<string, any>
    enableSort?: boolean

    cssClasses?: string
    styles?: Record<string, any>
    minWidth?: string
    maxWidth?: string
}



export type RendererType = "text" | "date" | "number" | "currency" | "autocolor" | "relativedate" | "lookup" | string