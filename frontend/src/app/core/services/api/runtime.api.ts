import {Backend} from './backend.api';


export interface Runtime extends Window {
  backend: Backend;
}
