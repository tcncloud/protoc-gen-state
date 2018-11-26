import { JSDOM } from 'jsdom';
console.log('Jasmine Helpers: Setting up jsdom...');

const dom = new JSDOM('<html><body></body></html>');
global.document = dom.window.document;
global.window = dom.window;
global.navigator = dom.window.navigator;

// rxjs hidden global dependencies
window.Object = Object;
window.Math = Math;
window.encodeURIComponent = encodeURIComponent;
window.localStorage = new LocalStorageMock();
window.sessionStorage = new LocalStorageMock();

function LocalStorageMock () {
  this.getItem = function (key) {
    return this[key];
  }

  this.setItem = function (key, value) {
    this[key] = value.toString();
  }

  this.clear = function () {
    Object.keys(this).forEach((key) => { delete this[key]; });
  }

  this.removeItem = function (key) {
    delete this[key];
  }
}
