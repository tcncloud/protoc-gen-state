import jasmineEnzyme from 'jasmine-enzyme';
import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
console.log('Jasmine Helpers: Setting up Enzyme...');

configure({ adapter: new Adapter() });

beforeEach(function () {
  jasmineEnzyme();
});
