
why does it at first not load the name giving me this error:
Uncaught TypeError: Cannot read properties of null (reading 'name')
    at Posts.vue:60:98
    at Proxy.renderFnWithContext (chunk-P4JRBGFP.js?v=c71a8d54:2176:13)
    at Proxy.<anonymous> (vue-router.js?v=c71a8d54:1593:54)
    at renderComponentRoot (chunk-P4JRBGFP.js?v=c71a8d54:7779:17)
    at ReactiveEffect.componentUpdateFn [as fn] (chunk-P4JRBGFP.js?v=c71a8d54:6459:46)
    at ReactiveEffect.run (chunk-P4JRBGFP.js?v=c71a8d54:435:19)
    at instance.update (chunk-P4JRBGFP.js?v=c71a8d54:6590:17)
    at setupRenderEffect (chunk-P4JRBGFP.js?v=c71a8d54:6600:5)
    at mountComponent (chunk-P4JRBGFP.js?v=c71a8d54:6368:7)
    at processComponent (chunk-P4JRBGFP.js?v=c71a8d54:6322:9)
(anonymous) @ Posts.vue:60
renderFnWithContext @ chunk-P4JRBGFP.js?v=c71a8d54:2176
(anonymous) @ vue-router.js?v=c71a8d54:1593
renderComponentRoot @ chunk-P4JRBGFP.js?v=c71a8d54:7779
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6459
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
render2 @ chunk-P4JRBGFP.js?v=c71a8d54:7125
mount @ chunk-P4JRBGFP.js?v=c71a8d54:4112
app.mount @ chunk-P4JRBGFP.js?v=c71a8d54:11171
(anonymous) @ main.js:21
Show 59 more frames
Show lessUnderstand this error
picsum.photos/id/87/300/320:1 
        
        
       GET https://picsum.photos/id/87/300/320 net::ERR_NAME_NOT_RESOLVED
Image
patchDOMProp @ chunk-P4JRBGFP.js?v=c71a8d54:10294
patchProp @ chunk-P4JRBGFP.js?v=c71a8d54:10403
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:6011
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
render2 @ chunk-P4JRBGFP.js?v=c71a8d54:7125
mount @ chunk-P4JRBGFP.js?v=c71a8d54:4112
app.mount @ chunk-P4JRBGFP.js?v=c71a8d54:11171
(anonymous) @ main.js:21
Show 51 more frames
Show lessUnderstand this error
tab.js:1 Refused to execute inline script because it violates the following Content Security Policy directive: "script-src 'self' 'wasm-unsafe-eval' 'inline-speculation-rules'". Either the 'unsafe-inline' keyword, a hash ('sha256-kPx0AsF0oz2kKiZ875xSvv693TBHkQ/0SkMJZnnNpnQ='), or a nonce ('nonce-...') is required to enable inline execution.

(anonymous) @ tab.js:1
(anonymous) @ tab.js:1
Show 2 more frames
Show lessUnderstand this error
User.vue:205 [Vue warn]: Failed to resolve component: Head
If this is a native custom element, make sure to exclude it from component resolution via compilerOptions.isCustomElement. 
  at <User onVnodeUnmounted=fn<onVnodeUnmounted> ref=Ref< undefined > > 
  at <RouterView> 
  at <App>
warn$1 @ chunk-P4JRBGFP.js?v=c71a8d54:1528
resolveAsset @ chunk-P4JRBGFP.js?v=c71a8d54:3079
resolveComponent @ chunk-P4JRBGFP.js?v=c71a8d54:3041
_sfc_render @ User.vue:205
renderComponentRoot @ chunk-P4JRBGFP.js?v=c71a8d54:7779
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6459
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6546
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
callWithErrorHandling @ chunk-P4JRBGFP.js?v=c71a8d54:1678
flushJobs @ chunk-P4JRBGFP.js?v=c71a8d54:1891
Promise.then
queueFlush @ chunk-P4JRBGFP.js?v=c71a8d54:1801
queuePostFlushCb @ chunk-P4JRBGFP.js?v=c71a8d54:1821
queueEffectWithSuspense @ chunk-P4JRBGFP.js?v=c71a8d54:8597
scheduler @ chunk-P4JRBGFP.js?v=c71a8d54:7444
resetScheduling @ chunk-P4JRBGFP.js?v=c71a8d54:516
triggerEffects @ chunk-P4JRBGFP.js?v=c71a8d54:560
triggerRefValue @ chunk-P4JRBGFP.js?v=c71a8d54:1326
set value @ chunk-P4JRBGFP.js?v=c71a8d54:1373
finalizeNavigation @ vue-router.js?v=c71a8d54:2487
(anonymous) @ vue-router.js?v=c71a8d54:2397
Promise.then
pushWithRedirect @ vue-router.js?v=c71a8d54:2365
push @ vue-router.js?v=c71a8d54:2291
install @ vue-router.js?v=c71a8d54:2644
use @ chunk-P4JRBGFP.js?v=c71a8d54:4034
(anonymous) @ main.js:19
Show 30 more frames
Show lessUnderstand this warning
main.js:19 [Vue warn]: Missing required prop: "to" 
  at <RouterLink class="mr-2" > 
  at <CreatePostBox image="https://picsum.photos/id/140/300/320" placeholder="What's on your mind Raymond?!" > 
  at <MainNavLayout> 
  at <User onVnodeUnmounted=fn<onVnodeUnmounted> ref=Ref< undefined > > 
  at <RouterView> 
  at <App>
warn$1 @ chunk-P4JRBGFP.js?v=c71a8d54:1528
validateProp @ chunk-P4JRBGFP.js?v=c71a8d54:4529
validateProps @ chunk-P4JRBGFP.js?v=c71a8d54:4517
initProps @ chunk-P4JRBGFP.js?v=c71a8d54:4210
setupComponent @ chunk-P4JRBGFP.js?v=c71a8d54:9138
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6356
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6546
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
callWithErrorHandling @ chunk-P4JRBGFP.js?v=c71a8d54:1678
flushJobs @ chunk-P4JRBGFP.js?v=c71a8d54:1891
Promise.then
queueFlush @ chunk-P4JRBGFP.js?v=c71a8d54:1801
queuePostFlushCb @ chunk-P4JRBGFP.js?v=c71a8d54:1821
queueEffectWithSuspense @ chunk-P4JRBGFP.js?v=c71a8d54:8597
scheduler @ chunk-P4JRBGFP.js?v=c71a8d54:7444
resetScheduling @ chunk-P4JRBGFP.js?v=c71a8d54:516
triggerEffects @ chunk-P4JRBGFP.js?v=c71a8d54:560
triggerRefValue @ chunk-P4JRBGFP.js?v=c71a8d54:1326
set value @ chunk-P4JRBGFP.js?v=c71a8d54:1373
finalizeNavigation @ vue-router.js?v=c71a8d54:2487
(anonymous) @ vue-router.js?v=c71a8d54:2397
Promise.then
pushWithRedirect @ vue-router.js?v=c71a8d54:2365
push @ vue-router.js?v=c71a8d54:2291
install @ vue-router.js?v=c71a8d54:2644
use @ chunk-P4JRBGFP.js?v=c71a8d54:4034
(anonymous) @ main.js:19
Show 77 more frames
Show lessUnderstand this warning
main.js:19 [Vue Router warn]: Invalid value for prop "to" in useLink()
- to: undefined 
- props: Proxy(Object) {replace: false, custom: false, ariaCurrentValue: 'page', to: undefined, activeClass: undefined, …}
warn @ vue-router.js?v=c71a8d54:44
(anonymous) @ vue-router.js?v=c71a8d54:1493
(anonymous) @ chunk-P4JRBGFP.js?v=c71a8d54:1246
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
get value @ chunk-P4JRBGFP.js?v=c71a8d54:1258
useLink @ vue-router.js?v=c71a8d54:1538
setup @ vue-router.js?v=c71a8d54:1581
callWithErrorHandling @ chunk-P4JRBGFP.js?v=c71a8d54:1678
setupStatefulComponent @ chunk-P4JRBGFP.js?v=c71a8d54:9179
setupComponent @ chunk-P4JRBGFP.js?v=c71a8d54:9140
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6356
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6546
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
callWithErrorHandling @ chunk-P4JRBGFP.js?v=c71a8d54:1678
flushJobs @ chunk-P4JRBGFP.js?v=c71a8d54:1891
Promise.then
queueFlush @ chunk-P4JRBGFP.js?v=c71a8d54:1801
queuePostFlushCb @ chunk-P4JRBGFP.js?v=c71a8d54:1821
queueEffectWithSuspense @ chunk-P4JRBGFP.js?v=c71a8d54:8597
scheduler @ chunk-P4JRBGFP.js?v=c71a8d54:7444
resetScheduling @ chunk-P4JRBGFP.js?v=c71a8d54:516
triggerEffects @ chunk-P4JRBGFP.js?v=c71a8d54:560
triggerRefValue @ chunk-P4JRBGFP.js?v=c71a8d54:1326
set value @ chunk-P4JRBGFP.js?v=c71a8d54:1373
finalizeNavigation @ vue-router.js?v=c71a8d54:2487
(anonymous) @ vue-router.js?v=c71a8d54:2397
Promise.then
pushWithRedirect @ vue-router.js?v=c71a8d54:2365
push @ vue-router.js?v=c71a8d54:2291
install @ vue-router.js?v=c71a8d54:2644
use @ chunk-P4JRBGFP.js?v=c71a8d54:4034
(anonymous) @ main.js:19
Show 82 more frames
Show lessUnderstand this warning
main.js:19 [Vue Router warn]: router.resolve() was passed an invalid location. This will fail in production.
- Location: undefined
warn @ vue-router.js?v=c71a8d54:44
resolve @ vue-router.js?v=c71a8d54:2218
(anonymous) @ vue-router.js?v=c71a8d54:1501
(anonymous) @ chunk-P4JRBGFP.js?v=c71a8d54:1246
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
get value @ chunk-P4JRBGFP.js?v=c71a8d54:1258
useLink @ vue-router.js?v=c71a8d54:1538
setup @ vue-router.js?v=c71a8d54:1581
callWithErrorHandling @ chunk-P4JRBGFP.js?v=c71a8d54:1678
setupStatefulComponent @ chunk-P4JRBGFP.js?v=c71a8d54:9179
setupComponent @ chunk-P4JRBGFP.js?v=c71a8d54:9140
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6356
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
mountElement @ chunk-P4JRBGFP.js?v=c71a8d54:5993
processElement @ chunk-P4JRBGFP.js?v=c71a8d54:5958
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5826
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
mountChildren @ chunk-P4JRBGFP.js?v=c71a8d54:6070
processFragment @ chunk-P4JRBGFP.js?v=c71a8d54:6252
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5812
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6466
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
setupRenderEffect @ chunk-P4JRBGFP.js?v=c71a8d54:6600
mountComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6368
processComponent @ chunk-P4JRBGFP.js?v=c71a8d54:6322
patch @ chunk-P4JRBGFP.js?v=c71a8d54:5838
componentUpdateFn @ chunk-P4JRBGFP.js?v=c71a8d54:6546
run @ chunk-P4JRBGFP.js?v=c71a8d54:435
instance.update @ chunk-P4JRBGFP.js?v=c71a8d54:6590
callWithErrorHandling @ chunk-P4JRBGFP.js?v=c71a8d54:1678
flushJobs @ chunk-P4JRBGFP.js?v=c71a8d54:1891
Promise.then
queueFlush @ chunk-P4JRBGFP.js?v=c71a8d54:1801
queuePostFlushCb @ chunk-P4JRBGFP.js?v=c71a8d54:1821
queueEffectWithSuspense @ chunk-P4JRBGFP.js?v=c71a8d54:8597
scheduler @ chunk-P4JRBGFP.js?v=c71a8d54:7444
resetScheduling @ chunk-P4JRBGFP.js?v=c71a8d54:516
triggerEffects @ chunk-P4JRBGFP.js?v=c71a8d54:560
triggerRefValue @ chunk-P4JRBGFP.js?v=c71a8d54:1326
set value @ chunk-P4JRBGFP.js?v=c71a8d54:1373
finalizeNavigation @ vue-router.js?v=c71a8d54:2487
(anonymous) @ vue-router.js?v=c71a8d54:2397
Promise.then
pushWithRedirect @ vue-router.js?v=c71a8d54:2365
push @ vue-router.js?v=c71a8d54:2291
install @ vue-router.js?v=c71a8d54:2644
use @ chunk-P4JRBGFP.js?v=c71a8d54:4034
(anonymous) @ main.js:19
Show 83 more frames
Show lessUnderstand this warning
api.js:37 Adding token to headers: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzIzNzI2MDMsInN1YiI6M30.QWLxRA_aN9_ii6JGfLo6lBeI5j8hqJbZx446WgDit8U
picsum.photos/id/87/300/320:1 
        
        
       GET https://picsum.photos/id/87/300/320 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/45/2000/320:1 
        
        
       GET https://picsum.photos/id/45/2000/320 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/200/500/500:1 
        
        
       GET https://picsum.photos/id/200/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/141/500/500:1 
        
        
       GET https://picsum.photos/id/141/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/142/500/500:1 
        
        
       GET https://picsum.photos/id/142/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/143/500/500:1 
        
        
       GET https://picsum.photos/id/143/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/144/500/500:1 
        
        
       GET https://picsum.photos/id/144/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/145/500/500:1 
        
        
       GET https://picsum.photos/id/145/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/146/500/500:1 
        
        
       GET https://picsum.photos/id/146/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/147/500/500:1 
        
        
       GET https://picsum.photos/id/147/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/149/500/500:1 
        
        
       GET https://picsum.photos/id/149/500/500 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/78/300/300:1 
        
        
       GET https://picsum.photos/id/78/300/300 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/140/300/320:1 
        
        
       GET https://picsum.photos/id/140/300/320 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/189/800/800:1 
        
        
       GET https://picsum.photos/id/189/800/800 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/199/800/800:1 
        
        
       GET https://picsum.photos/id/199/800/800 net::ERR_NAME_NOT_RESOLVEDUnderstand this error
picsum.photos/id/299/800/800:1 
        
        
       GET https://picsum.photos/id/299/800/800 net::ERR_NAME_NOT_RESOLVEDUnderstand this error


then when i first comment this {{userDetails.name}} then i refresh the page then i uncomment that's when it access the name and other properties on user object:

user object:
{
    "ID": 3,
    "CreatedAt": "2024-10-18T12:54:37.31008+03:00",
    "UpdatedAt": "2024-10-18T12:54:37.31008+03:00",
    "DeletedAt": null,
    "name": "Joe Fazer",
    "email": "Joe2@gmail.com",
    "password": "$2a$10$.iun931nu4TZqO.iMmDoOOP/2vUmrhOqDQgY6JsxUXKbWyymcmZOi",
    "posts": null,
    "comments": null,
    "image": ""
}



I think the issue is to do with something about when the component is mounted:
<script setup>
import { onMounted, ref } from "vue";
import CreatePostOverlay from "@/components/CreatePostOverlay.vue";

import Magnify from "vue-material-design-icons/Magnify.vue";
import Home from "vue-material-design-icons/Home.vue";
import HomeOutline from "vue-material-design-icons/HomeOutline.vue";
import TelevisionPlay from "vue-material-design-icons/TelevisionPlay.vue";
import StorefrontOutline from "vue-material-design-icons/StorefrontOutline.vue";
import AccountGroup from "vue-material-design-icons/AccountGroup.vue";
import ControllerClassicOutline from "vue-material-design-icons/ControllerClassicOutline.vue";
import DotsGrid from "vue-material-design-icons/DotsGrid.vue";
import FacebookMessenger from "vue-material-design-icons/FacebookMessenger.vue";
import Bell from "vue-material-design-icons/Bell.vue";
import Logout from "vue-material-design-icons/Logout.vue";
// Removed unused imports for brevity
import api from "@/config/api";

import CropperModal from "@/components/CropperModal.vue";
import ImageDisplay from "@/components/ImageDisplay.vue";

const handleLogout = () => {};

import { useGeneralStore } from "@/stores/general";
import { storeToRefs } from "pinia";
import { useAuthStore } from "@/stores/auth";
const useGeneral = useGeneralStore();
const { isPostOverlay, isCropperModal, isImageDisplay } =
  storeToRefs(useGeneral);

let showMenu = ref(false);
const userDetails = ref(null);
const authStore = useAuthStore();

// const getUserDetails = async () => {
//   // const token = authStore.token
//   const response = await api.get("/user");
//   userDetails.value = response.data
//   console.log(userDetails)
// }

// const getUserDetails = async () => {
//   const token = authStore.token; // Fetch the token from the authStore
//   const response = await api.get("/user", {
//     headers: {
//       Authorization: Bearer ${token}, // Attach the token
//     },
//   });
//   userDetails.value = response.data;
// };

const getUserDetails = async () => {
  const token = authStore.token;
  if (!token) {
    console.error("No token found");
    return;
  }
  
  try {
    // const response = await api.get("/user", {
    const response = await api.get("/validate", {
      headers: {
        Authorization: ${token},
      },
    });
    userDetails.value = response.data;
  } catch (error) {
    console.error("Failed to fetch user details", error);
  }
};



onMounted(() => {
  getUserDetails();
});
</script>


<template>
  <div
    id="MainNav"
    class="flex w-full fixed z-50 items-center justify-between bg-white shadow-xl border-b h-[56px]"
  >
    {{userDetails}} works immediately when i log into the system
{{userDetails.name}} doesn't work immediately when i log into the system//
    <div id="NavLeft" class="flex items-center justify-start w-[260px]">
      <router-link to="/home" class="pl-3 min-w-[55px]">
        <img src="/public/icons/FacebookLogoCircle.png" alt="" />
      </router-link>
      <div
        class="flex items-center justify-center p-1 rounded-full ml-2 bg-[#EFF2F5] h-[40px]"
      >
        <Magnify class="p-1" :size="22" fillColor="#64676B" />
        <input
          type="text"
          placeholder="Search Facebook"
          class="bg-[#EFF2F5] placeholder-[#64676B] lg:block hidden border-none p-0 ring-0 focus:ring-0"
        />
      </div>
    </div>

    <div
      id="NavCenter"
      class="hidden lg:flex items-center ml-5 justify-center w-8/12 max-w-[600px]"
    >
      <router-link to="/home" class="w-full">
        <div
          :class="$route.path === '/' ? 'mt-1.5' : ''"
          class="flex items-center justify-center p-1 w-full rounded-lg h-[48px] cursor-pointer hover:bg-[#F2F2F2]"
        >
          <div>
            <Home
              v-if="$route.path === '/home'"
              class="mx-auto"
              :size="27"
              fillColor="#1A73E3"
            />
            <HomeOutline
              v-else
              class="mx-auto"
              :size="32"
              fillColor="#64676B"
            />
          </div>
        </div>

        <div
          v-if="$route.path === '/'"
          class="border-b-4 border-b-blue-400 rounded-md"
        ></div>
      </router-link>
      <button
        class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]"
      >
        <TelevisionPlay class="mx-auto" :size="27" fillColor="#64676B" />
      </button>
      <button
        class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]"
      >
        <StorefrontOutline class="mx-auto" :size="27" fillColor="#64676B" />
      </button>
      <button
        class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]"
        @click="getUserDetails"
      >
        <span class="rounded-full p-0.5 border-[2px] border-[#64676B]">
          <AccountGroup class="mx-auto" :size="20" fillColor="#64676B" />
        </span>
      </button>
      <button
        class="flex items-center p-1 justify-center w-full rounded-lg mx-1 cursor-pointer h-[48px] hover:bg-[#F2F2F2]"
      >
        <ControllerClassicOutline
          class="mx-auto"
          :size="32"
          fillColor="#64676B"
        />
      </button>
    </div>

    <div class="flex items-center justify-end mr-4 w-2/12">
      <button
        class="rounded-full p-2 hover:bg-gray-300 mx-1 cursor-pointer bg-[#E3E6EA]"
      >
        <DotsGrid :size="23" fillColor="#050505" />
      </button>
      <button
        class="rounded-full p-2 hover:bg-gray-300 mx-1 cursor-pointer bg-[#E3E6EA]"
      >
        <FacebookMessenger :size="23" fillColor="#050505" />
      </button>
      <button
        class="rounded-full p-2 hover:bg-gray-300 mx-1 cursor-pointer bg-[#E3E6EA]"
      >
        <Bell :size="23" fillColor="#050505" />
      </button>
      <div class="flex items-center justify-center relative">
        <button @click="showMenu = !showMenu">
          <img
            src="https://picsum.photos/id/87/300/320"
            alt=""
            class="rounded-full ml-1 cursor-pointer min-w-[40px] max-h-[40px]"
          />
        </button>
        <div
          v-if="showMenu"
          class="absolute bg-white shadow-xl top-10 rounded-lg p-1 border mt-1 w-[330px] right-0"
        >
          <router-link to="/" @click="showMenu = !showMenu">
            <div
              class="flex items-center gap-3 rounded-lg hover:bg-gray-200 p-2"
            >
              <img
                src="https://picsum.photos/id/87/300/320"
                alt=""
                class="rounded-full ml-1 cursor-pointer min-w-[35px] max-h-[35px]"
              />
              <span>Raymond Mwebe</span>
              <!-- <span>{{userDetails.user.name}}</span> -->
            </div>
          </router-link>
          <button
            class="w-full flex items-center gap-3 hover:bg-gray-200 px-2 py-2.5 rounded-lg"
            as="button"
            @click="handleLogout"
            method="post"
          >
            <Logout class="pl-2" :size="30" />
            <span>Logout</span>
          </button>
          <div class="text-xs font-semi-bold p-2 pt-3 border-t mt-2">
            Privacy . Terms . Advertising . AdChoices . Cookies . Meta
          </div>
        </div>
      </div>
    </div>
  </div>
  <slot />
  <CreatePostOverlay v-if="isPostOverlay" @showModal="isPostOverlay = false" />

  <CropperModal v-if="isCropperModal" @showModal="isCropperModal = false" />

  <ImageDisplay v-if="isImageDisplay" />
</template>