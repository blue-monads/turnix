

window.simpleform = {
    init: async () => {
        console.log("simpleform client loaded");
    }
}


window.addEventListener("DOMContentLoaded", (event) => {
    console.log("DOMContentLoaded", event);
    simpleform.init();
});
