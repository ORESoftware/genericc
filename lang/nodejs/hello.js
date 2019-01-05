const bindings = require('bindings');
const mod = bindings('hello');

console.log(mod.hello()); // 'world'