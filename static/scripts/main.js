"use strict";(()=>{function p(){}function L(t){return t()}function B(){return Object.create(null)}function g(t){t.forEach(L)}function T(t){return typeof t=="function"}function G(t,e){return t!=t?e==e:t!==e||t&&typeof t=="object"||typeof t=="function"}function W(t){return Object.keys(t).length===0}function N(t){return t&&T(t.destroy)?t.destroy:p}var U=!1;function ut(){U=!0}function at(){U=!1}function y(t,e){t.appendChild(e)}function E(t,e,n){t.insertBefore(e,n||null)}function b(t){t.parentNode&&t.parentNode.removeChild(t)}function V(t,e){for(let n=0;n<t.length;n+=1)t[n]&&t[n].d(e)}function w(t){return document.createElement(t)}function k(t){return document.createTextNode(t)}function M(){return k(" ")}function X(){return k("")}function x(t,e,n,i){return t.addEventListener(e,n,i),()=>t.removeEventListener(e,n,i)}function m(t,e,n){n==null?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function ft(t){return Array.from(t.childNodes)}function j(t,e){e=""+e,t.wholeText!==e&&(t.data=e)}function I(t,e,n){t.classList[n?"add":"remove"](e)}var q;function S(t){q=t}var v=[];var H=[],$=[],z=[],dt=Promise.resolve(),O=!1;function _t(){O||(O=!0,dt.then(Y))}function D(t){$.push(t)}var A=new Set,F=0;function Y(){if(F!==0)return;let t=q;do{try{for(;F<v.length;){let e=v[F];F++,S(e),ht(e.$$)}}catch(e){throw v.length=0,F=0,e}for(S(null),v.length=0,F=0;H.length;)H.pop()();for(let e=0;e<$.length;e+=1){let n=$[e];A.has(n)||(A.add(n),n())}$.length=0}while(v.length);for(;z.length;)z.pop()();O=!1,A.clear(),S(t)}function ht(t){if(t.fragment!==null){t.update(),g(t.before_update);let e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(D)}}function mt(t){let e=[],n=[];$.forEach(i=>t.indexOf(i)===-1?e.push(i):n.push(i)),n.forEach(i=>i()),$=e}var pt=new Set;function yt(t,e){t&&t.i&&(pt.delete(t),t.i(e))}var Ot=typeof window<"u"?window:typeof globalThis<"u"?globalThis:global;var gt=["allowfullscreen","allowpaymentrequest","async","autofocus","autoplay","checked","controls","default","defer","disabled","formnovalidate","hidden","inert","ismap","itemscope","loop","multiple","muted","nomodule","novalidate","open","playsinline","readonly","required","reversed","selected"],Dt=new Set([...gt]);function bt(t,e,n,i){let{fragment:c,after_update:_}=t.$$;c&&c.m(e,n),i||D(()=>{let o=t.$$.on_mount.map(L).filter(T);t.$$.on_destroy?t.$$.on_destroy.push(...o):g(o),t.$$.on_mount=[]}),_.forEach(D)}function K(t,e){let n=t.$$;n.fragment!==null&&(mt(n.after_update),g(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function wt(t,e){t.$$.dirty[0]===-1&&(v.push(t),_t(),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function Q(t,e,n,i,c,_,o,u=[-1]){let d=q;S(t);let s=t.$$={fragment:null,ctx:[],props:_,update:p,not_equal:c,bound:B(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(e.context||(d?d.$$.context:[])),callbacks:B(),dirty:u,skip_bound:!1,root:e.target||d.$$.root};o&&o(s.root);let a=!1;if(s.ctx=n?n(t,e.props||{},(r,h,...l)=>{let f=l.length?l[0]:h;return s.ctx&&c(s.ctx[r],s.ctx[r]=f)&&(!s.skip_bound&&s.bound[r]&&s.bound[r](f),a&&wt(t,r)),h}):[],s.update(),a=!0,g(s.before_update),s.fragment=i?i(s.ctx):!1,e.target){if(e.hydrate){ut();let r=ft(e.target);s.fragment&&s.fragment.l(r),r.forEach(b)}else s.fragment&&s.fragment.c();e.intro&&yt(t.$$.fragment),bt(t,e.target,e.anchor,e.customElement),at(),Y()}S(d)}var Ft;typeof HTMLElement=="function"&&(Ft=class extends HTMLElement{constructor(){super(),this.attachShadow({mode:"open"})}connectedCallback(){let{on_mount:t}=this.$$;this.$$.on_disconnect=t.map(L).filter(T);for(let e in this.$$.slotted)this.appendChild(this.$$.slotted[e])}attributeChangedCallback(t,e,n){this[t]=n}disconnectedCallback(){g(this.$$.on_disconnect)}$destroy(){K(this,1),this.$destroy=p}$on(t,e){if(!T(e))return p;let n=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return n.push(e),()=>{let i=n.indexOf(e);i!==-1&&n.splice(i,1)}}$set(t){this.$$set&&!W(t)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}});var C=class{$destroy(){K(this,1),this.$destroy=p}$on(e,n){if(!T(n))return p;let i=this.$$.callbacks[e]||(this.$$.callbacks[e]=[]);return i.push(n),()=>{let c=i.indexOf(n);c!==-1&&i.splice(c,1)}}$set(e){this.$$set&&!W(e)&&(this.$$.skip_bound=!0,this.$$set(e),this.$$.skip_bound=!1)}};function J(t,e,n){let i=t.slice();return i[11]=e[n],i[13]=n,i}function Z(t){let e,n,i,c,_,o,u,d,s=t[0],a=[];for(let r=0;r<s.length;r+=1)a[r]=tt(J(t,s,r));return{c(){e=w("div"),n=w("div"),i=w("input"),_=M();for(let r=0;r<a.length;r+=1)a[r].c();m(i,"class","search__field"),m(i,"type","search"),m(i,"placeholder","Search composers by last name"),m(n,"class","search"),m(e,"class","search-wrapper")},m(r,h){E(r,e,h),y(e,n),y(n,i),y(n,_);for(let l=0;l<a.length;l+=1)a[l]&&a[l].m(n,null);u||(d=[x(i,"input",t[6]),N(c=Et.call(null,i)),N(o=kt.call(null,n,t[4])),x(e,"keydown",t[5])],u=!0)},p(r,h){if(h&133){s=r[0];let l;for(l=0;l<s.length;l+=1){let f=J(r,s,l);a[l]?a[l].p(f,h):(a[l]=tt(f),a[l].c(),a[l].m(n,null))}for(;l<a.length;l+=1)a[l].d(1);a.length=s.length}},d(r){r&&b(e),V(a,r),u=!1,g(d)}}}function tt(t){let e,n,i=t[11].lastName+"",c,_,o=t[11].firstName+"",u,d,s,a,r;function h(){return t[8](t[13])}return{c(){e=w("a"),n=w("div"),c=k(i),_=k(", "),u=k(o),d=M(),m(n,"class","search__result"),I(n,"search__result_selected",t[2]===t[13]),m(e,"href",s="/composer/"+t[11].slug)},m(l,f){E(l,e,f),y(e,n),y(n,c),y(n,_),y(n,u),y(e,d),a||(r=x(e,"mouseenter",h),a=!0)},p(l,f){t=l,f&1&&i!==(i=t[11].lastName+"")&&j(c,i),f&1&&o!==(o=t[11].firstName+"")&&j(u,o),f&4&&I(n,"search__result_selected",t[2]===t[13]),f&1&&s!==(s="/composer/"+t[11].slug)&&m(e,"href",s)},d(l){l&&b(e),a=!1,r()}}}function vt(t){let e,n,i,c,_,o=t[1]&&Z(t);return{c(){e=w("div"),e.innerHTML='<img class="icon" src="/static/img/search-icon.svg" alt="Search"/>',n=M(),o&&o.c(),i=X(),m(e,"class","search-button label")},m(u,d){E(u,e,d),E(u,n,d),o&&o.m(u,d),E(u,i,d),c||(_=[x(e,"click",t[3]),x(e,"keypress",t[3])],c=!0)},p(u,[d]){u[1]?o?o.p(u,d):(o=Z(u),o.c(),o.m(i.parentNode,i)):o&&(o.d(1),o=null)},i:p,o:p,d(u){u&&b(e),u&&b(n),o&&o.d(u),u&&b(i),c=!1,g(_)}}}async function $t(t){try{let e=await fetch(`/api/search?q=${t}`);return e.ok?await e.json():[]}catch(e){return console.log(e),[]}}function Et(t){t.focus()}function kt(t,e){let n=i=>t&&!t.contains(i.target)&&!i.defaultPrevented&&e();return document.addEventListener("click",n,!0),{destroy(){document.removeEventListener("click",n,!0)}}}function xt(t,e,n){"use strict";let i=[],c=[void 0,void 0],_=!1,o=0;async function u(){c[0]!==void 0?(n(0,i=await $t(c[0])),c[1]!==void 0?(c[0]=c[1],c[1]=void 0,await u()):c[0]=void 0):n(0,i=[])}function d(){n(1,_=!0)}function s(){n(0,i=[]),n(1,_=!1)}function a(f){f.code==="ArrowUp"&&i.length>0?n(2,o=o>0?o-1:i.length-1):f.code==="ArrowDown"?n(2,o=o<i.length-1?o+1:n(2,o=0)):f.code==="Escape"?s():f.code==="Enter"&&i.length>0&&(location.pathname=`/composer/${i[o].slug}`)}function r(f){let P=f.target.value||void 0;c[0]===void 0?(c[0]=P,u()):c[1]=P}function h(f){n(2,o=f)}return[i,_,o,d,s,a,r,h,f=>h(f)]}var R=class extends C{constructor(e){super(),Q(this,e,xt,vt,G,{})}},et=R;var nt=t=>!!t&&t instanceof HTMLInputElement,it=t=>{document.documentElement.dataset.theme=t},ot=t=>{localStorage.setItem("theme",t)},St=()=>localStorage.getItem("theme"),st=t=>t?"dark":"light",rt=(t,e)=>{nt(t)&&(t.checked=e)},Tt=({target:t})=>{if(nt(t)){let e=st(t.checked);ot(e),it(e)}},Ct=(t,e,n)=>{let i=st(n);ot(i),it(i),rt(e,n)},Mt=t=>{t&&t.classList.remove("d-none")},At=()=>{let t=document.getElementById("switcher"),e=window.matchMedia("(prefers-color-scheme: dark)");t&&(t.addEventListener("change",Tt),e.addEventListener("change",i=>{Ct(i,t,e.matches)}),St()==="dark"&&rt(t,!0),Mt(document.querySelector(".toggle-switch__label")))},ct=()=>{document.addEventListener("DOMContentLoaded",()=>At())};ct();var lt=document.getElementById("searchBlock");lt&&new et({target:lt});})();
