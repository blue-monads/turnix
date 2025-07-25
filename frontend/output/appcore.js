window.appcore = {
    _onInit: [],

    onInit: (callback) => {
        appcore._onInit.push(callback);
    },

    isLoggedIn: () => {
        return !!localStorage.getItem("access_token");
    },

    mustGetAccessToken: () => {
        const accessToken = localStorage.getItem("access_token");

        if (!accessToken) {
            appcore.redirectToLogin();
            return null;
        }

        return accessToken;
    },

    ensureLoggedIn: () => {
        if (!appcore.isLoggedIn()) {
            appcore.redirectToLogin();
            return false;
        }
        return true;
    },

    getHttpClient: () => {
        const accessToken = appcore.mustGetAccessToken();

        if (typeof axios === "undefined") {
            throw new Error("axios is not defined");
        }

        return axios.create({
            baseURL: appcore.getBaseURL(),
            headers: {
                "Authorization": accessToken
            }
        });
    },    
    getBaseURL: () => {
        return `${location.origin}/z/project/`;
    },

    getProjectURL: (ptype) => {
        return `${location.origin}/z/project/${ptype}/`;
    },

    redirectToLogin: () => {
        window.location.href = "/z/pages/auth/login";
    },

    redirectToHome: () => {
        window.location.href = "/";
    },

    redirectToPortal: () => {
        window.location.href = `${location.origin}/z/pages/portal`;
    },
}


window.addEventListener("DOMContentLoaded", (event) => {
    console.log("DOMContentLoaded/appcore");

    for (const callback of appcore._onInit) {
        callback();
    }
});
