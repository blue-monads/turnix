import"../chunks/disclose-version.BOfWGsg9.js";import{p as n}from"../chunks/proxy.DKs42Gzo.js";import{p as i,g as _,a as d,b as f,h as g,i as c}from"../chunks/runtime.CmPLUjJJ.js";import{s as u,a as y}from"../chunks/store.BN9B0iWR.js";import{A as l}from"../chunks/index.DQ4XhlMg.js";import"../chunks/autotable.D5y0wA8l.js";/* empty css                                                     */import{p as T}from"../chunks/index.NL8p-q19.js";import{g as $}from"../chunks/entry.CytXzrzW.js";function X(s,p){i(p,!0);const o=u();let e=y(T,"$params",o).pid;const r=_("__api__");let a=c("");l(s,{data:{},get message(){return f(a)},schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"SELECT",key_name:"ptype",name:"type",options:["sideapp_v1"]},{ftype:"LONG_TEXT",key_name:"server_code",name:"Server Code"},{ftype:"LONG_TEXT",key_name:"client_code",name:"Client Code"}],name:"New Plugin",required_fields:["name"]},onSave:async m=>{const t=await r.addProjectPlugin(e,m);if(t.status!==200){g(a,n(t.data));return}$(`/z/pages/portal/project/plugins?pid=${e}`)}}),d()}export{X as component};
