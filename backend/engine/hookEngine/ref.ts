

const fetchSync = (req: FetchRequest) : FetchResponse | undefined => { return undefined }


interface Runtime {
	fetchResolve(reqs: FetchRequest[] ): FetchResponse[]
}



interface FetchRequest {
	method: string,
	url: string,
	data: any,
	headers: Record<string, string>
}


interface FetchResponse {
    status: 200,
    statusText: string,
    data: () => any 
}



const resp = fetchSync({
	method: "post",
	url: "http://example.com",
	data: {},
	headers: {
		"Context-Type": "application/json"
	}
})

