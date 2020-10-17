import { Component } from '@angular/core';

declare var window: any;

@Component({
  selector: '[id="app"]',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'my-app';

  clickMessage = '';
  constructor() {
    // console.log(window.backend.logService.Append().then(v => v));
  }

  onClickMe() {
  }
}
