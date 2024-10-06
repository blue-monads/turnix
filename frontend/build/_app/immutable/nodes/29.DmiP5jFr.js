import{a as r,t as c,c as H}from"../chunks/disclose-version.wcsHG7oF.js";import{p as q,i as u,g as B,f as w,a as Q,s as h,c as i,b as t,h as v,r as d,e as D,d as I}from"../chunks/runtime.D6lgFUio.js";import{d as W}from"../chunks/utils.Dh-XvBAd.js";import{i as E}from"../chunks/if.lkQ-IRrr.js";import{s as F}from"../chunks/attributes.BFGPVQjG.js";import{p}from"../chunks/proxy.r5-P5EHK.js";import{s as G,a as J}from"../chunks/store.BaapSzAp.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DbB_HC5r.js";import{A as K}from"../chunks/AppBar.3srT6C5i.js";import{T as Z,a as N,h as V}from"../chunks/index.iPCAaS7f.js";import{C as A,a as ee,j as ae,l as te,s as se}from"../chunks/index.1_1nzF9G.js";import{N as re}from"../chunks/index.Duto3aX0.js";import{S as oe}from"../chunks/SvgIcon.Cj01WnO8.js";import{p as ne}from"../chunks/index.xGEmn6Hr.js";import{R as le}from"../chunks/Runner.DB5YsYnC.js";const ie=`<!DOCTYPE html>
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
</html>`,de=`


const myPluginDoXyz = (ctx) => {
    console.log("myPluginDoXyz/1");
};

`;var pe=c('<ol class="breadcrumb"><li class="crumb"><a class="anchor" href="/z/pages/portal/project">Projects</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li><a class="anchor">Plugins</a></li> <li class="crumb-separator" aria-hidden="">&rsaquo;</li> <li>Editor</li></ol>'),ce=(x,f,_)=>{v(f,p(t(_)))},me=c('<button class="btn variant-filled btn-sm"><!> Run</button>'),ge=c("<span>Server Code</span>"),ve=c("<span>Client Code</span>"),ue=c("<!> <!>",1),he=c('<div class="max-h-[40vh] md:max-h-[90vh] overflow-auto"><!></div>'),fe=c('<!> <div class=" flex flex-col md:flex-row w-full h-[94vh]"><div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"><!></div> <div class="flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"><!></div></div>',1);function Ne(x,f){q(f,!0);const _=G(),T=()=>J(ne,"$params",_);let m=u(0),j=u(p(de)),b=u(p(ie)),P=u("");const L=T().pid,Y=B("__api__");re(Y);const z={};var C=fe(),S=w(C);K(S,{$$slots:{lead:(o,$)=>{var e=pe(),s=h(i(e),4),n=i(s);F(n,"href",`/z/pages/portal/project/plugins?pid=${L}`),d(s),D(4),d(e),r(o,e)},trail:(o,$)=>{var e=me();e.__click=[ce,P,b];var s=i(e);oe(s,{className:"w-4 h-4",name:"play"}),D(),d(e),r(o,e)}}});var M=h(S,2),y=i(M),R=i(y);Z(R,{children:(o,$)=>{var e=ue(),s=w(e);N(s,{padding:"ml-6 p-2",get group(){return t(m)},set group(a){v(m,p(a))},name:"SQL Query",value:0,children:(a,g)=>{var l=ge();r(a,l)},$$slots:{default:!0}});var n=h(s,2);N(n,{get group(){return t(m)},set group(a){v(m,p(a))},name:"UI",value:1,children:(a,g)=>{var l=ve();r(a,l)},$$slots:{default:!0}}),r(o,e)},$$slots:{default:!0,panel:(o,$)=>{var e=he(),s=i(e);E(s,()=>t(m)===0,n=>{var a=I(()=>[ee({activateOnTyping:!0,override:[te,se(z)]})]),g=I(ae);A(n,{get value(){return t(j)},set value(l){v(j,p(l))},get extensions(){return t(a)},get lang(){return t(g)}})},n=>{var a=H(),g=w(a);E(g,()=>t(m)===1,l=>{var U=I(V);A(l,{get value(){return t(b)},set value(X){v(b,p(X))},get lang(){return t(U)}})},null,!0),r(n,a)}),d(e),r(o,e)}}}),d(y);var k=h(y,2),O=i(k);le(O,{get client_code(){return t(P)}}),d(k),d(M),r(x,C),Q()}W(["click"]);export{Ne as component};
