import{c as de,a as k,t as B}from"../chunks/disclose-version.BOfWGsg9.js";import{al as Bt,p as re,aC as ce,aD as bt,f as Z,a as ne,aO as gt,h as v,aE as O,b as a,at as f,c as n,r as o,t as M,s as y,a2 as Se,i as V,d as lt,e as Pe,ae as Et,n as At}from"../chunks/runtime.CmPLUjJJ.js";import{s as Lt}from"../chunks/snippet.UtI6fqho.js";import{i as Mt,g as be}from"../chunks/stores.C0Vv8-i4.js";import{w as Ot}from"../chunks/index.CdZvgsHB.js";import{i as Ft,a as _t,g as ht}from"../chunks/stores.BdRXiLpF.js";import{p as wt}from"../chunks/ProgressBar.svelte_svelte_type_style_lang.Kk8Ai2k7.js";import{s as me}from"../chunks/render.CZQ66Y_F.js";import{i as q}from"../chunks/if.DoU_1wcm.js";import{k as Ht}from"../chunks/key.B3I6Q0YW.js";import{h as it,S as pe}from"../chunks/SvgIcon.Cn9REWmg.js";import{c as mt}from"../chunks/svelte-component.DHZAe5Z3.js";import{s as Y,r as We,a as Rt}from"../chunks/attributes.yGGqmlPO.js";import{s as L}from"../chunks/class.BFWsl-fM.js";import{e as X,d as ze}from"../chunks/events.DnYA5zBv.js";import{f as Ge,t as je,a as st,A as ct}from"../chunks/autotable.D5y0wA8l.js";import{b as Oe}from"../chunks/input.BRsTNa69.js";import{b as Je}from"../chunks/this.DlUubvPq.js";import{a as qt}from"../chunks/shared.BFhNwdsm.js";import{i as Qe}from"../chunks/lifecycle.46CgxKeU.js";import{b as Ke}from"../chunks/misc.CT9jRRbj.js";import{l as xt,p,s as vt}from"../chunks/props.IjRlZquB.js";import{s as ue,a as oe}from"../chunks/store.BN9B0iWR.js";import{c as yt}from"../chunks/index-client.Caj9tY0k.js";import{a as Ut}from"../chunks/slot.D3lQyPcJ.js";import{e as dt,i as ut}from"../chunks/each.BJPUi07-.js";import{p as ve}from"../chunks/proxy.DKs42Gzo.js";import{L as Xe}from"../chunks/loader._7BlSOYJ.js";/* empty css                       */import"../chunks/index.DQ4XhlMg.js";/* empty css                                                     */import{r as Kt}from"../chunks/misc.CveY8KIp.js";import{F as Wt}from"../chunks/FilePicker.D0e64Dn9.js";import{r as Qt}from"../chunks/legacy-client.BxBrBVYa.js";import"../chunks/entry.CytXzrzW.js";import{f as Vt}from"../chunks/index.UdVxqQke.js";import{F as Yt}from"../chunks/FilePreview.CIuWA17h.js";function Gt(u,t){qt(window,["resize"],()=>t(window[u]))}const Jt="auto",Xt=!1,Zt=!0,Co=Object.freeze(Object.defineProperty({__proto__:null,csr:Zt,prerender:Jt,ssr:Xt},Symbol.toStringTag,{value:"Module"})),ea={message:"Missing Toast Message",autohide:!0,timeout:5e3},ta="toastStore";function aa(){const u=ra();return Bt(ta,u)}function oa(){const u=Math.random();return Number(u).toString(32)}function ra(){const{subscribe:u,set:t,update:g}=Ot([]),d=e=>g(r=>{if(r.length>0){const _=r.findIndex(h=>h.id===e),l=r[_];l&&(l.callback&&l.callback({id:e,status:"closed"}),l.timeoutId&&clearTimeout(l.timeoutId),r.splice(_,1))}return r});function s(e){if(e.autohide===!0)return setTimeout(()=>{d(e.id)},e.timeout)}return{subscribe:u,close:d,trigger:e=>{const r=oa();return g(_=>{e&&e.callback&&e.callback({id:r,status:"queued"}),e.hideDismiss&&(e.autohide=!0);const l={...ea,...e,id:r};return l.timeoutId=s(l),_.push(l),_}),r},freeze:e=>g(r=>(r.length>0&&clearTimeout(r[e].timeoutId),r)),unfreeze:e=>g(r=>(r.length>0&&(r[e].timeoutId=s(r[e])),r)),clear:()=>t([])}}function na(){Mt(),aa(),Ft()}function kt(u,t){const g='a[href]:not([tabindex="-1"]), button:not([tabindex="-1"]), input:not([tabindex="-1"]), textarea:not([tabindex="-1"]), select:not([tabindex="-1"]), details:not([tabindex="-1"]), [tabindex]:not([tabindex="-1"])';let d,s;function e(i){i.shiftKey&&i.code==="Tab"&&(i.preventDefault(),s.focus())}function r(i){!i.shiftKey&&i.code==="Tab"&&(i.preventDefault(),d.focus())}const _=i=>i.filter(c=>c.tabIndex>=0).sort((c,P)=>c.tabIndex===0&&P.tabIndex>0?1:c.tabIndex>0&&P.tabIndex===0?-1:c.tabIndex-P.tabIndex),l=i=>{const c=[...u.querySelectorAll("[data-focusindex]")];return!c||c.length===0?i:c.sort((P,m)=>+P.dataset.focusindex-+m.dataset.focusindex)[0]||i},h=i=>{if(t===!1)return;const c=_(Array.from(u.querySelectorAll(g)));c.length&&(d=c[0],s=c[c.length-1],i||l(d).focus(),d.addEventListener("keydown",e),s.addEventListener("keydown",r))};h(!1);function w(){d&&d.removeEventListener("keydown",e),s&&s.removeEventListener("keydown",r)}const b=(i,c)=>(i.length&&(w(),h(!0)),c),T=new MutationObserver(b);return T.observe(u,{childList:!0,subtree:!0}),{update(i){t=i,i?h(!1):w()},destroy(){w(),T.disconnect()}}}function Ne(u,t){const{transition:g,params:d,enabled:s}=t;return s?g(u,d):"duration"in d?g(u,{duration:0}):{duration:0}}function ia(u){return u*u*u}var sa=B("<header><!></header>"),la=B("<article><!></article>"),ca=B('<img alt="Modal">'),da=B('<footer><button type="button"> </button></footer>'),ua=B('<footer><button type="button"> </button> <button type="button"> </button></footer>'),fa=B('<form class="space-y-4"><input> <footer><button type="button"> </button> <button type="submit"> </button></footer></form>'),pa=B('<div data-testid="modal" role="dialog" aria-modal="true"><!> <!> <!> <!></div>'),ma=B('<div data-testid="modal-component" role="dialog" aria-modal="true"><!></div>'),va=B('<div data-testid="modal-backdrop"><div><!></div></div>');function ba(u,t){const g=xt(t,["children","$$slots","$$events","$$legacy"]);re(t,!1);const d=ue(),s=()=>oe(wt,"$prefersReducedMotionStore",d),e=()=>oe(te,"$modalStore",d),r=O(),_=O(),l=O(),h=O(),w=O(),b=yt();let T=p(t,"components",24,()=>({})),i=p(t,"position",8,"items-center"),c=p(t,"background",8,"bg-surface-100-800-token"),P=p(t,"width",8,"w-modal"),m=p(t,"height",8,"h-auto"),C=p(t,"padding",8,"p-4"),I=p(t,"spacing",8,"space-y-4"),z=p(t,"rounded",8,"rounded-container-token"),D=p(t,"shadow",8,"shadow-xl"),G=p(t,"zIndex",8,"z-[999]"),E=p(t,"buttonNeutral",8,"variant-ghost-surface"),U=p(t,"buttonPositive",8,"variant-filled"),N=p(t,"buttonTextCancel",12,"Cancel"),K=p(t,"buttonTextConfirm",12,"Confirm"),F=p(t,"buttonTextSubmit",12,"Submit"),ee=p(t,"regionBackdrop",8,""),H=p(t,"regionHeader",8,"text-2xl font-bold"),ie=p(t,"regionBody",8,"max-h-[200px] overflow-hidden"),W=p(t,"regionFooter",8,"flex justify-end space-x-2"),fe=p(t,"transitions",24,()=>!s()),De=p(t,"transitionIn",8,Ge),ge=p(t,"transitionInParams",24,()=>({duration:150,opacity:0,x:0,y:100})),_e=p(t,"transitionOut",8,Ge),Be=p(t,"transitionOutParams",24,()=>({duration:150,opacity:0,x:0,y:100}));const Ze="fixed top-0 left-0 right-0 bottom-0 bg-surface-backdrop-token p-4",et="w-full h-fit min-h-full overflow-y-auto flex justify-center",R="block overflow-y-auto",tt="w-full h-auto";let Ce=O();const Fe={buttonTextCancel:N(),buttonTextConfirm:K(),buttonTextSubmit:F()};let xe=O(),Ee=!1,Ie=O(),$=O(),Q=O("overflow-y-hidden");const te=be();function at(x){x[0].type==="prompt"&&v(Ce,x[0].value),N(x[0].buttonTextCancel||Fe.buttonTextCancel),K(x[0].buttonTextConfirm||Fe.buttonTextConfirm),F(x[0].buttonTextSubmit||Fe.buttonTextSubmit),v(xe,typeof x[0].component=="string"?T()[x[0].component]:x[0].component)}function se(x){var Ve;let le=x==null?void 0:x.clientHeight;le||(le=(Ve=x==null?void 0:x.firstChild)==null?void 0:Ve.clientHeight),le&&(le>a($)?v(Q,"overflow-y-auto"):v(Q,"overflow-y-hidden"))}function ot(x){if(!(x.target instanceof Element))return;const le=x.target.classList;(le.contains("modal-backdrop")||le.contains("modal-transition"))&&(Ee=!0)}function Tt(x){if(!(x.target instanceof Element))return;const le=x.target.classList;(le.contains("modal-backdrop")||le.contains("modal-transition"))&&Ee&&(e()[0].response&&e()[0].response(void 0),te.close(),b("backdrop",x)),Ee=!1}function He(){e()[0].response&&e()[0].response(!1),te.close()}function Pt(){e()[0].response&&e()[0].response(!0),te.close()}function It(x){x.preventDefault(),e()[0].response&&(e()[0].valueAttr!==void 0&&"type"in e()[0].valueAttr&&e()[0].valueAttr.type==="number"?e()[0].response(parseInt(a(Ce))):e()[0].response(a(Ce))),te.close()}function $t(x){e().length&&x.code==="Escape"&&He()}ce(()=>e(),()=>{e().length&&at(e())}),ce(()=>a(Ie),()=>{se(a(Ie))}),ce(()=>(e(),f(i())),()=>{var x;v(r,((x=e()[0])==null?void 0:x.position)??i())}),ce(()=>(f(ee()),f(G()),f(g),e()),()=>{var x;v(_,`${Ze} ${ee()} ${G()} ${g.class??""} ${((x=e()[0])==null?void 0:x.backdropClasses)??""}`)}),ce(()=>a(r),()=>{v(l,`${et} ${a(r)??""}`)}),ce(()=>(f(c()),f(P()),f(m()),f(C()),f(I()),f(z()),f(D()),e()),()=>{var x;v(h,`${R} ${c()} ${P()} ${m()} ${C()} ${I()} ${z()} ${D()} ${((x=e()[0])==null?void 0:x.modalClasses)??""}`)}),ce(()=>(f(i()),f(c()),f(P()),f(m()),f(C()),f(I()),f(z()),f(D()),f(E()),f(U()),f(N()),f(K()),f(F()),f(ee()),f(H()),f(ie()),f(W())),()=>{v(w,{position:i(),background:c(),width:P(),height:m(),padding:C(),spacing:I(),rounded:z(),shadow:D(),buttonNeutral:E(),buttonPositive:U(),buttonTextCancel:N(),buttonTextConfirm:K(),buttonTextSubmit:F(),regionBackdrop:ee(),regionHeader:H(),regionBody:ie(),regionFooter:W(),onClose:He})}),bt(),Qe();var ft=de();X("keydown",gt,$t);var St=Z(ft);q(St,()=>e().length>0,x=>{var le=de(),Ve=Z(le);Ht(Ve,e,Dt=>{var he=va(),Re=n(he),Ct=n(Re);q(Ct,()=>e()[0].type!=="component",ye=>{var J=pa();Je(J,S=>v(Ie,S),()=>a(Ie));var Ye=n(J);q(Ye,()=>{var S;return(S=e()[0])==null?void 0:S.title},S=>{var j=sa(),A=n(j);it(A,()=>e()[0].title,!1,!1),o(j),M(()=>L(j,`modal-header ${H()??""}`)),k(S,j)});var ae=y(Ye,2);q(ae,()=>{var S;return(S=e()[0])==null?void 0:S.body},S=>{var j=la(),A=n(j);it(A,()=>e()[0].body,!1,!1),o(j),M(()=>L(j,`modal-body ${ie()??""}`)),k(S,j)});var $e=y(ae,2);q($e,()=>{var S,j;return((S=e()[0])==null?void 0:S.image)&&typeof((j=e()[0])==null?void 0:j.image)=="string"},S=>{var j=ca();L(j,`modal-image ${tt}`),M(()=>{var A;return Y(j,"src",(A=e()[0])==null?void 0:A.image)}),k(S,j)});var qe=y($e,2);q(qe,()=>e()[0].type==="alert",S=>{var j=da(),A=n(j),Ae=n(A);o(A),o(j),M(()=>{L(j,`modal-footer ${W()??""}`),L(A,`btn ${E()??""}`),me(Ae,N())}),X("click",A,He),k(S,j)},S=>{var j=de(),A=Z(j);q(A,()=>e()[0].type==="confirm",Ae=>{var we=ua(),ke=n(we),Le=n(ke);o(ke);var Te=y(ke,2),Me=n(Te);o(Te),o(we),M(()=>{L(we,`modal-footer ${W()??""}`),L(ke,`btn ${E()??""}`),me(Le,N()),L(Te,`btn ${U()??""}`),me(Me,K())}),X("click",ke,He),X("click",Te,Pt),k(Ae,we)},Ae=>{var we=de(),ke=Z(we);q(ke,()=>e()[0].type==="prompt",Le=>{var Te=fa(),Me=n(Te);We(Me);let pt;var rt=y(Me,2),Ue=n(rt),jt=n(Ue);o(Ue);var nt=y(Ue,2),Nt=n(nt);o(nt),o(rt),o(Te),M(()=>{pt=Rt(Me,pt,{class:"modal-prompt-input input",name:"prompt",type:"text",...e()[0].valueAttr}),L(rt,`modal-footer ${W()??""}`),L(Ue,`btn ${E()??""}`),me(jt,N()),L(nt,`btn ${U()??""}`),me(Nt,F())}),Oe(Me,()=>a(Ce),zt=>v(Ce,zt)),X("click",Ue,He),X("submit",Te,It),k(Le,Te)},null,!0),k(Ae,we)},!0),k(S,j)}),o(J),M(()=>{L(J,`modal ${a(h)??""}`),Y(J,"aria-label",e()[0].title??"")}),k(ye,J)},ye=>{var J=ma();Je(J,ae=>v(Ie,ae),()=>a(Ie));var Ye=n(J);q(Ye,()=>{var ae;return(ae=a(xe))==null?void 0:ae.slot},ae=>{var $e=de(),qe=Z($e);mt(qe,()=>{var S;return(S=a(xe))==null?void 0:S.ref},(S,j)=>{j(S,vt(()=>{var A;return(A=a(xe))==null?void 0:A.props},{get parent(){return a(w)},children:(A,Ae)=>{var we=de(),ke=Z(we);it(ke,()=>{var Le;return(Le=a(xe))==null?void 0:Le.slot},!1,!1),k(A,we)},$$slots:{default:!0}}))}),k(ae,$e)},ae=>{var $e=de(),qe=Z($e);mt(qe,()=>{var S;return(S=a(xe))==null?void 0:S.ref},(S,j)=>{j(S,vt(()=>{var A;return(A=a(xe))==null?void 0:A.props},{get parent(){return a(w)}}))}),k(ae,$e)}),o(J),M(()=>{var ae;L(J,`modal contents ${((ae=e()[0])==null?void 0:ae.modalClasses)??""??""}`),Y(J,"aria-label",e()[0].title??"")}),k(ye,J)}),o(Re),o(he),M(()=>{L(he,`modal-backdrop ${a(_)??""} ${a(Q)??""}`),L(Re,`modal-transition ${a(l)??""}`)}),je(5,Re,()=>Ne,()=>({transition:De(),params:ge(),enabled:fe()})),je(6,Re,()=>Ne,()=>({transition:_e(),params:Be(),enabled:fe()})),Se(()=>X("mousedown",he,ot)),Se(()=>X("mouseup",he,Tt)),Se(()=>X("touchstart",he,function(ye){Ke.call(this,t,ye)},void 0,!0)),Se(()=>X("touchend",he,function(ye){Ke.call(this,t,ye)},void 0,!0)),je(7,he,()=>Ne,()=>({transition:st,params:{duration:150},enabled:fe()})),_t(he,(ye,J)=>kt(ye,J),()=>!0),k(Dt,he)}),k(x,le)}),Gt("innerHeight",x=>v($,x)),k(u,ft),ne()}var ga=B('<div data-testid="drawer-backdrop"><div data-testid="drawer" role="dialog" aria-modal="true"><!></div></div>');function _a(u,t){const g=xt(t,["children","$$slots","$$events","$$legacy"]);re(t,!1);const d=ue(),s=()=>oe(wt,"$prefersReducedMotionStore",d),e=()=>oe(Be,"$drawerStore",d),r=O(),_=O(),l=O(),h=O(),w=O(),b=O(),T=yt();let i=p(t,"position",12,"left"),c=p(t,"bgDrawer",12,"bg-surface-100-800-token"),P=p(t,"border",12,""),m=p(t,"rounded",12,""),C=p(t,"shadow",12,"shadow-xl"),I=p(t,"width",12,""),z=p(t,"height",12,""),D=p(t,"bgBackdrop",12,"bg-surface-backdrop-token"),G=p(t,"blur",12,""),E=p(t,"padding",12,""),U=p(t,"zIndex",8,"z-40"),N=p(t,"regionBackdrop",12,""),K=p(t,"regionDrawer",12,""),F=p(t,"labelledby",12,""),ee=p(t,"describedby",12,""),H=p(t,"duration",12,200),ie=p(t,"transitions",24,()=>!s()),W=p(t,"opacityTransition",12,!0);const fe={top:{alignment:"items-start",width:"w-full",height:"h-[50%]",rounded:"rounded-bl-container-token rounded-br-container-token"},bottom:{alignment:"items-end",width:"w-full",height:" h-[50%]",rounded:"rounded-tl-container-token rounded-tr-container-token"},left:{alignment:"justify-start",width:"w-[90%]",height:"h-full",rounded:"rounded-tr-container-token rounded-br-container-token"},right:{alignment:"justify-end",width:"w-[90%]",height:"h-full",rounded:"rounded-tl-container-token rounded-bl-container-token"}};let De=O(),ge=O(),_e=O({x:0,y:0});const Be=ht(),Ze="fixed top-0 left-0 right-0 bottom-0 flex",et="overflow-y-auto transition-transform",R={position:i(),bgBackdrop:D(),blur:G(),padding:E(),bgDrawer:c(),border:P(),rounded:m(),shadow:C(),width:I(),height:z(),opacityTransition:W(),regionBackdrop:N(),regionDrawer:K(),labelledby:F(),describedby:ee(),duration:H()};function tt($){i($.position||R.position),D($.bgBackdrop||R.bgBackdrop),G($.blur||R.blur),E($.padding||R.padding),c($.bgDrawer||R.bgDrawer),P($.border||R.border),m($.rounded||R.rounded),C($.shadow||R.shadow),I($.width||R.width),z($.height||R.height),N($.regionBackdrop||R.regionBackdrop),K($.regionDrawer||R.regionDrawer),F($.labelledby||R.labelledby),ee($.describedby||R.describedby),W($.opacityTransition||R.opacityTransition),H($.duration||R.duration)}function Ce(){switch(i()){case"top":v(_e,{x:0,y:-window.innerWidth});break;case"bottom":v(_e,{x:0,y:window.innerWidth});break;case"left":v(_e,{x:-window.innerHeight,y:0});break;case"right":v(_e,{x:window.innerHeight,y:0});break;default:console.error("Error: unknown position property value.");break}}function Fe($){$.target===a(De)?(Be.close(),T("backdrop",$)):T("drawer",$)}function xe($){e()&&$.code==="Escape"&&Be.close()}Be.subscribe($=>{$.open===!0&&(tt($),Ce())}),ce(()=>f(i()),()=>{v(r,fe[i()].alignment)}),ce(()=>(f(I()),f(i())),()=>{v(_,I()?I():fe[i()].width)}),ce(()=>(f(z()),f(i())),()=>{v(l,z()?z():fe[i()].height)}),ce(()=>(f(m()),f(i())),()=>{v(h,m()?m():fe[i()].rounded)}),ce(()=>(f(D()),f(E()),f(G()),a(r),f(N()),f(U()),f(g)),()=>{v(w,`${Ze} ${D()} ${E()} ${G()} ${a(r)} ${N()} ${U()} ${g.class??""}`)}),ce(()=>(f(c()),f(P()),f(m()),f(C()),a(_),a(l),a(h),f(K())),()=>{v(b,`${et} ${c()} ${P()} ${m()} ${C()} ${a(_)} ${a(l)} ${a(h)} ${K()}`)}),bt(),Qe();var Ee=de();X("keydown",gt,xe);var Ie=Z(Ee);q(Ie,()=>e().open===!0,$=>{var Q=ga();Je(Q,se=>v(De,se),()=>a(De));var te=n(Q);Je(te,se=>v(ge,se),()=>a(ge));var at=n(te);Ut(at,t,"default",{},null),o(te),o(Q),M(()=>{L(Q,`drawer-backdrop ${a(w)??""}`),L(te,`drawer ${a(b)??""}`),Y(te,"aria-labelledby",F()),Y(te,"aria-describedby",ee())}),je(1,te,()=>Ne,()=>({transition:Ge,params:{x:a(_e).x,y:a(_e).y,duration:H(),opacity:W()?void 0:1},enabled:ie()})),je(2,te,()=>Ne,()=>({transition:Ge,params:{x:a(_e).x,y:a(_e).y,duration:H(),opacity:W()?void 0:1,easing:ia},enabled:ie()})),Se(()=>X("mousedown",Q,Fe)),Se(()=>X("touchstart",Q,function(se){Ke.call(this,t,se)},void 0,!0)),Se(()=>X("touchend",Q,function(se){Ke.call(this,t,se)},void 0,!0)),Se(()=>X("keypress",Q,function(se){Ke.call(this,t,se)})),je(1,Q,()=>Ne,()=>({transition:st,params:{duration:H()},enabled:ie()&&W()})),je(2,Q,()=>Ne,()=>({transition:st,params:{duration:H()},enabled:ie()&&W()})),_t(Q,(se,ot)=>kt(se,ot),()=>!0),k($,Q)}),k(u,Ee),ne()}var ha=B('<a class="logo-item"><!> <span> </span></a>'),wa=B('<div class="card p-2 w-modal"><header class="header flex justify-center"><h4 class="h4"><span class="bg-gradient-to-br from-pink-500 to-violet-500 bg-clip-text text-transparent box-decoration-clone uppercase">Select New Project type</span></h4></header> <div class="logo-cloud grid-cols-1 lg:!grid-cols-3 gap-1 border p-2"></div></div>');function xa(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$modalStore",g),s=be();let e=V(ve([])),r=V(!0);(async()=>{const b=await d()[0].meta.api.listProjectTypes();b.status===200&&(v(e,ve(b.data)),v(r,!1))})();var l=de(),h=Z(l);q(h,()=>a(r),w=>{Xe(w,{})},w=>{var b=wa(),T=y(n(b),2);dt(T,21,()=>a(e),ut,(i,c)=>{var P=ha();const m=lt(()=>a(c).icon);P.__click=()=>{s.close()};var C=n(P);q(C,()=>a(m),D=>{pe(D,{get name(){return a(m)},className:"w-6 h-6"})});var I=y(C,2),z=n(I);o(I),o(P),M(()=>{Y(P,"href",`/z/pages/portal/project/new?ptype=${a(c).ptype??""}`),me(z,a(c).name)}),k(i,P)}),o(T),o(b),k(w,b)}),k(u,l),ne()}ze(["click"]);var ya=B('<div class="card p-2 w-modal"><header class="header flex justify-center"><h4 class="h4"><span class="bg-gradient-to-br from-pink-500 to-violet-500 bg-clip-text text-transparent box-decoration-clone uppercase">Options</span></h4></header> <div class="logo-cloud flex flex-col gap-1 border p-2"><a class="logo-item"><!> <span>Project Home</span></a> <a class="logo-item"><!> <span>Project Files</span></a> <a class="logo-item"><!> <span>Plugins</span></a> <a class="logo-item"><!> <span>Live Query</span></a> <a class="logo-item"><!> <span>Explore Data</span></a> <a class="logo-item"><!> <span>Users and Permissions</span></a> <a class="logo-item"><!> <span>Project Hooks</span></a> <a class="logo-item" href="/z/pages/portal/projects"><!> <span>Edit Project</span></a></div></div>');function ka(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$modalStore",g),s=be();let e=V(!0);(async()=>{const T=await d()[0].meta.api.listProjectTypes();T.status===200&&(T.data,v(e,!1))})();const _=d()[0].meta.data,l=_.id;var h=de(),w=Z(h);q(w,()=>a(e),b=>{Xe(b,{})},b=>{var T=ya(),i=y(n(T),2),c=n(i);c.__click=()=>{s.close()};var P=n(c);pe(P,{name:"document-text",className:"w-6 h-6"}),Pe(2),o(c);var m=y(c,2);m.__click=()=>{s.close()},Y(m,"href",`/z/pages/portal/project/files/${l}`);var C=n(m);pe(C,{name:"folder",className:"w-6 h-6"}),Pe(2),o(m);var I=y(m,2);I.__click=()=>{s.close()},Y(I,"href",`/z/pages/portal/project/plugins?pid=${l}`);var z=n(I);pe(z,{name:"squares-2x2",className:"w-6 h-6"}),Pe(2),o(I);var D=y(I,2);D.__click=()=>{s.close()},Y(D,"href",`/z/pages/portal/project/livequery?pid=${l}`);var G=n(D);pe(G,{name:"squares-2x2",className:"w-6 h-6"}),Pe(2),o(D);var E=y(D,2);E.__click=()=>{s.close()},Y(E,"href",`/z/pages/portal/project/explore?pid=${l}`);var U=n(E);pe(U,{name:"squares-2x2",className:"w-6 h-6"}),Pe(2),o(E);var N=y(E,2);N.__click=()=>{s.close()},Y(N,"href",`/z/pages/portal/project/users?pid=${l}`);var K=n(N);pe(K,{name:"users",className:"w-6 h-6"}),Pe(2),o(N);var F=y(N,2);F.__click=()=>{s.close()};var ee=n(F);pe(ee,{name:"code-bracket-square",className:"w-6 h-6"}),Pe(2),o(F);var H=y(F,2);H.__click=()=>{s.close()};var ie=n(H);pe(ie,{name:"pencil-square",className:"w-6 h-6"}),Pe(2),o(H),o(i),o(T),M(()=>{Y(c,"href",`/z/pages/portal/projects/${["ptype"]}?pid=${_.id??""}`),Y(F,"href",`/z/pages/portal/hooks?ptype=${_.ptype}&pid=${l}`)}),k(b,T)}),k(u,h),ne()}ze(["click"]);function Ta(u,t){re(t,!1);const g=ue(),d=()=>oe(s,"$store",g),s=be();let e=d()[0].meta.data;Qe();var r=Et(()=>[{Name:"select",icon:"plus",Action:async _=>{d()[0].meta.onClick(_),s.close()}}]);ct(u,{action_key:"id",key_names:[["id","ID"],["name","Name"],["info","Info"],["acc_type","Account type"]],datas:e,color:["acc_type"],get actions(){return a(r)}}),ne()}var Pa=B('<h4 class="h4">Pick Product</h4> <!>',1),Ia=B('<section class="p-4 flex flex-col gap-4"><label class="label"><span>Quantity</span> <input min="1" class="input p-1" type="number" placeholder="Input"></label> <label class="label"><span>Per Unit Amount</span> <input class="input p-1" min="0" type="number" placeholder="Input"></label> <label class="label flex flex-col gap-2 w-full"><span>Notes</span> <textarea class="textarea p-1 w-full" rows="4"></textarea></label></section> <footer class="card-footer flex justify-end"><button class="btn variant-filled">save</button></footer>',1),$a=B("<div><!></div>");function Sa(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$store",g),s=be();let e=V("pick_product"),r=V(""),_=V(1),l=V(0),h=V(void 0),w=V(0),b=V(ve(d()[0].meta.data||[]));(async()=>{const m=d()[0].meta.api;if(!m)return;const C=d()[0].meta.pid,I=await m.listProducts(C);I.status===200&&v(b,ve(I.data))})(),console.log("@parnet",t.parent);const i=async()=>{const m=d()[0].meta.onSave;m&&m({product_id:a(h),qty:a(_),amount:a(l),price:a(w),info:a(r),total_amount:a(l)*a(_),discount_amount:a(w)-a(l),tax_amount:0}),s.close()};var c=$a(),P=n(c);q(P,()=>d()[0],m=>{var C=de(),I=Z(C);q(I,()=>a(e)==="pick_product",z=>{var D=Pa(),G=y(Z(D),2),E=lt(()=>[{Name:"select",icon:"plus",Action:async(U,N)=>{v(h,ve(U)),v(l,ve(N.price||0)),v(w,ve(N.price||0)),v(r,(N.name||"")+"  "+(N.variant_id||"")),v(e,"set_quantity")}}]);ct(G,{action_key:"id",key_names:[["id","ID"],["name","Name"],["info","Info"],["price","Price"],["variant_id","Variant"]],get datas(){return a(b)},color:["variant_id"],get actions(){return a(E)}}),k(z,D)},z=>{var D=de(),G=Z(D);q(G,()=>a(e)==="loading",E=>{Xe(E,{})},E=>{var U=Ia(),N=Z(U),K=n(N),F=y(n(K),2);We(F),o(K);var ee=y(K,2),H=y(n(ee),2);We(H),o(ee);var ie=y(ee,2),W=y(n(ie),2);Kt(W),Y(W,"placeholder","information about account"),o(ie),o(N);var fe=y(N,2),De=n(fe);De.__click=i,o(fe),M(()=>Y(H,"max",a(w))),Oe(F,()=>a(_),ge=>v(_,ge)),Oe(H,()=>a(l),ge=>v(l,ge)),Oe(W,()=>a(r),ge=>v(r,ge)),k(E,U)},!0),k(z,D)}),k(m,C)}),o(c),M(()=>L(c,`card rounded p-4 w-modal-wide min-h-96 ${t.parent.backgroundColor??""}`)),k(u,c),ne()}ze(["click"]);var Da=B('<div><h4 class="h4">Pick Contact</h4> <!></div>');function Ca(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$store",g),s=be();let e=V(void 0),r=V(ve(d()[0].meta.data||[]));(async()=>{if(a(r).length)return;const w=d()[0].meta.api;if(!w)return;v(e,"loading");const b=d()[0].meta.pid,T=await w.listContacts(b);T.status===200&&(v(r,ve(T.data)),v(e,"main"))})();var l=Da(),h=y(n(l),2);q(h,()=>a(e)==="loading",w=>{Xe(w,{})},w=>{var b=lt(()=>[{Name:"select",icon:"plus",Action:async(T,i)=>{const c=d()[0].meta.onSelect;c==null||c(i),s.close()}}]);ct(w,{action_key:"id",key_names:[["id","ID"],["name","Name"],["ctype","Type"],["email","Email"],["phone","Phone"]],get datas(){return a(r)},color:["ctype"],get actions(){return a(b)}})}),o(l),M(()=>L(l,`card rounded p-4 w-modal-wide min-h-96 ${t.parent.backgroundColor??""}`)),k(u,l),ne()}var ja=B('<div class="w-full w-modal max-w-md bg-white rounded"><div class="text-center"><h2 class="text-xl font-bold text-gray-900">New Folder!</h2></div> <div class="p-4 flex flex-col gap-2"><label class="text-sm font-bold text-gray-500 tracking-wide">Name <input class="text-base p-2 border border-gray-300 rounded-lg focus:outline-none focus:border-indigo-500 w-full" type="text" placeholder="new_folder"></label> <div class="inline-flex justify-end"><button class="btn variant-filled-primary">Create</button></div></div></div>');function Na(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$store",g),s=be();let e=V("");const r=d()[0].meta.onNewName;var _=ja(),l=y(n(_),2),h=n(l),w=y(n(h));We(w),o(h);var b=y(h,2),T=n(b);T.__click=()=>{r(a(e)),s.close()},o(b),o(l),o(_),M(()=>T.disabled=a(e)===""),Oe(w,()=>a(e),i=>v(e,i)),k(u,_),ne()}ze(["click"]);var za=B('<div class="flex flex-row justify-between items-center hover:bg-gray-100 rounded p-1"><div class="flex flex-col"><h5 class="h-5 font-semibold"> </h5> <p class="p"> </p></div> <div class="flex flex-row"><a class="btn btn-sm variant-filled">open</a></div></div>'),Ba=B('<div class="w-modal-slim bg-white rounded card p-4"><h3 class="h3">Plugins Sources</h3> <section class="flex flex-col gap-2"></section></div>');function Ea(u,t){re(t,!1);const g=ue(),d=()=>oe(s,"$store",g),s=be(),r=d()[0].meta.pid,_=[{name:"URL",description:"Import plugin from a URL",link:`/z/pages/portal/project/plugins/new/URL?pid=${r}`},{name:"File",description:"Upload a plugin file",link:`/z/pages/portal/project/plugins/new/File?pid=${r}`},{name:"Live Editor",description:"Create a plugin in the editor",link:`/z/pages/portal/project/plugins/new/LiveEditor?pid=${r}`},{name:"Raw Create",description:"Create a plugin from scratch",link:`/z/pages/portal/project/plugins/new?pid=${r}`}];Qe();var l=Ba(),h=y(n(l),2);dt(h,5,()=>_,ut,(w,b)=>{var T=za(),i=n(T),c=n(i),P=n(c);o(c);var m=y(c,2),C=n(m);o(m),o(i);var I=y(i,2),z=n(I);z.__click=()=>{s.close()},o(I),o(T),M(()=>{me(P,a(b).name),me(C,a(b).description),Y(z,"href",a(b).link)}),k(w,T)}),o(h),o(l),k(u,l),ne()}ze(["click"]);var Aa=B('<div class="card p-2 shadow"><div class="flex justify-between"><div class="flex flex-col gap-1"><h6 class="h6"> </h6> <p class="text-sm text-gray-500"> </p></div></div></div>'),La=B('<div class="w-full h-full relative"><button class="absolute top-2 right-6 h-4 w-4 rounded border"><!></button> <h5 class="h5 p-2">Notifications</h5> <div class="flex flex-col gap-2 border-t border-gray-200 p-2 max-h-[90%] overflow-scroll"></div> <div class="flex justify-end p-2 gap-2"><a href="/z/pages/portal/self/inbox" class="underline inline-flex items-center gap-1"><!> inbox</a></div></div>');function Ma(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$drawer",g),s=ht(),e=d().meta.getRootAPI;let r=V(ve([]));(async()=>{const m=await e().listUserMessages();m.status===200&&v(r,ve(m.data))})();var l=La(),h=n(l);h.__click=()=>{s.close()};var w=n(h);pe(w,{name:"x-mark",className:"h-4 w-4"}),o(h);var b=y(h,4);dt(b,21,()=>a(r),ut,(P,m)=>{var C=Aa(),I=n(C),z=n(I),D=n(z),G=n(D);o(D);var E=y(D,2),U=n(E);o(E),o(z),o(I),o(C),M(()=>{me(G,a(m).title||""),me(U,a(m).contents||"")}),k(P,C)}),o(b);var T=y(b,2),i=n(T);i.__click=()=>{s.close()};var c=n(i);pe(c,{className:"w-4 h-4",name:"inbox"}),Pe(),o(i),o(T),o(l),k(u,l),ne()}ze(["click"]);var Oa=B("<div>TODO</div>");function Fa(u,t){re(t,!0),be(),p(t,"sub_total",3,0),p(t,"overall_tax_amount",3,0),p(t,"onSet",3,d=>{});var g=Oa();M(()=>L(g,`card rounded w-modal-slim ${t.parent.backgroundColor??""}`)),k(u,g),ne()}var Ha=B('<div><section class="p-4 flex flex-col gap-4"><label class="label"><span>Overall Discounted Amount</span> <input class="input p-1" type="number" placeholder="Input"></label></section> <footer class="card-footer flex justify-between"><div class="text-xs"><span>Overall Discount Percentage</span> <span> </span></div> <button class="btn variant-filled">set</button></footer></div>');function Ra(u,t){re(t,!0);const g=ue(),d=()=>oe(s,"$store",g),s=be(),e=d()[0].meta||{};let r=e.sub_total||0,_=e.overall_discount_amount||0,l=e.onSet||(()=>{}),h=V(r-_),w=V(0);Qt(()=>{v(w,(r-a(h))/r*100)});var b=Ha(),T=n(b),i=n(T),c=y(n(i),2);We(c),o(i),o(T);var P=y(T,2),m=n(P),C=y(n(m),2),I=n(C);M(()=>me(I,`${Vt(a(w))??""}%`)),o(C),o(m);var z=y(m,2);z.__click=()=>{l(r-a(h)),s.close()},o(P),o(b),M(()=>L(b,`card rounded w-modal-slim ${t.parent.backgroundColor??""}`)),Oe(c,()=>a(h),D=>v(h,D)),k(u,b),ne()}ze(["click"]);var qa=B('<div class="w-fit min-w-96 bg-white p-4 max-h-[90vh] overflow-auto"><!></div>');function Ua(u,t){re(t,!1);const g=ue(),d=()=>oe(s,"$store",g),s=be(),e=d()[0].meta;Qe();var r=qa(),_=n(r);Yt(_,{get api(){return e.api},get fileId(){return e.fileId},get filename(){return e.filename}}),o(r),k(u,r),ne()}var Ka=B("<!> <!> <!>",1);function jo(u,t){re(t,!0),na();const g={project_picker:{ref:xa},project_options:{ref:ka},project_new_plugin_picker:{ref:Ea},new_folder_panel:{ref:Na},books_account_picker:{ref:Ta},books_sales_item_pick:{ref:Sa},books_client_picker:{ref:Ca},books_sales_overall_tax_picker:{ref:Fa},books_sales_overall_discount_picker:{ref:Ra},file_picker:{ref:Wt},file_preview_dialog:{ref:Ua}};var d=Ka(),s=Z(d);_a(s,{children:(_,l)=>{Ma(_,{})},$$slots:{default:!0}});var e=y(s,2);ba(e,{components:g});var r=y(e,2);Lt(r,()=>t.children??At),k(u,d),ne()}export{jo as component,Co as universal};
