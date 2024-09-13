import{s as q,a as L,e as _,g as M,c as h,b as y,f,m as $,i as D,h as m,k as B,q as Q,t as U,w as P,d as V,n as O,y as W,D as T,E as A}from"../chunks/scheduler.Dy-hT7lj.js";import{S as F,i as G,c as E,b as S,m as j,t as C,a as I,e as x,f as z,g as J,d as K}from"../chunks/index.jcCBAMyh.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.DUTEVMI9.js";import{A as Z}from"../chunks/AppBar.C2N-oVhp.js";import{T as ee,a as H,h as te}from"../chunks/index.BKLWt5jm.js";import{C as X,a as ae,l as se,s as ne,j as re}from"../chunks/index.C-Ut8Aow.js";import{N as oe}from"../chunks/index.ZlEQwogg.js";import{S as le}from"../chunks/SvgIcon.D5ofxCQw.js";import{p as ie}from"../chunks/index.NnJejm2M.js";import{R as ce}from"../chunks/Runner.1eq8diuF.js";const ue=`<!DOCTYPE html>
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
</html>`,pe=`


const myPluginDoXyz = (ctx) => {
    console.log("myPluginDoXyz/1");
};

`;function de(u){let e,a,t='<a class="anchor" href="/z/pages/portal/project">Projects</a>',o,l,s="›",i,n,d,c,r,p,g="›",v,k,b="Editor";return{c(){e=_("ol"),a=_("li"),a.innerHTML=t,o=L(),l=_("li"),l.textContent=s,i=L(),n=_("li"),d=_("a"),c=U("Plugins"),r=L(),p=_("li"),p.textContent=g,v=L(),k=_("li"),k.textContent=b,this.h()},l(N){e=h(N,"OL",{class:!0});var w=y(e);a=h(w,"LI",{class:!0,"data-svelte-h":!0}),P(a)!=="svelte-16t9qza"&&(a.innerHTML=t),o=M(w),l=h(w,"LI",{class:!0,"aria-hidden":!0,"data-svelte-h":!0}),P(l)!=="svelte-16i7nlm"&&(l.textContent=s),i=M(w),n=h(w,"LI",{});var R=y(n);d=h(R,"A",{class:!0,href:!0});var Y=y(d);c=V(Y,"Plugins"),Y.forEach(f),R.forEach(f),r=M(w),p=h(w,"LI",{class:!0,"aria-hidden":!0,"data-svelte-h":!0}),P(p)!=="svelte-16i7nlm"&&(p.textContent=g),v=M(w),k=h(w,"LI",{"data-svelte-h":!0}),P(k)!=="svelte-fudknd"&&(k.textContent=b),w.forEach(f),this.h()},h(){$(a,"class","crumb"),$(l,"class","crumb-separator"),$(l,"aria-hidden",""),$(d,"class","anchor"),$(d,"href",`/z/pages/portal/project/plugins?pid=${u[4]}`),$(p,"class","crumb-separator"),$(p,"aria-hidden",""),$(e,"class","breadcrumb")},m(N,w){D(N,e,w),m(e,a),m(e,o),m(e,l),m(e,i),m(e,n),m(n,d),m(d,c),m(e,r),m(e,p),m(e,v),m(e,k)},p:O,d(N){N&&f(e)}}}function fe(u){let e,a,t,o,l,s;return a=new le({props:{className:"w-4 h-4",name:"play"}}),{c(){e=_("button"),E(a.$$.fragment),t=U(`
            Run`),this.h()},l(i){e=h(i,"BUTTON",{class:!0});var n=y(e);S(a.$$.fragment,n),t=V(n,`
            Run`),n.forEach(f),this.h()},h(){$(e,"class","btn variant-filled btn-sm")},m(i,n){D(i,e,n),j(a,e,null),m(e,t),o=!0,l||(s=W(e,"click",u[6]),l=!0)},p:O,i(i){o||(C(a.$$.fragment,i),o=!0)},o(i){I(a.$$.fragment,i),o=!1},d(i){i&&f(e),x(a),l=!1,s()}}}function me(u){let e,a="Server Code";return{c(){e=_("span"),e.textContent=a},l(t){e=h(t,"SPAN",{"data-svelte-h":!0}),P(e)!=="svelte-ecmg0i"&&(e.textContent=a)},m(t,o){D(t,e,o)},p:O,d(t){t&&f(e)}}}function ge(u){let e,a="Client Code";return{c(){e=_("span"),e.textContent=a},l(t){e=h(t,"SPAN",{"data-svelte-h":!0}),P(e)!=="svelte-v6x68u"&&(e.textContent=a)},m(t,o){D(t,e,o)},p:O,d(t){t&&f(e)}}}function _e(u){let e,a,t,o,l,s;function i(r){u[9](r)}let n={padding:"ml-6 p-2",name:"SQL Query",value:0,$$slots:{default:[me]},$$scope:{ctx:u}};u[0]!==void 0&&(n.group=u[0]),e=new H({props:n}),T.push(()=>z(e,"group",i));function d(r){u[10](r)}let c={name:"UI",value:1,$$slots:{default:[ge]},$$scope:{ctx:u}};return u[0]!==void 0&&(c.group=u[0]),o=new H({props:c}),T.push(()=>z(o,"group",d)),{c(){E(e.$$.fragment),t=L(),E(o.$$.fragment)},l(r){S(e.$$.fragment,r),t=M(r),S(o.$$.fragment,r)},m(r,p){j(e,r,p),D(r,t,p),j(o,r,p),s=!0},p(r,p){const g={};p&32768&&(g.$$scope={dirty:p,ctx:r}),!a&&p&1&&(a=!0,g.group=r[0],A(()=>a=!1)),e.$set(g);const v={};p&32768&&(v.$$scope={dirty:p,ctx:r}),!l&&p&1&&(l=!0,v.group=r[0],A(()=>l=!1)),o.$set(v)},i(r){s||(C(e.$$.fragment,r),C(o.$$.fragment,r),s=!0)},o(r){I(e.$$.fragment,r),I(o.$$.fragment,r),s=!1},d(r){r&&f(t),x(e,r),x(o,r)}}}function he(u){let e,a,t;function o(s){u[8](s)}let l={lang:te()};return u[2]!==void 0&&(l.value=u[2]),e=new X({props:l}),T.push(()=>z(e,"value",o)),{c(){E(e.$$.fragment)},l(s){S(e.$$.fragment,s)},m(s,i){j(e,s,i),t=!0},p(s,i){const n={};!a&&i&4&&(a=!0,n.value=s[2],A(()=>a=!1)),e.$set(n)},i(s){t||(C(e.$$.fragment,s),t=!0)},o(s){I(e.$$.fragment,s),t=!1},d(s){x(e,s)}}}function $e(u){let e,a,t;function o(s){u[7](s)}let l={extensions:[ae({activateOnTyping:!0,override:[se,ne(u[5])]})],lang:re()};return u[1]!==void 0&&(l.value=u[1]),e=new X({props:l}),T.push(()=>z(e,"value",o)),{c(){E(e.$$.fragment)},l(s){S(e.$$.fragment,s)},m(s,i){j(e,s,i),t=!0},p(s,i){const n={};!a&&i&2&&(a=!0,n.value=s[1],A(()=>a=!1)),e.$set(n)},i(s){t||(C(e.$$.fragment,s),t=!0)},o(s){I(e.$$.fragment,s),t=!1},d(s){x(e,s)}}}function ve(u){let e,a,t,o;const l=[$e,he],s=[];function i(n,d){return n[0]===0?0:n[0]===1?1:-1}return~(a=i(u))&&(t=s[a]=l[a](u)),{c(){e=_("div"),t&&t.c(),this.h()},l(n){e=h(n,"DIV",{class:!0});var d=y(e);t&&t.l(d),d.forEach(f),this.h()},h(){$(e,"class","max-h-[40vh] md:max-h-[90vh] overflow-auto")},m(n,d){D(n,e,d),~a&&s[a].m(e,null),o=!0},p(n,d){let c=a;a=i(n),a===c?~a&&s[a].p(n,d):(t&&(J(),I(s[c],1,1,()=>{s[c]=null}),K()),~a?(t=s[a],t?t.p(n,d):(t=s[a]=l[a](n),t.c()),C(t,1),t.m(e,null)):t=null)},i(n){o||(C(t),o=!0)},o(n){I(t),o=!1},d(n){n&&f(e),~a&&s[a].d()}}}function be(u){let e,a,t,o,l,s,i,n,d;return e=new Z({props:{$$slots:{trail:[fe],lead:[de]},$$scope:{ctx:u}}}),l=new ee({props:{$$slots:{panel:[ve],default:[_e]},$$scope:{ctx:u}}}),n=new ce({props:{client_code:u[3]}}),{c(){E(e.$$.fragment),a=L(),t=_("div"),o=_("div"),E(l.$$.fragment),s=L(),i=_("div"),E(n.$$.fragment),this.h()},l(c){S(e.$$.fragment,c),a=M(c),t=h(c,"DIV",{class:!0});var r=y(t);o=h(r,"DIV",{class:!0});var p=y(o);S(l.$$.fragment,p),p.forEach(f),s=M(r),i=h(r,"DIV",{class:!0});var g=y(i);S(n.$$.fragment,g),g.forEach(f),r.forEach(f),this.h()},h(){$(o,"class","flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"),$(i,"class","flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"),$(t,"class","flex flex-col md:flex-row w-full h-[94vh]")},m(c,r){j(e,c,r),D(c,a,r),D(c,t,r),m(t,o),j(l,o,null),m(t,s),m(t,i),j(n,i,null),d=!0},p(c,[r]){const p={};r&32780&&(p.$$scope={dirty:r,ctx:c}),e.$set(p);const g={};r&32775&&(g.$$scope={dirty:r,ctx:c}),l.$set(g);const v={};r&8&&(v.client_code=c[3]),n.$set(v)},i(c){d||(C(e.$$.fragment,c),C(l.$$.fragment,c),C(n.$$.fragment,c),d=!0)},o(c){I(e.$$.fragment,c),I(l.$$.fragment,c),I(n.$$.fragment,c),d=!1},d(c){c&&(f(a),f(t)),x(e,c),x(l),x(n)}}}function we(u,e,a){let t;B(u,ie,b=>a(11,t=b));let o=0,l=pe,s=ue,i="";const n=t.pid,d=Q("__api__");oe(d);const c={},r=()=>{a(3,i=s)};function p(b){l=b,a(1,l)}function g(b){s=b,a(2,s)}function v(b){o=b,a(0,o)}function k(b){o=b,a(0,o)}return[o,l,s,i,n,c,r,p,g,v,k]}class Me extends F{constructor(e){super(),G(this,e,we,be,q,{})}}export{Me as component};
