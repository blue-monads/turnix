import{s as i,y as p}from"../chunks/scheduler.CPmtgLt4.js";import{S as c,i as f,c as u,b as _,m as l,t as $,a as d,d as g}from"../chunks/index.CsITmn2N.js";import{b as y}from"../chunks/index.fwYLBDZW.js";import{A as S}from"../chunks/auto_form.LV0wk951.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.XpbzWZVH.js";/* empty css                                                     */function P(a){let t,o;return t=new S({props:{data:{},message:a[0],schema:{fields:[{ftype:"TEXT",key_name:"name",name:"Name"},{ftype:"SELECT",key_name:"ptype",name:"type",options:["onloop"]}],name:"Add Project",required_fields:["name"]},onSave:a[2]}}),{c(){u(t.$$.fragment)},l(e){_(t.$$.fragment,e)},m(e,n){l(t,e,n),o=!0},p(e,[n]){const s={};n&1&&(s.message=e[0]),n&1&&(s.onSave=e[2]),t.$set(s)},i(e){o||($(t.$$.fragment,e),o=!0)},o(e){d(t.$$.fragment,e),o=!1},d(e){g(t,e)}}}function h(a,t,o){const e=p("__api__");let n="";return[n,e,async m=>{const r=await e.addProject(m);if(r.status!==200){o(0,n=r.data);return}y()}]}class b extends c{constructor(t){super(),f(this,t,h,P,i,{})}}export{b as component};
