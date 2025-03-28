

window.simpleform = {
    httpClient: null,
    _onInit: [],

    init: () => {
        const appcore = window.appcore;
        if (appcore.ensureLoggedIn()) {
            return
        }

        simpleform.httpClient = appcore.getHttpClient();
    },

    onInit: (callback) => {
        simpleform._onInit.push(callback);
    },

    getClient: () => {
        if (!simpleform.httpClient) {
            simpleform.init();
        }

        return simpleform.httpClient;
    },

    listTemplates: async () => {
        return getClient().get("/templates");
    },

    createTemplate: async (template) => {
        return getClient().post("/templates", template);
    },

    getTemplate: async (id) => {
        return getClient().get(`/templates/${id}`);
    },

    updateTemplate: async (id, template) => {
        return getClient().put(`/templates/${id}`, template);
    },

    deleteTemplate: async (id) => {
        return getClient().delete(`/templates/${id}`);
    },
}


window.addEventListener("DOMContentLoaded", (event) => {
    console.log("DOMContentLoaded", event);
    simpleform.init();

    for (const callback of simpleform._onInit) {
        callback();
    }

});
