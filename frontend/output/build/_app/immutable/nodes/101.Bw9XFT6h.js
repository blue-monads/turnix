import{c as x,a as A}from"../chunks/Celh2SBM.js";import{p as L,i as r,g as T,f as h,a as k,h as m,d as p}from"../chunks/BC9Cf8Q3.js";import{i as w}from"../chunks/BgSg1_cg.js";import{p as i}from"../chunks/CKwprqj2.js";import{s as C,a as n}from"../chunks/BvXhjoeW.js";import{A as S}from"../chunks/CMsATdMh.js";import"../chunks/C6IzKD7A.js";import{L as U}from"../chunks/JNyEhxU_.js";import{p as K}from"../chunks/DK_nkBS0.js";import{N as M}from"../chunks/DQVP37Ea.js";import{p as N}from"../chunks/C1rFHlC-.js";function H(d,f){L(f,!0);const[a,u]=C(),l=()=>n(N,"$page",a),c=()=>n(K,"$params",a);let _=l().params.pid,g=c().mid,y="",t=r(i({})),s=r(!0);const E=M(T("__api__"));(async()=>{const e=await E.getQueueMessage(_,g);e.status===200&&(m(t,i(e.data)),m(s,!1))})();var o=x(),v=h(o);{var $=e=>{U(e,{})},b=e=>{S(e,{get data(){return p(t)},message:y,schema:{fields:[{ftype:"KEY_VALUE_TEXT",key_name:"contents",name:"Contents",disabled:!0},{ftype:"SELECT",key_name:"status",name:"Status",options:["submited","processed","denied"],disabled:!0},{ftype:"KEY_VALUE_TEXT",key_name:"extrameta",name:"Extra Meta",disabled:!0}],name:"Preview Queue message",required_fields:[]},onSave:async Q=>{}})};w(v,e=>{p(s)?e($):e(b,!1)})}A(d,o),k(),u()}export{H as component};
