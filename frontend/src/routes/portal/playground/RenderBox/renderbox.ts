export interface IFrameIPCMessage {
    type: "sql_query" | "api_call" | "ping" | string;
    name?: string;
    data: any;
    msgId: number;
}

export interface RenderBoxHandle {
    reload: () => void;
}
