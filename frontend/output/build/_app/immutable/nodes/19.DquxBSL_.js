import"../chunks/uYeY98Ep.js";import{p as r,g as m,a as n,d as i,i as f,h as d}from"../chunks/CRNwc6nF.js";import{p as c}from"../chunks/Ycl9y-hy.js";import{g}from"../chunks/DUS9ZIvE.js";import{A as _}from"../chunks/D8sggRce.js";import"../chunks/g85Fs1Im.js";/* empty css                */function E(a,o){r(o,!0);const s=m("__api__");let e=d("");_(a,{data:{},get message(){return i(e)},schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"SELECT",key_name:"ptype",name:"type",options:["onloop"]}],name:"Add Project",required_fields:["name"]},onSave:async p=>{const t=await s.addProject(p);if(t.status!==200){f(e,c(t.data));return}g()}}),n()}export{E as component};
