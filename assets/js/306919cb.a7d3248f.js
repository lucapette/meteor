"use strict";(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[825],{3905:function(e,t,r){r.d(t,{Zo:function(){return p},kt:function(){return m}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=n.createContext({}),u=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},p=function(e){var t=u(e.components);return n.createElement(l.Provider,{value:t},e.children)},f={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},s=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,l=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),s=u(r),m=o,d=s["".concat(l,".").concat(m)]||s[m]||f[m]||a;return r?n.createElement(d,i(i({ref:t},p),{},{components:r})):n.createElement(d,i({ref:t},p))}));function m(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,i=new Array(a);i[0]=s;var c={};for(var l in t)hasOwnProperty.call(t,l)&&(c[l]=t[l]);c.originalType=e,c.mdxType="string"==typeof e?e:o,i[1]=c;for(var u=2;u<a;u++)i[u]=r[u];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}s.displayName="MDXCreateElement"},4480:function(e,t,r){r.r(t),r.d(t,{contentTitle:function(){return l},default:function(){return s},frontMatter:function(){return c},metadata:function(){return u},toc:function(){return p}});var n=r(7462),o=r(3366),a=(r(7294),r(3905)),i=["components"],c={},l="Configuration",u={unversionedId:"reference/configuration",id:"reference/configuration",isDocsHomePage:!1,title:"Configuration",description:"This page contains references for all the application configurations for Meteor.",source:"@site/docs/reference/configuration.md",sourceDirName:"reference",slug:"/reference/configuration",permalink:"/meteor/docs/reference/configuration",editUrl:"https://github.com/odpf/meteor/edit/master/docs/docs/reference/configuration.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Commands",permalink:"/meteor/docs/reference/commands"},next:{title:"Meteor Metadata Model",permalink:"/meteor/docs/reference/metadata_models"}},p=[{value:"Table of Contents",id:"table-of-contents",children:[]},{value:"Generic",id:"generic",children:[{value:"<code>PORT</code>",id:"port",children:[]},{value:"<code>RECIPE_STORAGE_URL</code>",id:"recipe_storage_url",children:[]}]}],f={toc:p};function s(e){var t=e.components,r=(0,o.Z)(e,i);return(0,a.kt)("wrapper",(0,n.Z)({},f,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"configuration"},"Configuration"),(0,a.kt)("p",null,"This page contains references for all the application configurations for Meteor."),(0,a.kt)("h2",{id:"table-of-contents"},"Table of Contents"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/meteor/docs/reference/configuration#generic"},"Generic"))),(0,a.kt)("h2",{id:"generic"},"Generic"),(0,a.kt)("p",null,"Meteor's required variables to start using it."),(0,a.kt)("h3",{id:"port"},(0,a.kt)("inlineCode",{parentName:"h3"},"PORT")),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Example value: ",(0,a.kt)("inlineCode",{parentName:"li"},"8080")),(0,a.kt)("li",{parentName:"ul"},"Type: ",(0,a.kt)("inlineCode",{parentName:"li"},"optional")),(0,a.kt)("li",{parentName:"ul"},"Default: ",(0,a.kt)("inlineCode",{parentName:"li"},"3000")),(0,a.kt)("li",{parentName:"ul"},"Port to listen on.")),(0,a.kt)("h3",{id:"recipe_storage_url"},(0,a.kt)("inlineCode",{parentName:"h3"},"RECIPE_STORAGE_URL")),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Example value: ",(0,a.kt)("inlineCode",{parentName:"li"},"s3://my-bucket?region=us-west-1")),(0,a.kt)("li",{parentName:"ul"},"Type: ",(0,a.kt)("inlineCode",{parentName:"li"},"optional")),(0,a.kt)("li",{parentName:"ul"},"Default: ",(0,a.kt)("inlineCode",{parentName:"li"},"mem://")),(0,a.kt)("li",{parentName:"ul"},"Object storage URL to persist recipes. Can be a gcs, an aws bucket or even a local folder. Check this ",(0,a.kt)("a",{parentName:"li",href:"https://github.com/odpf/meteor/tree/27f39fe2f83b657d4ecb9eb2c2a8794c6c0671b6/docs/guides/setup_storage.md"},"guide")," for url format and how to setup each available storage.")))}s.isMDXComponent=!0}}]);