import{a as E,t as R}from"../chunks/disclose-version.wcsHG7oF.js";import{p as O,g as Q,i as l,a as T,h as p,s as t,c as r,b as e,r as s,e as U,d as W,t as Y}from"../chunks/runtime.D6lgFUio.js";import{s as Z}from"../chunks/render.CIW-w3Th.js";import{e as aa,i as ea}from"../chunks/each.CWrzZiCN.js";import{s as ta,r as m}from"../chunks/attributes.BFGPVQjG.js";import{f as sa}from"../chunks/utils.Dh-XvBAd.js";import{r as ra}from"../chunks/misc.CFrgioPu.js";import{b as n}from"../chunks/input.DAoTAKLA.js";import{b as oa}from"../chunks/select.BfGKzECA.js";import{p as la}from"../chunks/event-modifiers.Bfc47y5P.js";import{p as G}from"../chunks/proxy.r5-P5EHK.js";import{s as pa,a as ia}from"../chunks/store.BaapSzAp.js";import{N as na}from"../chunks/index.Duto3aX0.js";import{p as ca}from"../chunks/stores.D4Dcgzbh.js";import"../chunks/auto_form.CxbPX2mo.js";import"../chunks/autotable.D6q4cofd.js";import"../chunks/index.BXdKDRLQ.js";/* empty css                                                     */import{g as ma}from"../chunks/entry.Y6HcnCx3.js";import{S as ua}from"../chunks/SvgIcon.Cj01WnO8.js";var da=R("<option> </option>"),va=R('<form class="p-2"><div class="card"><header class="card-header"><h4 class="h4">Add Product</h4></header> <section class="p-4 flex flex-col gap-4"><label class="label"><span>Name</span> <input class="input p-1" type="text" placeholder="Input"></label> <label class="label"><span>Catagory</span> <select class="select" required></select></label> <label class="label"><span>Variant</span> <input type="text" name="variant_id" class="input p-1" placeholder="LARGE XL"></label> <label class="label" for="price"><span>Price</span></label> <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]"><span class="w-8"><!></span> <input id="price" class="input p-1" type="number" placeholder="Input"></div> <label class="label"><span>In Stock</span> <input id="stock_count" class="input p-1" type="number" placeholder="Input"></label> <label class="label"><span>Info</span> <textarea class="textarea p-1" rows="4"></textarea></label></section> <footer class="card-footer flex justify-end"><button type="submit" class="btn variant-filled">save</button></footer></div></form>');function Ea(V,X){O(X,!0);const F=pa(),u=ia(ca,"$page",F).params.pid,C=na(Q("__api__"));let d=l(""),v=l(""),f=l(0),b=l(0),_=l(""),g=l(0),P=l(G([]));(async()=>{const a=await C.listCatagories(u);a.status===200&&p(P,G(a.data))})();const H=async()=>{const a=await C.addProduct(u,{name:e(d),info:e(v),catagory_id:Number(e(f)),price:e(b),variant_id:e(_),stock_count:e(g)});if(a.status!==200){a.data.message;return}ma(`/z/pages/portal/projects/books/${u}/inventory/products`)};var c=va(),J=W(()=>la(H)),j=r(c),A=t(r(j),2),h=r(A),S=t(r(h),2);m(S),s(h);var x=t(h,2),y=t(r(x),2);aa(y,21,()=>e(P),ea,(a,o)=>{var i=da(),D={},M=r(i);s(i),Y(()=>{D!==(D=e(o).id)&&(i.value=(i.__value=e(o).id)==null?"":e(o).id),Z(M,e(o).name)}),E(a,i)}),s(y),s(x);var I=t(x,2),L=t(r(I),2);m(L),s(I);var k=t(I,4),w=r(k),K=r(w);ua(K,{name:"currency-dollar",className:"w-5 h-5 text-gray-700 inline-flex justify-self-center items-center mt-1 ml-1"}),s(w);var q=t(w,2);m(q),s(k);var $=t(k,2),z=t(r($),2);m(z),s($);var B=t($,2),N=t(r(B),2);ra(N),ta(N,"placeholder","information about account"),s(B),s(A),U(2),s(j),s(c),sa("submit",c,function(...a){var o;(o=e(J))==null||o.apply(this,a)}),n(S,()=>e(d),a=>p(d,a)),oa(y,()=>e(f),a=>p(f,a)),n(L,()=>e(_),a=>p(_,a)),n(q,()=>e(b),a=>p(b,a)),n(z,()=>e(g),a=>p(g,a)),n(N,()=>e(v),a=>p(v,a)),E(V,c),T()}export{Ea as component};
