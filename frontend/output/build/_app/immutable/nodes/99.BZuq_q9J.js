import"../chunks/Celh2SBM.js";import{p as g,g as f,i as $,a as _,h as y,d as A}from"../chunks/BC9Cf8Q3.js";import{p as n}from"../chunks/CKwprqj2.js";import{s as P,a as m}from"../chunks/BvXhjoeW.js";import{P as b}from"../chunks/CMsATdMh.js";import{A as w}from"../chunks/C6IzKD7A.js";/* empty css                */import{g as Q,a as C}from"../chunks/C_5QEsW9.js";import{p as N}from"../chunks/DK_nkBS0.js";import{N as h}from"../chunks/DQVP37Ea.js";import{p as v}from"../chunks/C1rFHlC-.js";function E(u,c){g(c,!0);const[s,d]=P(),l=()=>m(v,"$page",s),o=()=>m(N,"$params",s);let t=l().params.pid;const r=h(f("__api__"));let p=$(n([]));const i=async()=>{const a=await r.listQueueMessages(t);a.status===200&&y(p,n(a.data))};i(),b(u,{title:"Onloop Queue Messages",actions:[{name:"add",icon:"plus",actionFn:()=>{Q(t,o().tid)}}],children:(a,x)=>{w(a,{action_key:"id",key_names:[["id","ID"],["submitter","Submitter"],["status","Status"],["createdAt","Created At"]],get datas(){return A(p)},color:["ttype"],actions:[{Name:"preview",Class:"bg-green-400",Action:async(e,M)=>{C(t,o().tid,e)}},{Name:"delete",Class:"bg-red-400",Action:async e=>{await r.removeUpdateQueueMessage(t,e),i()}}]})},$$slots:{default:!0}}),_(),d()}export{E as component};
