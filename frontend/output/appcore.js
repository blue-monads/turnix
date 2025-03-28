// attach to the window object
window.appcore = {
    isLoggedIn: () => {
        return false;
    },

    getUserInfo: () => {
        return null;
    },

    getBaseUrl: () => {

    },

    redirectToLogin: () => {
        window.location.href = "/z/pages/auth/login";
    },

    redirectToHome: () => {
        window.location.href = "/";
    },
}


