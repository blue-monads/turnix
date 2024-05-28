export const code = `


/**
 * @typedef {Object} Context
 * @property {number} event_id - The ID of the event.
 * @property {string} event_type - The type of the event.
 * @property {string} hook_type - The type of the hook.
 * @property {number} session_user - The ID of the session user.
 * @property {number} runas_user - The ID of the user running as another user.
 * @property {function(): Object.<string, any>} getData - Function to get data.
 * @property {function(): [string, string]} getUserToken - Function to get the user token.
 * 
 * Handles the given context.
 *
 * @param {Context} ctx - The context object.
 */
const handle = (ctx) => {
  
};

`


export const completionObject = {
    "ctx": {
        event_id: 0,
        event_type: "",
        hook_type: "",
        session_user: 0,
        runas_user: 0,
        getUserToken: () => ["", ""],
        getData: () => {}
    },
    "fetchSync": null,
}