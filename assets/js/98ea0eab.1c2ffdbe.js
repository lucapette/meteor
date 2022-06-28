"use strict";(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[142],{3905:(e,t,r)=>{r.d(t,{Zo:()=>u,kt:()=>m});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=n.createContext({}),c=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},u=function(e){var t=c(e.components);return n.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},p=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),p=c(r),m=o,f=p["".concat(l,".").concat(m)]||p[m]||d[m]||i;return r?n.createElement(f,a(a({ref:t},u),{},{components:r})):n.createElement(f,a({ref:t},u))}));function m(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,a=new Array(i);a[0]=p;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:o,a[1]=s;for(var c=2;c<i;c++)a[c]=r[c];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}p.displayName="MDXCreateElement"},9139:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>d,frontMatter:()=>i,metadata:()=>s,toc:()=>c});var n=r(7462),o=(r(7294),r(3905));const i={},a="Guide",s={unversionedId:"contribute/guide",id:"contribute/guide",title:"Guide",description:"Adding a new Extractor",source:"@site/docs/contribute/guide.md",sourceDirName:"contribute",slug:"/contribute/guide",permalink:"/meteor/docs/contribute/guide",draft:!1,editUrl:"https://github.com/odpf/meteor/edit/master/docs/docs/contribute/guide.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Sinks",permalink:"/meteor/docs/reference/sinks"},next:{title:"Contribution Process",permalink:"/meteor/docs/contribute/contributing"}},l={},c=[{value:"Adding a new Extractor",id:"adding-a-new-extractor",level:2},{value:"Adding a new Processor",id:"adding-a-new-processor",level:2},{value:"Adding a new Sink",id:"adding-a-new-sink",level:2}],u={toc:c};function d(e){let{components:t,...r}=e;return(0,o.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"guide"},"Guide"),(0,o.kt)("h2",{id:"adding-a-new-extractor"},"Adding a new Extractor"),(0,o.kt)("p",null,"Please follow this list when adding a new Extractor:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Create unit test for the new extractor."),(0,o.kt)("li",{parentName:"ul"},"Register your extractor ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/main/plugins/extractors/populate.go"},"here"),". This is also where you would inject any dependencies needed for your extractor."),(0,o.kt)("li",{parentName:"ul"},"Create a markdown with your extractor details. ","(",(0,o.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/main/plugins/extractors/mysql/README.md"},"example"),")"),(0,o.kt)("li",{parentName:"ul"},"Add your extractor to one of the extractor list in ",(0,o.kt)("inlineCode",{parentName:"li"},"docs/reference/extractors.md"),".")),(0,o.kt)("h2",{id:"adding-a-new-processor"},"Adding a new Processor"),(0,o.kt)("p",null,"Please follow this list when adding a new Processor:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Create unit test for the new processor."),(0,o.kt)("li",{parentName:"ul"},"If the source instance is required for testing, Meteor provides a utility to easily create a docker container to help with your test as shown ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/main/plugins/extractors/mysql/extractor_test.go#L35"},"here"),"."),(0,o.kt)("li",{parentName:"ul"},"Register your processor ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/main/plugins/processors/populate.go"},"here"),". This is also where you would inject any dependencies needed for your processor."),(0,o.kt)("li",{parentName:"ul"},"Update ",(0,o.kt)("inlineCode",{parentName:"li"},"docs/reference/processors.md")," with guide to use the new processor.")),(0,o.kt)("h2",{id:"adding-a-new-sink"},"Adding a new Sink"),(0,o.kt)("p",null,"Please follow this list when adding a new Sink:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Create unit test for the new processor."),(0,o.kt)("li",{parentName:"ul"},"If the source instance is required for testing, Meteor provides a utility to easily create a docker container to help with your test as shown ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/main/plugins/extractors/mysql/extractor_test.go#L35"},"here"),"."),(0,o.kt)("li",{parentName:"ul"},"Register your sink ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/main/plugins/sinks/populate.go"},"here"),". This is also where you would inject any dependencies needed for your sink."),(0,o.kt)("li",{parentName:"ul"},"Update ",(0,o.kt)("inlineCode",{parentName:"li"},"docs/reference/sinks.md")," with guide to use the new sink.")))}d.isMDXComponent=!0}}]);