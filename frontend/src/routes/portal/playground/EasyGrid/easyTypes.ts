

export interface GridOptions {
    columns: Column[]
    onLoad: LoaderFunc
    actions?: Action[]
    key: string
    enableSort?: boolean
    enableSidebar?: boolean
    enablePagination?: boolean
    handle?: GridHandle

}


export interface GridHandle {
    reload: () => void
    next: () => void
    previous: () => void
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
    filterModels: Record<string, FilterModel[]>
    orderBy?: string
    orderDirection?: "asc" | "desc"
    pageId?: number
    pageSize?: number
}


export interface FilterModel {
    fiterType: ValueType
    operator: OperatorType
    value: string[]
    depth?: number
}

export interface OperatorValue {
    operator: OperatorType
    value: string
}


type OperatorInputType = "novalue" | "value" | "range"

type OperatorMapValue = [OperatorInputType]

const OPERATORS_MAP= {
    "equal": ["value"] as OperatorMapValue,
    "not_equal": ["value"] as OperatorMapValue,
    "contains": ["value"] as OperatorMapValue,
    "not_contains": ["value"] as OperatorMapValue,
    "starts_with": ["value"] as OperatorMapValue,
    "ends_with": ["value"] as OperatorMapValue,
    "is_null": ["novalue"] as OperatorMapValue,
    "is_not_null": ["novalue"] as OperatorMapValue,
    "between": ["range"] as OperatorMapValue,
    "not_between": ["range"] as OperatorMapValue,
    "greater_than": ["value"] as OperatorMapValue,
    "less_than": ["value"] as OperatorMapValue,
    "greater_than_or_equal": ["value"] as OperatorMapValue,
    "less_than_or_equal": ["value"] as OperatorMapValue,
} as const;

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
    disableSort?: boolean

    cssClasses?: string
    styles?: Record<string, any>
    minWidth?: string
    maxWidth?: string
}



export type RendererType = "text" | "date" | "number" | "currency" | "autocolor" | "relativedate" | "lookup" | string
