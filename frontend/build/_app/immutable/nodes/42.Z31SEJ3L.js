import{s as d,y as w,n as A}from"../chunks/scheduler.CPmtgLt4.js";import{S as _,i as $,c as i,b as l,m,t as p,a as f,d as u}from"../chunks/index.CsITmn2N.js";import{A as h}from"../chunks/autotable.azE3OIp3.js";import{P as C}from"../chunks/pagelayout.DsbCfEMs.js";/* empty css                                                     */import{c as N,d as k,e as L}from"../chunks/index.fwYLBDZW.js";function O(a){let e,n;return e=new h({props:{action_key:"id",key_names:[["id","ID"],["name","Name"],["ptype","Project type"],["owned_by","Owner"]],datas:a[1],color:["ptype"],actions:[{Name:"explore",Class:"bg-green-400",Action:q},{Name:"edit",Action:a[5]},{Name:"hooks",Class:"bg-yellow-400",Action:a[6]},{Name:"delete",Class:"bg-red-400",Action:a[7]}]}}),{c(){i(e.$$.fragment)},l(t){l(e.$$.fragment,t)},m(t,o){m(e,t,o),n=!0},p(t,o){const r={};o&2&&(r.datas=t[1]),e.$set(r)},i(t){n||(p(e.$$.fragment,t),n=!0)},o(t){f(e.$$.fragment,t),n=!1},d(t){u(e,t)}}}function S(a){let e,n;return e=new C({props:{title:a[0],actions:[{name:"add",actionFn:N}],$$slots:{default:[O]},$$scope:{ctx:a}}}),{c(){i(e.$$.fragment)},l(t){l(e.$$.fragment,t)},m(t,o){m(e,t,o),n=!0},p(t,[o]){const r={};o&1&&(r.title=t[0]),o&258&&(r.$$scope={dirty:o,ctx:t}),e.$set(r)},i(t){n||(p(e.$$.fragment,t),n=!0)},o(t){f(e.$$.fragment,t),n=!1},d(t){u(e,t)}}}const q=async(a,e)=>{};function x(a,e,n){let{ptype:t=""}=e,{title:o="projects"}=e;const r=w("__api__");let g=[];const c=async()=>{const s=await r.listProjects(t);s.status===200&&n(1,g=s.data)};c();const y=async s=>{k(s)},P=async(s,b)=>{L(b.ptype,s)},j=async s=>{await r.removeProject(s),c()};return a.$$set=s=>{"ptype"in s&&n(4,t=s.ptype),"title"in s&&n(0,o=s.title)},[o,g,r,c,t,y,P,j]}class D extends _{constructor(e){super(),$(this,e,x,S,d,{ptype:4,title:0})}}function E(a){let e,n;return e=new D({props:{title:"OnLoop Projects",ptype:"onloop"}}),{c(){i(e.$$.fragment)},l(t){l(e.$$.fragment,t)},m(t,o){m(e,t,o),n=!0},p:A,i(t){n||(p(e.$$.fragment,t),n=!0)},o(t){f(e.$$.fragment,t),n=!1},d(t){u(e,t)}}}class G extends _{constructor(e){super(),$(this,e,null,E,d,{})}}export{G as component};
