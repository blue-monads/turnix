import{a as r,t as n,c as J}from"../chunks/disclose-version.BOfWGsg9.js";import{p as z,i as $,h as u,b as t,g as R,a as O,f as E,c as p,s as j,d as D,r as c,e as N}from"../chunks/runtime.CmPLUjJJ.js";import{d as K}from"../chunks/events.DnYA5zBv.js";import{s as Z}from"../chunks/attributes.yGGqmlPO.js";import{p as s}from"../chunks/proxy.DKs42Gzo.js";import{s as U,a as X}from"../chunks/store.BN9B0iWR.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{A as V}from"../chunks/AppBar.pSMxueMf.js";import{C as L,h as tt}from"../chunks/index.63AqBVKj.js";import{N as H}from"../chunks/index.Dq_7tmzI.js";import{S as et}from"../chunks/SvgIcon.Cn9REWmg.js";import{p as q}from"../chunks/index.CqTR9n69.js";import{i as T}from"../chunks/if.DoU_1wcm.js";import{p as at}from"../chunks/props.IjRlZquB.js";import{T as st,a as Y}from"../chunks/Tab.DYwUpmHc.js";import{h as rt}from"../chunks/index.DtK7kxBu.js";import{j as ot,l as nt,s as it}from"../chunks/index.BW_8NPNh.js";import{R as lt}from"../chunks/Runner.BVfP6L_V.js";const B=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="//unpkg.com/alpinejs" defer><\/script>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"><\/script>
    <script>
        let pendingMessages = {}
        let messageId = 0;
        window.addEventListener("message", (ev) => {
            console.log("message", ev);
            if (ev.data !== "transfer_port") {
                console.log("wrong message", ev);
                return;
            }
            const port = ev.ports[0];
            port.onmessage = handleMessage;
            window["__parent_port__"] = port;
        }, false);
        const handleMessage = (ev) => {
            const data = ev.data;
            const messageId = data.msgId;
            const resolve = pendingMessages[messageId];
            delete pendingMessages[messageId];
            resolve(data.data);
            console.log("data", data);
        }
        const sendMessage = async (data) => {
            console.log("sendMessage", data);
            const msgId = messageId + 1;
            messageId = msgId;
            const p = new Promise((resolve) => {
                pendingMessages[msgId] = resolve;
            });
            data["msgId"] = msgId;
            const port = window["__parent_port__"]
            port?.postMessage(data);
            return p;
        }
    <\/script>
</head>
<body>
    <div class="container">
        <h1>Example Plugin UI</h1>          
    </div>
    <div class="container-xl mt-5" x-data="Loader()">
        <div class="container">
            <button id="start" class="btn btn-primary" type="button" x-on:click="load">
                <span class="icon icon-thumbs-up"></span>
                Start
            </button>    
        </div>
    </div>

    <script>
    function Loader() {
      return {
          datas: [],
          loading: true,
          async load() {
              this.loading = true;
              try {

                  const data = await sendMessage({
                      type: "ping",
                      name: "hello",
                      data: [],
                  });

                  console.log("@data", data);

                  this.data = data.data;
              } catch (error) {
                  console.error("Error fetching data:", error);
              } finally {
                  this.loading = false;
              }
          },
          formatDate(dateString) {
              return new Date(dateString).toLocaleDateString();
          }
      }
  }


    <\/script>
</body>
</html>`,Q=`


const myPluginDoXyz = (ctx) => {
    console.log("myPluginDoXyz/1");
};

`;var dt=n("<span>Server Code</span>"),pt=n("<span>Client Code</span>"),ct=n("<!> <!>",1),mt=n('<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"><!></div>'),vt=n('<div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><!></div></div>');function gt(A,y){z(y,!0);const P=U(),C=()=>X(q,"$params",P);let o=$(0),w=$(s(Q)),i=$(s(B)),h=$("");at(y,"runAction",15)(()=>{u(h,s(t(i)))}),C().pid;const M=R("__api__");H(M);const l={};var f=vt(),e=p(f),m=p(e);st(m,{children:(k,W)=>{var g=ct(),x=E(g);Y(x,{padding:"ml-6 p-2",get group(){return t(o)},set group(a){u(o,s(a))},name:"SQL Query",value:0,children:(a,b)=>{var d=dt();r(a,d)},$$slots:{default:!0}});var _=j(x,2);Y(_,{get group(){return t(o)},set group(a){u(o,s(a))},name:"UI",value:1,children:(a,b)=>{var d=pt();r(a,d)},$$slots:{default:!0}}),r(k,g)},$$slots:{default:!0,panel:(k,W)=>{var g=mt(),x=p(g);T(x,()=>t(o)===0,_=>{var a=D(()=>[tt({activateOnTyping:!0,override:[nt,it(l)]})]),b=D(ot);L(_,{get value(){return t(w)},set value(d){u(w,s(d))},get extensions(){return t(a)},get lang(){return t(b)}})},_=>{var a=J(),b=E(a);T(b,()=>t(o)===1,d=>{var F=D(rt);L(d,{get value(){return t(i)},set value(G){u(i,s(G))},get lang(){return t(F)}})},null,!0),r(_,a)}),c(g),r(k,g)}}}),c(e);var v=j(e,2),I=p(v);lt(I,{get client_code(){return t(h)}}),c(v),c(f),r(A,f),O()}var ut=n('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/project">Projects</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li><a class="anchor">Plugins</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Editor</li></ol>'),ht=n('<button class="btn variant-filled btn-sm"><!> Run</button>'),ft=n("<!> <!>",1);function Tt(A,y){z(y,!0);const P=U(),C=()=>X(q,"$params",P);s(Q),s(B);const o=C().pid,w=R("__api__");H(w);let i=$(void 0);var h=ft(),S=E(h);V(S,{$$slots:{lead:(l,f)=>{var e=ut(),m=j(p(e),4),v=p(m);Z(v,"href",`/z/pages/portal/project/plugins?pid=${o}`),c(m),N(4),c(e),r(l,e)},trail:(l,f)=>{var e=ht();e.__click=function(...v){var I;(I=t(i))==null||I.apply(this,v)};var m=p(e);et(m,{className:"w-4 h-4",name:"play"}),N(),c(e),r(l,e)}}});var M=j(S,2);gt(M,{get runAction(){return t(i)},set runAction(l){u(i,s(l))}}),r(A,h),O()}K(["click"]);export{Tt as component};
