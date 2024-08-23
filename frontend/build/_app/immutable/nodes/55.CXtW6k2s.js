import{s as Ce,a as M,e as Q,g as Z,c as w,b as S,f as b,m as v,i as y,h as L,x as H,k as Se,r as Pe,v as re,n as Y,t as oe,d as le,A as Te,R as X,T as z}from"../chunks/scheduler.D2OtNeTK.js";import{S as Ie,i as Ue,c as P,b as T,m as I,t as $,a as C,e as U,f as V,g as Be,d as qe}from"../chunks/index.DD9KLV65.js";import"../chunks/ProgressBar.svelte_svelte_type_style_lang.B6Y8wNrO.js";import{A as De}from"../chunks/AppBar.C3Ur4xTI.js";import{T as Le,a as ie,h as Re}from"../chunks/index.CxVLFxMs.js";import{L as Ee,b as je,E as Xe,i as Ne,c as Ae,f as Me,d as Ze,t as h,e as ze,g as Ve,h as We,k as Ye,C as ve}from"../chunks/index.BTX9jO80.js";import{p as Fe}from"../chunks/stores.D0bQ952I.js";import{N as Ge}from"../chunks/index.ZlEQwogg.js";import{S as ce}from"../chunks/SvgIcon.DW0uSnC-.js";const Ke=36,de=1,He=2,j=3,G=4,Je=5,et=6,tt=7,at=8,rt=9,st=10,nt=11,ot=12,lt=13,it=14,ct=15,dt=16,ft=17,fe=18,mt=19,ke=20,_e=21,me=22,ut=23,pt=24;function J(t){return t>=65&&t<=90||t>=97&&t<=122||t>=48&&t<=57}function ht(t){return t>=48&&t<=57||t>=97&&t<=102||t>=65&&t<=70}function D(t,e,r){for(let a=!1;;){if(t.next<0)return;if(t.next==e&&!a){t.advance();return}a=r&&!a&&t.next==92,t.advance()}}function gt(t,e){e:for(;;){if(t.next<0)return console.log("exit at end",t.pos);if(t.next==36){t.advance();for(let r=0;r<e.length;r++){if(t.next!=e.charCodeAt(r))continue e;t.advance()}if(t.next==36){t.advance();return}}else t.advance()}}function Ot(t,e){let r="[{<(".indexOf(String.fromCharCode(e)),a=r<0?e:"]}>)".charCodeAt(r);for(;;){if(t.next<0)return;if(t.next==a&&t.peek(1)==39){t.advance(2);return}t.advance()}}function ee(t,e){for(;!(t.next!=95&&!J(t.next));)e!=null&&(e+=String.fromCharCode(t.next)),t.advance();return e}function bt(t){if(t.next==39||t.next==34||t.next==96){let e=t.next;t.advance(),D(t,e,!1)}else ee(t)}function ue(t,e){for(;t.next==48||t.next==49;)t.advance();e&&t.next==e&&t.advance()}function pe(t,e){for(;;){if(t.next==46){if(e)break;e=!0}else if(t.next<48||t.next>57)break;t.advance()}if(t.next==69||t.next==101)for(t.advance(),(t.next==43||t.next==45)&&t.advance();t.next>=48&&t.next<=57;)t.advance()}function he(t){for(;!(t.next<0||t.next==10);)t.advance()}function q(t,e){for(let r=0;r<e.length;r++)if(e.charCodeAt(r)==t)return!0;return!1}const K=` 	\r
`;function xe(t,e,r){let a=Object.create(null);a.true=a.false=Je,a.null=a.unknown=et;for(let s of t.split(" "))s&&(a[s]=ke);for(let s of e.split(" "))s&&(a[s]=_e);for(let s of(r||"").split(" "))s&&(a[s]=pt);return a}const Qe="array binary bit boolean char character clob date decimal double float int integer interval large national nchar nclob numeric object precision real smallint time timestamp varchar varying ",we="absolute action add after all allocate alter and any are as asc assertion at authorization before begin between both breadth by call cascade cascaded case cast catalog check close collate collation column commit condition connect connection constraint constraints constructor continue corresponding count create cross cube current current_date current_default_transform_group current_transform_group_for_type current_path current_role current_time current_timestamp current_user cursor cycle data day deallocate declare default deferrable deferred delete depth deref desc describe descriptor deterministic diagnostics disconnect distinct do domain drop dynamic each else elseif end end-exec equals escape except exception exec execute exists exit external fetch first for foreign found from free full function general get global go goto grant group grouping handle having hold hour identity if immediate in indicator initially inner inout input insert intersect into is isolation join key language last lateral leading leave left level like limit local localtime localtimestamp locator loop map match method minute modifies module month names natural nesting new next no none not of old on only open option or order ordinality out outer output overlaps pad parameter partial path prepare preserve primary prior privileges procedure public read reads recursive redo ref references referencing relative release repeat resignal restrict result return returns revoke right role rollback rollup routine row rows savepoint schema scroll search second section select session session_user set sets signal similar size some space specific specifictype sql sqlexception sqlstate sqlwarning start state static system_user table temporary then timezone_hour timezone_minute to trailing transaction translation treat trigger under undo union unique unnest until update usage user using value values view when whenever where while with without work write year zone ",te={backslashEscapes:!1,hashComments:!1,spaceAfterDashes:!1,slashComments:!1,doubleQuotedStrings:!1,doubleDollarQuotedStrings:!1,unquotedBitLiterals:!1,treatBitsAsBytes:!1,charSetCasts:!1,plsqlQuotingMechanism:!1,operatorChars:"*+-%<>!=&|~^/",specialVar:"?",identifierQuotes:'"',caseInsensitiveIdentifiers:!1,words:xe(we,Qe)};function vt(t,e,r,a){let s={};for(let l in te)s[l]=(t.hasOwnProperty(l)?t:te)[l];return e&&(s.words=xe(e,r||"",a)),s}function ye(t){return new Xe(e=>{var r;let{next:a}=e;if(e.advance(),q(a,K)){for(;q(e.next,K);)e.advance();e.acceptToken(Ke)}else if(a==36&&t.doubleDollarQuotedStrings){let s=ee(e,"");e.next==36&&(e.advance(),gt(e,s),e.acceptToken(j))}else if(a==39||a==34&&t.doubleQuotedStrings)D(e,a,t.backslashEscapes),e.acceptToken(j);else if(a==35&&t.hashComments||a==47&&e.next==47&&t.slashComments)he(e),e.acceptToken(de);else if(a==45&&e.next==45&&(!t.spaceAfterDashes||e.peek(1)==32))he(e),e.acceptToken(de);else if(a==47&&e.next==42){e.advance();for(let s=1;;){let l=e.next;if(e.next<0)break;if(e.advance(),l==42&&e.next==47){if(s--,e.advance(),!s)break}else l==47&&e.next==42&&(s++,e.advance())}e.acceptToken(He)}else if((a==101||a==69)&&e.next==39)e.advance(),D(e,39,!0),e.acceptToken(j);else if((a==110||a==78)&&e.next==39&&t.charSetCasts)e.advance(),D(e,39,t.backslashEscapes),e.acceptToken(j);else if(a==95&&t.charSetCasts)for(let s=0;;s++){if(e.next==39&&s>1){e.advance(),D(e,39,t.backslashEscapes),e.acceptToken(j);break}if(!J(e.next))break;e.advance()}else if(t.plsqlQuotingMechanism&&(a==113||a==81)&&e.next==39&&e.peek(1)>0&&!q(e.peek(1),K)){let s=e.peek(1);e.advance(2),Ot(e,s),e.acceptToken(j)}else if(a==40)e.acceptToken(tt);else if(a==41)e.acceptToken(at);else if(a==123)e.acceptToken(rt);else if(a==125)e.acceptToken(st);else if(a==91)e.acceptToken(nt);else if(a==93)e.acceptToken(ot);else if(a==59)e.acceptToken(lt);else if(t.unquotedBitLiterals&&a==48&&e.next==98)e.advance(),ue(e),e.acceptToken(me);else if((a==98||a==66)&&(e.next==39||e.next==34)){const s=e.next;e.advance(),t.treatBitsAsBytes?(D(e,s,t.backslashEscapes),e.acceptToken(ut)):(ue(e,s),e.acceptToken(me))}else if(a==48&&(e.next==120||e.next==88)||(a==120||a==88)&&e.next==39){let s=e.next==39;for(e.advance();ht(e.next);)e.advance();s&&e.next==39&&e.advance(),e.acceptToken(G)}else if(a==46&&e.next>=48&&e.next<=57)pe(e,!0),e.acceptToken(G);else if(a==46)e.acceptToken(it);else if(a>=48&&a<=57)pe(e,!1),e.acceptToken(G);else if(q(a,t.operatorChars)){for(;q(e.next,t.operatorChars);)e.advance();e.acceptToken(ct)}else if(q(a,t.specialVar))e.next==a&&e.advance(),bt(e),e.acceptToken(ft);else if(q(a,t.identifierQuotes))D(e,a,!1),e.acceptToken(mt);else if(a==58||a==44)e.acceptToken(dt);else if(J(a)){let s=ee(e,String.fromCharCode(a));e.acceptToken(e.next==46||e.peek(-s.length-1)==46?fe:(r=t.words[s.toLowerCase()])!==null&&r!==void 0?r:fe)}})}const $e=ye(te),kt=ze.deserialize({version:14,states:"%vQ]QQOOO#wQRO'#DSO$OQQO'#CwO%eQQO'#CxO%lQQO'#CyO%sQQO'#CzOOQQ'#DS'#DSOOQQ'#C}'#C}O'UQRO'#C{OOQQ'#Cv'#CvOOQQ'#C|'#C|Q]QQOOQOQQOOO'`QQO'#DOO(xQRO,59cO)PQQO,59cO)UQQO'#DSOOQQ,59d,59dO)cQQO,59dOOQQ,59e,59eO)jQQO,59eOOQQ,59f,59fO)qQQO,59fOOQQ-E6{-E6{OOQQ,59b,59bOOQQ-E6z-E6zOOQQ,59j,59jOOQQ-E6|-E6|O+VQRO1G.}O+^QQO,59cOOQQ1G/O1G/OOOQQ1G/P1G/POOQQ1G/Q1G/QP+kQQO'#C}O+rQQO1G.}O)PQQO,59cO,PQQO'#Cw",stateData:",[~OtOSPOSQOS~ORUOSUOTUOUUOVROXSOZTO]XO^QO_UO`UOaPObPOcPOdUOeUOfUOgUOhUO~O^]ORvXSvXTvXUvXVvXXvXZvX]vX_vX`vXavXbvXcvXdvXevXfvXgvXhvX~OsvX~P!jOa_Ob_Oc_O~ORUOSUOTUOUUOVROXSOZTO^tO_UO`UOa`Ob`Oc`OdUOeUOfUOgUOhUO~OWaO~P$ZOYcO~P$ZO[eO~P$ZORUOSUOTUOUUOVROXSOZTO^QO_UO`UOaPObPOcPOdUOeUOfUOgUOhUO~O]hOsoX~P%zOajObjOcjO~O^]ORkaSkaTkaUkaVkaXkaZka]ka_ka`kaakabkackadkaekafkagkahka~Oska~P'kO^]O~OWvXYvX[vX~P!jOWnO~P$ZOYoO~P$ZO[pO~P$ZO^]ORkiSkiTkiUkiVkiXkiZki]ki_ki`kiakibkickidkiekifkigkihki~Oski~P)xOWkaYka[ka~P'kO]hO~P$ZOWkiYki[ki~P)xOasObsOcsO~O",goto:"#hwPPPPPPPPPPPPPPPPPPPPPPPPPPx||||!Y!^!d!xPPP#[TYOZeUORSTWZbdfqT[OZQZORiZSWOZQbRQdSQfTZgWbdfqQ^PWk^lmrQl_Qm`RrseVORSTWZbdfq",nodeNames:"⚠ LineComment BlockComment String Number Bool Null ( ) { } [ ] ; . Operator Punctuation SpecialVar Identifier QuotedIdentifier Keyword Type Bits Bytes Builtin Script Statement CompositeIdentifier Parens Braces Brackets Statement",maxTerm:38,nodeProps:[["isolate",-4,1,2,3,19,""]],skippedNodes:[0,1,2],repeatNodeCount:3,tokenData:"RORO",tokenizers:[0,$e],topRules:{Script:[0,25]},tokenPrec:0});function ae(t){let e=t.cursor().moveTo(t.from,-1);for(;/Comment/.test(e.name);)e.moveTo(e.from,-1);return e.node}function N(t,e){let r=t.sliceString(e.from,e.to),a=/^([`'"])(.*)\1$/.exec(r);return a?a[2]:r}function W(t){return t&&(t.name=="Identifier"||t.name=="QuotedIdentifier")}function _t(t,e){if(e.name=="CompositeIdentifier"){let r=[];for(let a=e.firstChild;a;a=a.nextSibling)W(a)&&r.push(N(t,a));return r}return[N(t,e)]}function ge(t,e){for(let r=[];;){if(!e||e.name!=".")return r;let a=ae(e);if(!W(a))return r;r.unshift(N(t,a)),e=ae(a)}}function xt(t,e){let r=Ye(t).resolveInner(e,-1),a=wt(t.doc,r);return r.name=="Identifier"||r.name=="QuotedIdentifier"||r.name=="Keyword"?{from:r.from,quoted:r.name=="QuotedIdentifier"?t.doc.sliceString(r.from,r.from+1):null,parents:ge(t.doc,ae(r)),aliases:a}:r.name=="."?{from:e,quoted:null,parents:ge(t.doc,r),aliases:a}:{from:e,quoted:null,parents:[],empty:!0,aliases:a}}const Qt=new Set("where group having order union intersect except all distinct limit offset fetch for".split(" "));function wt(t,e){let r;for(let s=e;!r;s=s.parent){if(!s)return null;s.name=="Statement"&&(r=s)}let a=null;for(let s=r.firstChild,l=!1,n=null;s;s=s.nextSibling){let c=s.name=="Keyword"?t.sliceString(s.from,s.to).toLowerCase():null,i=null;if(!l)l=c=="from";else if(c=="as"&&n&&W(s.nextSibling))i=N(t,s.nextSibling);else{if(c&&Qt.has(c))break;n&&W(s)&&(i=N(t,s))}i&&(a||(a=Object.create(null)),a[i]=_t(t,n)),n=/Identifier$/.test(s.name)?s:null}return a}function yt(t,e){return t?e.map(r=>Object.assign(Object.assign({},r),{label:r.label[0]==t?r.label:t+r.label+t,apply:void 0})):e}const $t=/^\w*$/,Ct=/^[`'"]?\w*[`'"]?$/;function Oe(t){return t.self&&typeof t.self.label=="string"}class se{constructor(e,r){this.idQuote=e,this.idCaseInsensitive=r,this.list=[],this.children=void 0}child(e){let r=this.children||(this.children=Object.create(null)),a=r[e];return a||(e&&!this.list.some(s=>s.label==e)&&this.list.push(be(e,"type",this.idQuote,this.idCaseInsensitive)),r[e]=new se(this.idQuote,this.idCaseInsensitive))}maybeChild(e){return this.children?this.children[e]:null}addCompletion(e){let r=this.list.findIndex(a=>a.label==e.label);r>-1?this.list[r]=e:this.list.push(e)}addCompletions(e){for(let r of e)this.addCompletion(typeof r=="string"?be(r,"property",this.idQuote,this.idCaseInsensitive):r)}addNamespace(e){Array.isArray(e)?this.addCompletions(e):Oe(e)?this.addNamespace(e.children):this.addNamespaceObject(e)}addNamespaceObject(e){for(let r of Object.keys(e)){let a=e[r],s=null,l=r.replace(/\\?\./g,c=>c=="."?"\0":c).split("\0"),n=this;Oe(a)&&(s=a.self,a=a.children);for(let c=0;c<l.length;c++)s&&c==l.length-1&&n.addCompletion(s),n=n.child(l[c].replace(/\\\./g,"."));n.addNamespace(a)}}}function be(t,e,r,a){return new RegExp("^[a-z_][a-z_\\d]*$",a?"i":"").test(t)?{label:t,type:e}:{label:t,type:e,apply:r+t+r}}function St(t,e,r,a,s,l){var n;let c=((n=l==null?void 0:l.spec.identifierQuotes)===null||n===void 0?void 0:n[0])||'"',i=new se(c,!!(l!=null&&l.spec.caseInsensitiveIdentifiers)),d=s?i.child(s):null;return i.addNamespace(t),e&&(d||i).addCompletions(e),r&&i.addCompletions(r),d&&i.addCompletions(d.list),a&&i.addCompletions((d||i).child(a).list),m=>{let{parents:o,from:u,quoted:f,empty:p,aliases:_}=xt(m.state,m.pos);if(p&&!m.explicit)return null;_&&o.length==1&&(o=_[o[0]]||o);let g=i;for(let E of o){for(;!g.children||!g.children[E];)if(g==i&&d)g=d;else if(g==d&&a)g=g.child(a);else return null;let A=g.maybeChild(E);if(!A)return null;g=A}let R=f&&m.state.sliceDoc(m.pos,m.pos+1)==f,B=g.list;return g==i&&_&&(B=B.concat(Object.keys(_).map(E=>({label:E,type:"constant"})))),{from:u,to:R?m.pos+1:void 0,options:yt(f,B),validFor:f?Ct:$t}}}function Pt(t,e){let r=Object.keys(t).map(a=>({label:e?a.toUpperCase():a,type:t[a]==_e?"type":t[a]==ke?"keyword":"variable",boost:-1}));return Ve(["QuotedIdentifier","SpecialVar","String","LineComment","BlockComment","."],We(r))}let Tt=kt.configure({props:[Ne.add({Statement:Ae()}),Me.add({Statement(t,e){return{from:Math.min(t.from+100,e.doc.lineAt(t.from).to),to:t.to}},BlockComment(t){return{from:t.from+2,to:t.to-2}}}),Ze({Keyword:h.keyword,Type:h.typeName,Builtin:h.standard(h.name),Bits:h.number,Bytes:h.string,Bool:h.bool,Null:h.null,Number:h.number,String:h.string,Identifier:h.name,QuotedIdentifier:h.special(h.string),SpecialVar:h.special(h.name),LineComment:h.lineComment,BlockComment:h.blockComment,Operator:h.operator,"Semi Punctuation":h.punctuation,"( )":h.paren,"{ }":h.brace,"[ ]":h.squareBracket})]});class F{constructor(e,r,a){this.dialect=e,this.language=r,this.spec=a}get extension(){return this.language.extension}static define(e){let r=vt(e,e.keywords,e.types,e.builtin),a=je.define({name:"sql",parser:Tt.configure({tokenizers:[{from:$e,to:ye(r)}]}),languageData:{commentTokens:{line:"--",block:{open:"/*",close:"*/"}},closeBrackets:{brackets:["(","[","{","'",'"',"`"]}}});return new F(r,a,e)}}function It(t,e=!1){return Pt(t.dialect.words,e)}function Ut(t,e=!1){return t.language.data.of({autocomplete:It(t,e)})}function Bt(t){return t.schema?St(t.schema,t.tables,t.schemas,t.defaultTable,t.defaultSchema,t.dialect||ne):()=>null}function qt(t){return t.schema?(t.dialect||ne).language.data.of({autocomplete:Bt(t)}):[]}function Dt(t={}){let e=t.dialect||ne;return new Ee(e.language,[qt(t),Ut(e,!!t.upperCaseKeywords)])}const ne=F.define({}),Lt=F.define({keywords:we+"abort analyze attach autoincrement conflict database detach exclusive fail glob ignore index indexed instead isnull notnull offset plan pragma query raise regexp reindex rename replace temp vacuum virtual",types:Qe+"bool blob long longblob longtext medium mediumblob mediumint mediumtext tinyblob tinyint tinytext text bigint int2 int8 unsigned signed real",builtin:"auth backup bail changes clone databases dbinfo dump echo eqp explain fullschema headers help import imposter indexes iotrace lint load log mode nullvalue once print prompt quit restore save scanstats separator shell show stats system tables testcase timeout timer trace vfsinfo vfslist vfsname width",operatorChars:"*+-%<>!=&|/~",identifierQuotes:'`"',specialVar:"@:?$"});let Rt=`
    
-- query_name: list_accounts
select * from __project__Accounts;

































`;const Et=`

<!DOCTYPE html>
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
        <h1>Reports Editor Example</h1> 
        <p>This is a sample preview</p>
        
    </div>
    <div class="container-xl mt-5" x-data="accountTable()" id="ac_table">
        
    <div class="container">
        <button id="start" class="btn btn-primary" type="button" x-on:click="fetchAccounts">
            <span class="icon icon-thumbs-up"></span>
            Start
        </button>    
    </div>


        <div x-show="!loading" class="table-responsive">
            <table class="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Type</th>
                        <th>Info</th>
                        <th>Created At</th>
                        <th>Updated At</th>
                    </tr>
                </thead>
                <tbody>
                    <template x-for="account in accounts" :key="account.id">
                        <tr>
                            <td x-text="account.id"></td>
                            <td x-text="account.name"></td>
                            <td x-text="account.acc_type" class="text-capitalize"></td>
                            <td x-text="account.info"></td>
                            <td x-text="formatDate(account.created_at)"></td>
                            <td x-text="formatDate(account.updated_at)"></td>
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
    </div>

    <script>
    function accountTable() {
      return {
          accounts: [],
          loading: true,
          async fetchAccounts() {
              this.loading = true;
              try {
                  const data = await sendMessage({
                      type: "sql_query",
                      name: "list_accounts",
                      data: [],
                  });
                  console.log("accounts", data);

                  this.accounts = data.data;
              } catch (error) {
                  console.error("Error fetching accounts:", error);
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
</html>







`;function jt(t){let e,r='<li class="crumb"><a class="anchor" href="/z/pages/portal/projects/books">Books</a></li> <li class="crumb-separator" aria-hidden="">›</li> <li>Reports</li> <li class="crumb-separator" aria-hidden="">›</li> <li>Editor</li>';return{c(){e=Q("ol"),e.innerHTML=r,this.h()},l(a){e=w(a,"OL",{class:!0,"data-svelte-h":!0}),re(e)!=="svelte-1enmns1"&&(e.innerHTML=r),this.h()},h(){v(e,"class","breadcrumb")},m(a,s){y(a,e,s)},p:Y,d(a){a&&b(e)}}}function Xt(t){let e,r,a,s,l,n,c,i,d,m;return r=new ce({props:{className:"w-4 h-4",name:"play"}}),n=new ce({props:{className:"w-4 h-4",name:"printer"}}),{c(){e=Q("button"),P(r.$$.fragment),a=oe(`
            Run`),s=M(),l=Q("button"),P(n.$$.fragment),c=oe(`
            Print`),this.h()},l(o){e=w(o,"BUTTON",{class:!0});var u=S(e);T(r.$$.fragment,u),a=le(u,`
            Run`),u.forEach(b),s=Z(o),l=w(o,"BUTTON",{class:!0});var f=S(l);T(n.$$.fragment,f),c=le(f,`
            Print`),f.forEach(b),this.h()},h(){v(e,"class","btn variant-filled btn-sm"),v(l,"class","btn variant-filled-secondary btn-sm")},m(o,u){y(o,e,u),I(r,e,null),L(e,a),y(o,s,u),y(o,l,u),I(n,l,null),L(l,c),i=!0,d||(m=[H(e,"click",t[7]),H(l,"click",t[8])],d=!0)},p:Y,i(o){i||($(r.$$.fragment,o),$(n.$$.fragment,o),i=!0)},o(o){C(r.$$.fragment,o),C(n.$$.fragment,o),i=!1},d(o){o&&(b(e),b(s),b(l)),U(r),U(n),d=!1,Te(m)}}}function Nt(t){let e,r="SQL Query";return{c(){e=Q("span"),e.textContent=r},l(a){e=w(a,"SPAN",{"data-svelte-h":!0}),re(e)!=="svelte-92glzm"&&(e.textContent=r)},m(a,s){y(a,e,s)},p:Y,d(a){a&&b(e)}}}function At(t){let e,r="UI";return{c(){e=Q("span"),e.textContent=r},l(a){e=w(a,"SPAN",{"data-svelte-h":!0}),re(e)!=="svelte-16okyjk"&&(e.textContent=r)},m(a,s){y(a,e,s)},p:Y,d(a){a&&b(e)}}}function Mt(t){let e,r,a,s,l,n;function c(o){t[11](o)}let i={padding:"ml-6 p-2",name:"SQL Query",value:0,$$slots:{default:[Nt]},$$scope:{ctx:t}};t[1]!==void 0&&(i.group=t[1]),e=new ie({props:i}),X.push(()=>V(e,"group",c));function d(o){t[12](o)}let m={name:"UI",value:1,$$slots:{default:[At]},$$scope:{ctx:t}};return t[1]!==void 0&&(m.group=t[1]),s=new ie({props:m}),X.push(()=>V(s,"group",d)),{c(){P(e.$$.fragment),a=M(),P(s.$$.fragment)},l(o){T(e.$$.fragment,o),a=Z(o),T(s.$$.fragment,o)},m(o,u){I(e,o,u),y(o,a,u),I(s,o,u),n=!0},p(o,u){const f={};u&524288&&(f.$$scope={dirty:u,ctx:o}),!r&&u&2&&(r=!0,f.group=o[1],z(()=>r=!1)),e.$set(f);const p={};u&524288&&(p.$$scope={dirty:u,ctx:o}),!l&&u&2&&(l=!0,p.group=o[1],z(()=>l=!1)),s.$set(p)},i(o){n||($(e.$$.fragment,o),$(s.$$.fragment,o),n=!0)},o(o){C(e.$$.fragment,o),C(s.$$.fragment,o),n=!1},d(o){o&&b(a),U(e,o),U(s,o)}}}function Zt(t){let e,r,a;function s(n){t[10](n)}let l={lang:Re()};return t[3]!==void 0&&(l.value=t[3]),e=new ve({props:l}),X.push(()=>V(e,"value",s)),{c(){P(e.$$.fragment)},l(n){T(e.$$.fragment,n)},m(n,c){I(e,n,c),a=!0},p(n,c){const i={};!r&&c&8&&(r=!0,i.value=n[3],z(()=>r=!1)),e.$set(i)},i(n){a||($(e.$$.fragment,n),a=!0)},o(n){C(e.$$.fragment,n),a=!1},d(n){U(e,n)}}}function zt(t){let e,r,a;function s(n){t[9](n)}let l={lang:Dt({dialect:Lt})};return t[2]!==void 0&&(l.value=t[2]),e=new ve({props:l}),X.push(()=>V(e,"value",s)),{c(){P(e.$$.fragment)},l(n){T(e.$$.fragment,n)},m(n,c){I(e,n,c),a=!0},p(n,c){const i={};!r&&c&4&&(r=!0,i.value=n[2],z(()=>r=!1)),e.$set(i)},i(n){a||($(e.$$.fragment,n),a=!0)},o(n){C(e.$$.fragment,n),a=!1},d(n){U(e,n)}}}function Vt(t){let e,r,a,s;const l=[zt,Zt],n=[];function c(i,d){return i[1]===0?0:i[1]===1?1:-1}return~(r=c(t))&&(a=n[r]=l[r](t)),{c(){e=Q("div"),a&&a.c(),this.h()},l(i){e=w(i,"DIV",{class:!0});var d=S(e);a&&a.l(d),d.forEach(b),this.h()},h(){v(e,"class","max-h-[40vh] md:max-h-[90vh] overflow-auto")},m(i,d){y(i,e,d),~r&&n[r].m(e,null),s=!0},p(i,d){let m=r;r=c(i),r===m?~r&&n[r].p(i,d):(a&&(Be(),C(n[m],1,1,()=>{n[m]=null}),qe()),~r?(a=n[r],a?a.p(i,d):(a=n[r]=l[r](i),a.c()),$(a,1),a.m(e,null)):a=null)},i(i){s||($(a),s=!0)},o(i){C(a),s=!1},d(i){i&&b(e),~r&&n[r].d()}}}function Wt(t){let e,r,a,s,l,n,c,i,d,m,o,u;return e=new De({props:{$$slots:{trail:[Xt],lead:[jt]},$$scope:{ctx:t}}}),l=new Le({props:{$$slots:{panel:[Vt],default:[Mt]},$$scope:{ctx:t}}}),{c(){P(e.$$.fragment),r=M(),a=Q("div"),s=Q("div"),P(l.$$.fragment),n=M(),c=Q("div"),i=Q("div"),d=Q("iframe"),this.h()},l(f){T(e.$$.fragment,f),r=Z(f),a=w(f,"DIV",{class:!0});var p=S(a);s=w(p,"DIV",{class:!0});var _=S(s);T(l.$$.fragment,_),_.forEach(b),n=Z(p),c=w(p,"DIV",{class:!0});var g=S(c);i=w(g,"DIV",{class:!0});var R=S(i);d=w(R,"IFRAME",{title:!0,srcdoc:!0,width:!0,height:!0,class:!0,allow:!0,sandbox:!0});var B=S(d);B.forEach(b),R.forEach(b),g.forEach(b),p.forEach(b),this.h()},h(){v(s,"class","flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full"),v(d,"title","preview"),v(d,"srcdoc",t[4]),v(d,"width","100%"),v(d,"height","100%"),v(d,"class","border-green-200 w-full h-full transition-all"),v(d,"allow","accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"),v(d,"sandbox","allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"),v(i,"class","card p-2 h-full w-full"),v(c,"class","flex-1 w-full md:w-1/2 border border-slate-50 h-1/2 md:h-full p-2"),v(a,"class","flex flex-col md:flex-row w-full h-[94vh]")},m(f,p){I(e,f,p),y(f,r,p),y(f,a,p),L(a,s),I(l,s,null),L(a,n),L(a,c),L(c,i),L(i,d),t[14](d),m=!0,o||(u=H(d,"load",t[13]),o=!0)},p(f,[p]){const _={};p&524313&&(_.$$scope={dirty:p,ctx:f}),e.$set(_);const g={};p&524302&&(g.$$scope={dirty:p,ctx:f}),l.$set(g),(!m||p&16)&&v(d,"srcdoc",f[4])},i(f){m||($(e.$$.fragment,f),$(l.$$.fragment,f),m=!0)},o(f){C(e.$$.fragment,f),C(l.$$.fragment,f),m=!1},d(f){f&&(b(r),b(a)),U(e,f),U(l),t[14](null),o=!1,u()}}}function Yt(t,e,r){let a;Se(t,Fe,O=>r(15,a=O));let s,l=0,n=Rt,c=Et,i="";const d=a.params.pid,m=Pe("__api__");Ge(m);let o;const u=async O=>{const k=O.data;if(console.log("onFrameMessage",k),k.type==="sql_query"){console.log("sql_query",k);try{const x=await m.runProjectSQL(d,{input:n,name:k.name,data:k.data});if(x.status!==200){o==null||o.postMessage({msgId:k.msgId,data:{msg:"error",data:x.data}});return}o==null||o.postMessage({msgId:k.msgId,data:{msg:"success",data:x.data}})}catch{}}else k.type==="api_call"?console.log("api_call",k):k.type==="ping"&&(o==null||o.postMessage({msgId:k.msgId,data:{msg:"pong"}}))},f=()=>{r(4,i=c)},p=()=>{var O;(O=s.contentWindow)==null||O.print()};function _(O){n=O,r(2,n)}function g(O){c=O,r(3,c)}function R(O){l=O,r(1,l)}function B(O){l=O,r(1,l)}const E=O=>{var k;try{let x=new MessageChannel;x.port1.onmessage=u,r(5,o=x.port1),console.log("chan.port2 type:",x.port2 instanceof MessagePort),(k=s==null?void 0:s.contentWindow)==null||k.postMessage("transfer_port","*",[x.port2])}catch(x){console.error("Error in postMessage:",x)}};function A(O){X[O?"unshift":"push"](()=>{s=O,r(0,s)})}return[s,l,n,c,i,o,u,f,p,_,g,R,B,E,A]}class sa extends Ie{constructor(e){super(),Ue(this,e,Yt,Wt,Ce,{})}}export{sa as component};
